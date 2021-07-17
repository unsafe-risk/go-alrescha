package parser

import (
	"sort"

	"github.com/buger/jsonparser"
	"github.com/unsafe-risk/go-alrescha/nameconv"
)

type IDLFile struct {
	Version int

	Types []IDLType

	IndexPathMap map[int]string
	TypeIndexMap map[string]int
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
	Key      string
	Type     string
	CodeType string

	Index int

	IsRawType bool
	IsFixed   bool
}

type GernerateInfo struct {
	Structs []*GenerateStruct
	IDL     *IDLFile
}

func ParseGenerateInfo(data []byte) (*GernerateInfo, error) {
	idl, err := ParseJSONData(data)
	if err != nil {
		return nil, err
	}
	var structs []*GenerateStruct
	for _, t := range idl.Types {
		var fields []*GenerateField
		for _, f := range t.Fields {
			TraceType(idl, &fields, f, nil)
		}
		sort.Sort(GenTypes(fields))
		structs = append(structs, &GenerateStruct{
			Name:  idl.IndexPathMap[t.Index],
			Types: fields,
		})
	}
	return &GernerateInfo{
		Structs: structs,
		IDL:     idl,
	}, nil
}

func ParseJSONData(data []byte) (*IDLFile, error) {
	idl := new(IDLFile)
	idl.IndexPathMap = make(map[int]string)
	idl.TypeIndexMap = make(map[string]int)
	ctr := 0
	ver, err := jsonparser.GetInt(data, "$version")
	if err != nil {
		return nil, err
	}
	idl.Version = int(ver)

	jsonparser.ObjectEach(data, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
		ctr++
		t := IDLType{}
		t.Name = string(key)
		t.Index = ctr
		idl.IndexPathMap[ctr] = nameconv.Snake2Pascal(string(key))
		idl.TypeIndexMap[t.Name] = ctr

		jsonparser.ObjectEach(value, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
			ctr++
			isRaw, isFixed, _ := GetRawTypeInfo(string(value))
			var CodeType string

			if isRaw {
				CodeType = string(value)
			} else {
				CodeType = nameconv.Snake2Pascal(string(value))
			}

			t.Fields = append(t.Fields, IDLField{
				Key:       string(key),
				Type:      string(value),
				Index:     ctr,
				CodeType:  CodeType,
				IsRawType: isRaw,
				IsFixed:   isFixed,
			})

			idl.IndexPathMap[ctr] = nameconv.Snake2Pascal(string(key))
			return nil
		})

		idl.Types = append(idl.Types, t)
		return nil
	}, "$types")

	return idl, nil
}
