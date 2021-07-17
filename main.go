package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/unsafe-risk/go-alrescha/goqtpl"
	"github.com/unsafe-risk/go-alrescha/parser"
)

const InputFileName = "test.json"
const OutputFileName = "test.out.json"

func main() {
	data, err := os.ReadFile(InputFileName)
	if err != nil {
		log.Fatal(err)
	}
	info, err := parser.ParseGenerateInfo(data)
	if err != nil {
		log.Fatal(err)
	}
	GenDatas, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(OutputFileName, GenDatas, 0644)

	log.Println("Generate info file:", OutputFileName)

	for i := range info.IDL.Types {
		qtplStructs := make([]goqtpl.StructField, 0, len(info.IDL.Types))
		for j := range info.IDL.Types[i].Fields {
			qtplStructs = append(qtplStructs, goqtpl.StructField{
				Name: info.IDL.IndexPathMap[info.IDL.Types[i].Fields[j].Index],
				Type: ConvertToGoType(info.IDL.Types[i].Fields[j].CodeType),
			})
		}
		fmt.Println(string(goqtpl.GenerateGoStruct(info.IDL.IndexPathMap[info.IDL.Types[i].Index], qtplStructs)))
	}
	for i := range info.Structs {
		fmt.Println(string(goqtpl.MakeMarshal(info.Structs[i].Name, info.Structs[i].Types)))
	}
}

func ConvertToGoType(RawType string) string {
	switch RawType {
	case "i8":
		return "int8"
	case "i16":
		return "int16"
	case "i32":
		return "int32"
	case "i64":
		return "int64"
	case "u8":
		return "uint8"
	case "u16":
		return "uint16"
	case "u32":
		return "uint32"
	case "u64":
		return "uint64"
	case "f32":
		return "float32"
	case "f64":
		return "float64"
	case "str":
		return "string"
	case "bytes":
		return "[]byte"
	}
	return RawType
}
