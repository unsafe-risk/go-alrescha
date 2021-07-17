package parser

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
	Name    string
	Size    int
	Offset  int
	IsFixed bool

	RawType string

	Index int
}

func (f *GenerateField) String() string {
	return fmt.Sprintf("Field(Path=%v, Size=%d, Type=%v)", f.Path, f.Size, f.RawType)
}

type GenTypes []*GenerateField

func (g GenTypes) Len() int      { return len(g) }
func (g GenTypes) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g GenTypes) Less(i, j int) bool {
	a, b := g[i], g[j]
	if a.IsFixed && b.IsFixed {
		if a.Size == b.Size {
			return a.Index < b.Index
		}
		return a.Size > b.Size
	} else if !a.IsFixed && !b.IsFixed {
		return a.Index < b.Index
	}
	if a.IsFixed {
		return true
	}
	return false
}

var ErrTypeNotFound = errors.New("type not found")

func TraceType(f *IDLFile, to *[]*GenerateField, ff IDLField, path []int) error {
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
			gopath += "." + f.IndexPathMap[path[i]]
		}
		gf.Path = gopath + "." + f.IndexPathMap[ff.Index]
		gf.Name = f.IndexPathMap[ff.Index]
		gf.Index = ff.Index
		gf.RawType = ff.Type
		return nil
	}
	newpath := make([]int, len(path), len(path)+1)
	copy(newpath, path)
	newpath = append(newpath, ff.Index)

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
