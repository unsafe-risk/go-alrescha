package main

import (
	"encoding/json"
	"log"
	"os"

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
}
