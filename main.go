//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=templates

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buger/jsonparser"
)

const fileName = "test.json"

type IDLFile struct {
	Version int

	Types []Type
}

type Type struct {
	Name string

	Fields []Field
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

	fmt.Println(ParseIDL(data))
}

func ParseIDL(data []byte) *IDLFile {
	idl := new(IDLFile)
	ctr := 0
	ver, err := jsonparser.GetInt(data, "$version")
	if err != nil {
		log.Fatal(err)
	}

	idl.Version = int(ver)

	jsonparser.ObjectEach(data, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
		ctr++
		t := Type{}
		t.Name = string(key)

		jsonparser.ObjectEach(value, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
			ctr++
			t.Fields = append(t.Fields, Field{
				Key:  string(key),
				Type: string(value),
				idx:  ctr,
			})
			return nil
		})

		idl.Types = append(idl.Types, t)
		return nil
	}, "$types")

	return idl
}
