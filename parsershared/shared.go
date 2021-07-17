package main

import "C"

import (
	"encoding/json"

	"github.com/unsafe-risk/go-alrescha/parser"
)

func main() {

}

//export ParseIDL
func ParseIDL(data []byte) (output string, returnCode int) {
	info, err := parser.ParseGenerateInfo(data)
	if err != nil {
		return "", -2
	}
	GenDatas, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return "", -1
	}
	return string(GenDatas), 0
}
