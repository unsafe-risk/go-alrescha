package main

import (
	"errors"
	"fmt"
)

type GenerateStruct struct {
	Name  string
	Types []*GenerateField
}

type GenerateField struct {
	Path    string
	Size    int
	IsFixed bool

	RawType string
	GoType  string
	GoName  string

	Index int
}

func (f *GenerateField) String() string {
	return fmt.Sprintf("Path=%v, Size=%d, Type=%v", f.Path, f.Size, f.RawType)
}

type GenTypes []*GenerateField

var ErrTypeNotFound = errors.New("type not found")

func TraceType(f *IDLFile, to *[]*GenerateField, ff Field, path []int) error {
	isRaw, isFixed, size := GetRawTypeInfo(ff.Type)
	if isRaw {
		gf := &GenerateField{
			Size:    size,
			IsFixed: isFixed,
		}
		*to = append(*to, gf)
		gopath := ""
		//fmt.Println(path)
		for i := range path {
			gopath += "." + f.idxPathMap[path[i]]
		}
		gf.Path = gopath + "." + f.idxPathMap[ff.idx]
		gf.GoName = f.idxPathMap[ff.idx]
		gf.Index = ff.idx
		gf.RawType = ff.Type
		return nil
	}
	newpath := make([]int, len(path))
	copy(newpath, path)
	newpath = append(newpath, ff.idx)

	t := f.GetType(ff.Type)
	if t == nil {
		return nil
	}

	for i := range t.Fields {
		err := TraceType(f, to, t.Fields[i], newpath)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRawTypeInfo(t string) (isRaw bool, isFixed bool, Size int) {
	switch t {
	case "i8":
		return true, true, 1
	case "i16":
		return true, true, 2
	case "i32":
		return true, true, 4
	case "i64":
		return true, true, 8
	case "u8":
		return true, true, 1
	case "u16":
		return true, true, 2
	case "u32":
		return true, true, 4
	case "u64":
		return true, true, 8
	case "f32":
		return true, true, 4
	case "f64":
		return true, true, 8
	case "str":
		return true, false, 0
	case "bytes":
		return true, false, 0
	}
	return false, false, 0
}
