//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=templates

package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/buger/jsonparser"
	"github.com/unsafe-risk/go-alrescha/nameconv"
)

const fileName = "test.json"

type IDLFile struct {
	Version int

	Types []Type

	idxPathMap map[int]string
}

func (f *IDLFile) GetType(t string) *Type {
	for i := range f.Types {
		if f.Types[i].Name == t {
			return &f.Types[i]
		}
	}
	return nil
}

type Type struct {
	Name string

	Fields []Field

	idx int
}

type Field struct {
	Key  string
	Type string

	idx int
}

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	idl, err := ParseIDL(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(idl)
	for _, t := range idl.Types {
		var fields []*GenerateField
		for _, f := range t.Fields {
			TraceType(idl, &fields, f, nil)
		}
		sort.Sort(GenTypes(fields))
		fmt.Println(fields)
	}
}

func ParseIDL(data []byte) (*IDLFile, error) {
	idl := new(IDLFile)
	idl.idxPathMap = make(map[int]string)
	ctr := 0
	ver, err := jsonparser.GetInt(data, "$version")
	if err != nil {
		return nil, err
	}
	idl.Version = int(ver)

	jsonparser.ObjectEach(data, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
		ctr++
		t := Type{}
		t.Name = string(key)
		t.idx = ctr
		idl.idxPathMap[ctr] = nameconv.Snake2Pascal(string(key))

		jsonparser.ObjectEach(value, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
			ctr++
			t.Fields = append(t.Fields, Field{
				Key:  string(key),
				Type: string(value),
				idx:  ctr,
			})

			idl.idxPathMap[ctr] = nameconv.Snake2Pascal(string(key))
			return nil
		})

		idl.Types = append(idl.Types, t)
		return nil
	}, "$types")

	return idl, nil
}
