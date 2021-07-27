//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=goqtpl

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

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
	fmt.Println("package main")
	fmt.Println()
	fmt.Println("import (")
	fmt.Println("\t\"io\"")
	fmt.Println("\t\"math\"")
	fmt.Println(")")
	fmt.Println()
	fmt.Println("var _ = math.Float64bits")
	fmt.Println("var _ = io.EOF")

	for i := range info.IDL.Types {
		qtplStructs := make([]goqtpl.StructField, 0, len(info.IDL.Types))
		for j := range info.IDL.Types[i].Fields {
			if info.IDL.Types[i].Fields[j].FieldType == "array" {
				qtplStructs = append(qtplStructs, goqtpl.StructField{
					Name: info.IDL.IndexPathMap[info.IDL.Types[i].Fields[j].Index],
					Type: "[" + strconv.Itoa(info.IDL.Types[i].Fields[j].ArrayLength) + "]" + ConvertToGoType(info.IDL.Types[i].Fields[j].CodeType),
				})
			} else if info.IDL.Types[i].Fields[j].FieldType == "list" {
				qtplStructs = append(qtplStructs, goqtpl.StructField{
					Name: info.IDL.IndexPathMap[info.IDL.Types[i].Fields[j].Index],
					Type: "[]" + ConvertToGoType(info.IDL.Types[i].Fields[j].CodeType),
				})
			} else {
				qtplStructs = append(qtplStructs, goqtpl.StructField{
					Name: info.IDL.IndexPathMap[info.IDL.Types[i].Fields[j].Index],
					Type: ConvertToGoType(info.IDL.Types[i].Fields[j].CodeType),
				})
			}
		}
		fmt.Println(CleanCode(goqtpl.GenerateGoStruct(info.IDL.IndexPathMap[info.IDL.Types[i].Index], qtplStructs)))
	}
	for i := range info.Structs {
		fmt.Println(CleanCode(goqtpl.MakeMarshal(info.Structs[i].Name, info.Structs[i].Types)))
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

func CleanCode(in string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(in)
	outbuf := bytes.NewBuffer(nil)
	gfmt := exec.Command("gofmt")
	gfmt.Stdin = buf
	gfmt.Stdout = outbuf
	gfmt.Start()
	gfmt.Wait()
	return outbuf.String()
}
