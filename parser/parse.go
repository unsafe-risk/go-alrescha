package parser

import (
	"encoding/json"
	"sort"

	"github.com/buger/jsonparser"
	"github.com/unsafe-risk/go-alrescha/nameconv"
)

type IDLFile struct {
	Version int

	Types []IDLType

	IndexPathMap       map[int]string
	TypeIndexMap       map[string]int
	IsArrayMap         map[int]bool
	IndexArrayIndexMap map[int]int

	Ctr int
}

func (f *IDLFile) GetType(t string) *IDLType {
	for i := range f.Types {
		if f.Types[i].Name == t {
			return &f.Types[i]
		}
	}
	return nil
}

type IDLType struct {
	Name string

	Fields []IDLField

	Index int
}

type IDLField struct {
	FieldType string // array, var, list...
	Key       string
	Type      string
	CodeType  string

	Index int

	IsRawType bool
	IsFixed   bool
	IsArray   bool
	IsList    bool

	ArrayLength int

	ArrayIndex int
}

type GernerateInfo struct {
	Structs []*GenerateStruct
	IDL     *IDLFile

	Max MaxLen
}

type MaxLen struct {
	StringMaxLen int `json:"$StringMaxLen"`
	BytesMaxLen  int `json:"$BytesMaxLen"`
	ListMaxItem  int `json:"$ListMaxItem"`
}

func ParseGenerateInfo(data []byte) (*GernerateInfo, error) {
	idl, err := ParseJSONData(data)
	if err != nil {
		return nil, err
	}
	structIDL := &IDLFile{}
	err = DeepCopy(idl, structIDL)
	if err != nil {
		return nil, err
	}
	for i := range structIDL.Types {
		t := &structIDL.Types[i]
		var fields []IDLField = make([]IDLField, len(t.Fields))
		copy(fields, t.Fields)
		t.Fields = t.Fields[:0]
		for j := range fields {
			f := &fields[j]
			if f.FieldType == "var" {
				t.Fields = append(t.Fields, *f)
			} else if f.FieldType == "array" {
				for k := 0; k < f.ArrayLength; k++ {
					fk := *f
					fk.ArrayIndex = k
					structIDL.Ctr++
					structIDL.IndexPathMap[structIDL.Ctr] = idl.IndexPathMap[f.Index]
					structIDL.IsArrayMap[structIDL.Ctr] = idl.IsArrayMap[f.Index]
					structIDL.IndexArrayIndexMap[structIDL.Ctr] = k
					fk.Index = structIDL.Ctr
					//fmt.Println(fk.Index)
					t.Fields = append(t.Fields, fk)
				}
			} else if f.FieldType == "list" {
				t.Fields = append(t.Fields, *f)
			}
		}
	}
	//fmt.Println(structIDL)
	var structs []*GenerateStruct
	for _, t := range structIDL.Types {
		var fields []*GenerateField
		for _, f := range t.Fields {
			TraceType(structIDL, &fields, f, nil)
		}
		//fmt.Println("// " + fmt.Sprint(fields))
		sort.Sort(GenTypes(fields))
		off := 0
		for _, f := range fields {
			if f.IsFixed {
				f.Offset = off
				off += f.Size
			}
		}
		structs = append(structs, &GenerateStruct{
			Name:  structIDL.IndexPathMap[t.Index],
			Types: fields,
		})
	}
	idl.IndexArrayIndexMap = structIDL.IndexArrayIndexMap
	idl.IndexPathMap = structIDL.IndexPathMap
	idl.IsArrayMap = structIDL.IsArrayMap
	geninfo := &GernerateInfo{
		Structs: structs,
		IDL:     idl,
	}
	err = json.Unmarshal(data, &geninfo.Max)
	if err != nil {
		return nil, err
	}
	if geninfo.Max.StringMaxLen == 0 {
		geninfo.Max.StringMaxLen = 65535
	}
	if geninfo.Max.BytesMaxLen == 0 {
		geninfo.Max.BytesMaxLen = 65535
	}
	if geninfo.Max.ListMaxItem == 0 {
		geninfo.Max.ListMaxItem = 65535
	}
	return geninfo, nil
}

func ParseJSONData(data []byte) (*IDLFile, error) {
	idl := new(IDLFile)
	idl.IndexPathMap = make(map[int]string)
	idl.TypeIndexMap = make(map[string]int)
	idl.IsArrayMap = make(map[int]bool)
	idl.IndexArrayIndexMap = make(map[int]int)
	ctr := 0
	ver, err := jsonparser.GetInt(data, "$version")
	if err != nil {
		return nil, err
	}
	idl.Version = int(ver)

	err = jsonparser.ObjectEach(data, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
		ctr++
		t := IDLType{}
		t.Name = string(key)
		t.Index = ctr
		idl.IndexPathMap[ctr] = nameconv.Snake2Pascal(string(key))
		idl.TypeIndexMap[t.Name] = ctr

		jsonparser.ObjectEach(value, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
			ctr++
			var CodeType string

			switch TypeOf(string(value)) {
			case TYPE_ARRAY:
				Type, Size, err := ParseArray(string(value))
				if err != nil {
					return err
				}
				isRaw, isFixed, _ := GetRawTypeInfo(Type)
				if isRaw {
					CodeType = string(Type)
				} else {
					CodeType = nameconv.Snake2Pascal(Type)
				}
				t.Fields = append(t.Fields, IDLField{
					FieldType:   "array",
					Key:         string(key),
					Type:        Type,
					Index:       ctr,
					CodeType:    CodeType,
					IsRawType:   isRaw,
					IsFixed:     isFixed,
					IsArray:     true,
					ArrayLength: Size,
				})
				idl.IndexPathMap[ctr] = nameconv.Snake2Pascal(string(key))
				idl.IsArrayMap[ctr] = true
			case TYPE_LIST:
				Type, err := ParseList(string(value))
				if err != nil {
					return err
				}
				isRaw, _, _ := GetRawTypeInfo(Type)
				if isRaw {
					CodeType = string(Type)
				} else {
					CodeType = nameconv.Snake2Pascal(Type)
				}
				t.Fields = append(t.Fields, IDLField{
					FieldType: "list",
					Key:       string(key),
					Type:      Type,
					Index:     ctr,
					CodeType:  CodeType,
					IsRawType: isRaw,
					IsFixed:   false,
					IsArray:   false,
					IsList:    true,
				})
				idl.IndexPathMap[ctr] = nameconv.Snake2Pascal(string(key))
				idl.IsArrayMap[ctr] = false
			case TYPE_VARIABLE:
				isRaw, isFixed, _ := GetRawTypeInfo(string(value))
				if isRaw {
					CodeType = string(value)
				} else {
					CodeType = nameconv.Snake2Pascal(string(value))
				}

				t.Fields = append(t.Fields, IDLField{
					FieldType: "var",
					Key:       string(key),
					Type:      string(value),
					Index:     ctr,
					CodeType:  CodeType,
					IsRawType: isRaw,
					IsFixed:   isFixed,
				})
				idl.IndexPathMap[ctr] = nameconv.Snake2Pascal(string(key))
				idl.IsArrayMap[ctr] = false
			}

			return nil
		})

		idl.Types = append(idl.Types, t)
		return nil
	}, "$types")
	idl.Ctr = ctr
	if err != nil {
		return nil, err
	}

	return idl, nil
}
