// Code generated by qtc from "unmarshal.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line goqtpl/unmarshal.qtpl:1
package goqtpl

//line goqtpl/unmarshal.qtpl:1
import "github.com/unsafe-risk/go-alrescha/parser"

//line goqtpl/unmarshal.qtpl:2
import "github.com/unsafe-risk/go-alrescha/nameconv"

//line goqtpl/unmarshal.qtpl:3
import "fmt"

//line goqtpl/unmarshal.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line goqtpl/unmarshal.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line goqtpl/unmarshal.qtpl:5
func StreamMakeUnmarshal(qw422016 *qt422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/unmarshal.qtpl:5
	qw422016.N().S(`
func (v *`)
//line goqtpl/unmarshal.qtpl:6
	qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:6
	qw422016.N().S(`) rf(r io.Reader) error {
    var err error
    var n int
    _ = n
    `)
//line goqtpl/unmarshal.qtpl:10
	if sumSize(fields) > 0 {
//line goqtpl/unmarshal.qtpl:10
		qw422016.N().S(`
    var staticBuffer [`)
//line goqtpl/unmarshal.qtpl:11
		qw422016.N().D(sumSize(fields))
//line goqtpl/unmarshal.qtpl:11
		qw422016.N().S(`]byte
    _, err = io.ReadAtLeast(r, staticBuffer[:], `)
//line goqtpl/unmarshal.qtpl:12
		qw422016.N().D(sumSize(fields))
//line goqtpl/unmarshal.qtpl:12
		qw422016.N().S(`)
    if err != nil {
		return err
	}
    `)
//line goqtpl/unmarshal.qtpl:16
	}
//line goqtpl/unmarshal.qtpl:16
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:17
	for i, field := range fields {
//line goqtpl/unmarshal.qtpl:17
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:18
		if field.IsFixed && !field.IsList {
//line goqtpl/unmarshal.qtpl:18
			qw422016.N().S(`
    // `)
//line goqtpl/unmarshal.qtpl:19
			qw422016.N().D(i)
//line goqtpl/unmarshal.qtpl:19
			qw422016.N().S(` : `)
//line goqtpl/unmarshal.qtpl:19
			qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:19
			qw422016.E().S(field.Path)
//line goqtpl/unmarshal.qtpl:19
			qw422016.N().S(`
    // Type : `)
//line goqtpl/unmarshal.qtpl:20
			qw422016.E().S(field.RawType)
//line goqtpl/unmarshal.qtpl:20
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:21
			if field.IsFixed {
//line goqtpl/unmarshal.qtpl:21
				qw422016.N().S(`// Size : `)
//line goqtpl/unmarshal.qtpl:21
				qw422016.N().D(field.Size)
//line goqtpl/unmarshal.qtpl:21
			} else {
//line goqtpl/unmarshal.qtpl:21
				qw422016.N().S(`// Size : Variable`)
//line goqtpl/unmarshal.qtpl:21
			}
//line goqtpl/unmarshal.qtpl:21
			qw422016.N().S(`
    // Offset : `)
//line goqtpl/unmarshal.qtpl:22
			qw422016.N().D(field.Offset)
//line goqtpl/unmarshal.qtpl:22
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:23
			StreamDeSerializeStatic(qw422016, "v"+field.Path, field.RawType, field.Offset, field.Size)
//line goqtpl/unmarshal.qtpl:23
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:24
		}
//line goqtpl/unmarshal.qtpl:24
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:25
	}
//line goqtpl/unmarshal.qtpl:25
	qw422016.N().S(`

    `)
//line goqtpl/unmarshal.qtpl:27
	for i, field := range fields {
//line goqtpl/unmarshal.qtpl:27
		qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:28
		if field.IsList {
//line goqtpl/unmarshal.qtpl:28
			qw422016.N().S(`
	// `)
//line goqtpl/unmarshal.qtpl:29
			qw422016.N().D(i)
//line goqtpl/unmarshal.qtpl:29
			qw422016.N().S(` : `)
//line goqtpl/unmarshal.qtpl:29
			qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:29
			qw422016.E().S(field.Path)
//line goqtpl/unmarshal.qtpl:29
			qw422016.N().S(`
    // Type : `)
//line goqtpl/unmarshal.qtpl:30
			qw422016.E().S(field.RawType)
//line goqtpl/unmarshal.qtpl:30
			qw422016.N().S(` (List)
	// Size : Variable
	`)
//line goqtpl/unmarshal.qtpl:32
			StreamDeSerializeList(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/unmarshal.qtpl:32
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:33
		}
//line goqtpl/unmarshal.qtpl:33
		qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:34
		if !field.IsFixed {
//line goqtpl/unmarshal.qtpl:34
			qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:35
			if isBytes(field.RawType) && !field.IsList {
//line goqtpl/unmarshal.qtpl:35
				qw422016.N().S(`
	// `)
//line goqtpl/unmarshal.qtpl:36
				qw422016.N().D(i)
//line goqtpl/unmarshal.qtpl:36
				qw422016.N().S(` : `)
//line goqtpl/unmarshal.qtpl:36
				qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:36
				qw422016.E().S(field.Path)
//line goqtpl/unmarshal.qtpl:36
				qw422016.N().S(`
    // Type : `)
//line goqtpl/unmarshal.qtpl:37
				qw422016.E().S(field.RawType)
//line goqtpl/unmarshal.qtpl:37
				qw422016.N().S(`
    // Size : Variable
	`)
//line goqtpl/unmarshal.qtpl:39
				StreamDeSerializeBytes(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/unmarshal.qtpl:39
				qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:40
			}
//line goqtpl/unmarshal.qtpl:40
			qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:41
		}
//line goqtpl/unmarshal.qtpl:41
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:42
	}
//line goqtpl/unmarshal.qtpl:42
	qw422016.N().S(`
    return err
}
`)
//line goqtpl/unmarshal.qtpl:45
}

//line goqtpl/unmarshal.qtpl:45
func WriteMakeUnmarshal(qq422016 qtio422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/unmarshal.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:45
	StreamMakeUnmarshal(qw422016, name, fields)
//line goqtpl/unmarshal.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:45
}

//line goqtpl/unmarshal.qtpl:45
func MakeUnmarshal(name string, fields []*parser.GenerateField) string {
//line goqtpl/unmarshal.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:45
	WriteMakeUnmarshal(qb422016, name, fields)
//line goqtpl/unmarshal.qtpl:45
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:45
	return qs422016
//line goqtpl/unmarshal.qtpl:45
}

//line goqtpl/unmarshal.qtpl:47
func StreamDeSerializeList(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:47
	qw422016.N().S(`
	var v`)
//line goqtpl/unmarshal.qtpl:48
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:48
	qw422016.N().S(` uint32
	`)
//line goqtpl/unmarshal.qtpl:49
	StreamDeSerializeInteger(qw422016, fmt.Sprintf("v%d", xh(VarName+RawType+"DeSerializeList")), "u32", 4)
//line goqtpl/unmarshal.qtpl:49
	qw422016.N().S(`
    if v`)
//line goqtpl/unmarshal.qtpl:50
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:50
	qw422016.N().S(` > ALRESCHA_MAX_LISTS_ITEM {
        return ErrSizeExceeded
    }
	// == List ==
	// List VarName : `)
//line goqtpl/unmarshal.qtpl:54
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:54
	qw422016.N().S(`
	// List VarType : `)
//line goqtpl/unmarshal.qtpl:55
	qw422016.E().S(RawType)
//line goqtpl/unmarshal.qtpl:55
	qw422016.N().S(`
    v`)
//line goqtpl/unmarshal.qtpl:56
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList+intvar"))
//line goqtpl/unmarshal.qtpl:56
	qw422016.N().S(` := int(v`)
//line goqtpl/unmarshal.qtpl:56
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:56
	qw422016.N().S(`)
    if cap(`)
//line goqtpl/unmarshal.qtpl:57
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:57
	qw422016.N().S(`) < v`)
//line goqtpl/unmarshal.qtpl:57
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList+intvar"))
//line goqtpl/unmarshal.qtpl:57
	qw422016.N().S(` {
        `)
//line goqtpl/unmarshal.qtpl:58
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:58
	qw422016.N().S(` = make([]`)
//line goqtpl/unmarshal.qtpl:58
	if ConvertToGoType(RawType) == RawType {
//line goqtpl/unmarshal.qtpl:58
		qw422016.N().S(nameconv.Snake2Pascal(ConvertToGoType(RawType)))
//line goqtpl/unmarshal.qtpl:58
	} else {
//line goqtpl/unmarshal.qtpl:58
		qw422016.N().S(ConvertToGoType(RawType))
//line goqtpl/unmarshal.qtpl:58
	}
//line goqtpl/unmarshal.qtpl:58
	qw422016.N().S(`, v`)
//line goqtpl/unmarshal.qtpl:58
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:58
	qw422016.N().S(`)
    } else {
        `)
//line goqtpl/unmarshal.qtpl:60
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:60
	qw422016.N().S(` = `)
//line goqtpl/unmarshal.qtpl:60
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:60
	qw422016.N().S(`[:v`)
//line goqtpl/unmarshal.qtpl:60
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList+intvar"))
//line goqtpl/unmarshal.qtpl:60
	qw422016.N().S(`]
    }
	for i := 0; i < v`)
//line goqtpl/unmarshal.qtpl:62
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList+intvar"))
//line goqtpl/unmarshal.qtpl:62
	qw422016.N().S(`; i++ {
		`)
//line goqtpl/unmarshal.qtpl:63
	if isBytes(RawType) {
//line goqtpl/unmarshal.qtpl:63
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:64
		StreamDeSerializeBytes(qw422016, VarName+"[i]", RawType)
//line goqtpl/unmarshal.qtpl:64
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:65
	} else if isInteger(RawType) {
//line goqtpl/unmarshal.qtpl:65
		qw422016.N().S(`
        `)
//line goqtpl/unmarshal.qtpl:66
		StreamDeSerializeInteger(qw422016, VarName+"[i]", RawType, varSize(RawType))
//line goqtpl/unmarshal.qtpl:66
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:67
	} else if isFloat(RawType) {
//line goqtpl/unmarshal.qtpl:67
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:68
		StreamDeSerializeFloat(qw422016, VarName+"[i]", RawType, varSize(RawType))
//line goqtpl/unmarshal.qtpl:68
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:69
	} else {
//line goqtpl/unmarshal.qtpl:69
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:70
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:70
		qw422016.N().S(`[i].rf(r)
		`)
//line goqtpl/unmarshal.qtpl:71
	}
//line goqtpl/unmarshal.qtpl:71
	qw422016.N().S(`
	}
`)
//line goqtpl/unmarshal.qtpl:73
}

//line goqtpl/unmarshal.qtpl:73
func WriteDeSerializeList(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:73
	StreamDeSerializeList(qw422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:73
}

//line goqtpl/unmarshal.qtpl:73
func DeSerializeList(VarName string, RawType string) string {
//line goqtpl/unmarshal.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:73
	WriteDeSerializeList(qb422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:73
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:73
	return qs422016
//line goqtpl/unmarshal.qtpl:73
}

//line goqtpl/unmarshal.qtpl:75
func StreamDeSerializeBytes(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:75
	qw422016.N().S(`
	var v`)
//line goqtpl/unmarshal.qtpl:76
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:76
	qw422016.N().S(` uint32
	`)
//line goqtpl/unmarshal.qtpl:77
	StreamDeSerializeInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeBytes")), "u32", 4)
//line goqtpl/unmarshal.qtpl:77
	qw422016.N().S(`

	`)
//line goqtpl/unmarshal.qtpl:79
	if RawType == "bytes" {
//line goqtpl/unmarshal.qtpl:79
		qw422016.N().S(`
        if v`)
//line goqtpl/unmarshal.qtpl:80
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:80
		qw422016.N().S(` > ALRESCHA_MAX_BYTES_SIZE {
            return ErrSizeExceeded
        }
        if cap(`)
//line goqtpl/unmarshal.qtpl:83
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:83
		qw422016.N().S(`) < int(v`)
//line goqtpl/unmarshal.qtpl:83
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:83
		qw422016.N().S(`) {
            `)
//line goqtpl/unmarshal.qtpl:84
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:84
		qw422016.N().S(` = make([]byte, v`)
//line goqtpl/unmarshal.qtpl:84
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:84
		qw422016.N().S(`)
        } else {
            `)
//line goqtpl/unmarshal.qtpl:86
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:86
		qw422016.N().S(` = `)
//line goqtpl/unmarshal.qtpl:86
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:86
		qw422016.N().S(`[:v`)
//line goqtpl/unmarshal.qtpl:86
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:86
		qw422016.N().S(`]
        }
		_, err = io.ReadAtLeast(r, `)
//line goqtpl/unmarshal.qtpl:88
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:88
		qw422016.N().S(`, int(v`)
//line goqtpl/unmarshal.qtpl:88
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:88
		qw422016.N().S(`))
        if err != nil {
            return err
        }
	`)
//line goqtpl/unmarshal.qtpl:92
	} else if RawType == "str" {
//line goqtpl/unmarshal.qtpl:92
		qw422016.N().S(`
        if v`)
//line goqtpl/unmarshal.qtpl:93
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:93
		qw422016.N().S(` > ALRESCHA_MAX_STRING_SIZE {
            return ErrSizeExceeded
        }
        var Buffer`)
//line goqtpl/unmarshal.qtpl:96
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:96
		qw422016.N().S(` []byte = make([]byte, v`)
//line goqtpl/unmarshal.qtpl:96
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:96
		qw422016.N().S(`)
		_, err = io.ReadAtLeast(r, Buffer`)
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().S(`, int(v`)
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().S(`))
        if err != nil {
            return err
        }
        `)
//line goqtpl/unmarshal.qtpl:101
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:101
		qw422016.N().S(` = string(Buffer`)
//line goqtpl/unmarshal.qtpl:101
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:101
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:102
	}
//line goqtpl/unmarshal.qtpl:102
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:103
}

//line goqtpl/unmarshal.qtpl:103
func WriteDeSerializeBytes(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:103
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:103
	StreamDeSerializeBytes(qw422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:103
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:103
}

//line goqtpl/unmarshal.qtpl:103
func DeSerializeBytes(VarName string, RawType string) string {
//line goqtpl/unmarshal.qtpl:103
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:103
	WriteDeSerializeBytes(qb422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:103
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:103
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:103
	return qs422016
//line goqtpl/unmarshal.qtpl:103
}

//line goqtpl/unmarshal.qtpl:105
func StreamDeSerializeStatic(qw422016 *qt422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:105
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:106
	if isInteger(RawType) {
//line goqtpl/unmarshal.qtpl:106
		qw422016.N().S(`
        `)
//line goqtpl/unmarshal.qtpl:107
		StreamDeSerializeIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:107
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:108
	} else if isFloat(RawType) {
//line goqtpl/unmarshal.qtpl:108
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:109
		StreamDeSerializeFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:109
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:110
	} else {
//line goqtpl/unmarshal.qtpl:110
		qw422016.N().S(`
	// Failed to serialize `)
//line goqtpl/unmarshal.qtpl:111
		qw422016.E().S(RawType)
//line goqtpl/unmarshal.qtpl:111
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:112
	}
//line goqtpl/unmarshal.qtpl:112
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:113
}

//line goqtpl/unmarshal.qtpl:113
func WriteDeSerializeStatic(qq422016 qtio422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:113
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:113
	StreamDeSerializeStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:113
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:113
}

//line goqtpl/unmarshal.qtpl:113
func DeSerializeStatic(VarName, RawType string, Offset, Size int) string {
//line goqtpl/unmarshal.qtpl:113
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:113
	WriteDeSerializeStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:113
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:113
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:113
	return qs422016
//line goqtpl/unmarshal.qtpl:113
}

//line goqtpl/unmarshal.qtpl:115
func StreamDeSerializeInteger(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:115
	qw422016.N().S(`
    // Size : `)
//line goqtpl/unmarshal.qtpl:116
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:116
	qw422016.N().S(`, VarName : `)
//line goqtpl/unmarshal.qtpl:116
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:116
	qw422016.N().S(`
    var v`)
//line goqtpl/unmarshal.qtpl:117
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:117
	qw422016.N().S(` `)
//line goqtpl/unmarshal.qtpl:117
	qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:117
	qw422016.N().S(`
    var Buffer`)
//line goqtpl/unmarshal.qtpl:118
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:118
	qw422016.N().S(` [`)
//line goqtpl/unmarshal.qtpl:118
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:118
	qw422016.N().S(`]byte
    _, err = io.ReadAtLeast(r, Buffer`)
//line goqtpl/unmarshal.qtpl:119
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:119
	qw422016.N().S(`[:], `)
//line goqtpl/unmarshal.qtpl:119
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:119
	qw422016.N().S(`)
    if err != nil {
		return err
	}
    `)
//line goqtpl/unmarshal.qtpl:123
	for i := 0; i < Size; i++ {
//line goqtpl/unmarshal.qtpl:123
		qw422016.N().S(`
    v`)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().S(` |= `)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().S(`(Buffer`)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().S(`[`)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().D(Size - 1 - i)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().S(`]) << `)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().D(8 * i)
//line goqtpl/unmarshal.qtpl:124
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:125
	}
//line goqtpl/unmarshal.qtpl:125
	qw422016.N().S(`

    `)
//line goqtpl/unmarshal.qtpl:127
	if isUint(RawType) {
//line goqtpl/unmarshal.qtpl:127
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:128
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:128
		qw422016.N().S(` = v`)
//line goqtpl/unmarshal.qtpl:128
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:128
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:129
	} else {
//line goqtpl/unmarshal.qtpl:129
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:130
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:130
		qw422016.N().S(` = `)
//line goqtpl/unmarshal.qtpl:130
		qw422016.N().S(ConvertToGoType(RawType))
//line goqtpl/unmarshal.qtpl:130
		qw422016.N().S(`(v`)
//line goqtpl/unmarshal.qtpl:130
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:130
		qw422016.N().S(`)
    `)
//line goqtpl/unmarshal.qtpl:131
	}
//line goqtpl/unmarshal.qtpl:131
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:132
}

//line goqtpl/unmarshal.qtpl:132
func WriteDeSerializeInteger(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:132
	StreamDeSerializeInteger(qw422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:132
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:132
}

//line goqtpl/unmarshal.qtpl:132
func DeSerializeInteger(VarName string, RawType string, Size int) string {
//line goqtpl/unmarshal.qtpl:132
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:132
	WriteDeSerializeInteger(qb422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:132
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:132
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:132
	return qs422016
//line goqtpl/unmarshal.qtpl:132
}

//line goqtpl/unmarshal.qtpl:134
func StreamDeSerializeFloat(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:134
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:135
	if Size == 4 {
//line goqtpl/unmarshal.qtpl:135
		qw422016.N().S(`
        var v`)
//line goqtpl/unmarshal.qtpl:136
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:136
		qw422016.N().S(` uint32
		`)
//line goqtpl/unmarshal.qtpl:137
		StreamDeSerializeInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloat")), "u32", Size)
//line goqtpl/unmarshal.qtpl:137
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:138
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:138
		qw422016.N().S(` = math.Float32frombits(v`)
//line goqtpl/unmarshal.qtpl:138
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:138
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:139
	} else if Size == 8 {
//line goqtpl/unmarshal.qtpl:139
		qw422016.N().S(`
		var v`)
//line goqtpl/unmarshal.qtpl:140
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:140
		qw422016.N().S(` uint64
		`)
//line goqtpl/unmarshal.qtpl:141
		StreamDeSerializeInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloat")), "u64", Size)
//line goqtpl/unmarshal.qtpl:141
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:142
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:142
		qw422016.N().S(` = math.Float64frombits(v`)
//line goqtpl/unmarshal.qtpl:142
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:142
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:143
	}
//line goqtpl/unmarshal.qtpl:143
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:144
}

//line goqtpl/unmarshal.qtpl:144
func WriteDeSerializeFloat(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:144
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:144
	StreamDeSerializeFloat(qw422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:144
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:144
}

//line goqtpl/unmarshal.qtpl:144
func DeSerializeFloat(VarName string, RawType string, Size int) string {
//line goqtpl/unmarshal.qtpl:144
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:144
	WriteDeSerializeFloat(qb422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:144
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:144
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:144
	return qs422016
//line goqtpl/unmarshal.qtpl:144
}

//line goqtpl/unmarshal.qtpl:146
func StreamDeSerializeIntegerStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:146
	qw422016.N().S(`
    // Size : `)
//line goqtpl/unmarshal.qtpl:147
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:147
	qw422016.N().S(`, Offset : `)
//line goqtpl/unmarshal.qtpl:147
	qw422016.N().D(Offset)
//line goqtpl/unmarshal.qtpl:147
	qw422016.N().S(`, VarName : `)
//line goqtpl/unmarshal.qtpl:147
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:147
	qw422016.N().S(`
    var v`)
//line goqtpl/unmarshal.qtpl:148
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:148
	qw422016.N().S(` `)
//line goqtpl/unmarshal.qtpl:148
	qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:148
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:149
	for i := 0; i < Size; i++ {
//line goqtpl/unmarshal.qtpl:149
		qw422016.N().S(`
    v`)
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().S(` |= `)
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().S(`(staticBuffer[`)
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().D(Offset + Size - 1 - i)
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().S(`]) << `)
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().D(8 * i)
//line goqtpl/unmarshal.qtpl:150
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:151
	}
//line goqtpl/unmarshal.qtpl:151
	qw422016.N().S(`

    `)
//line goqtpl/unmarshal.qtpl:153
	if isUint(RawType) {
//line goqtpl/unmarshal.qtpl:153
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:154
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:154
		qw422016.N().S(` = v`)
//line goqtpl/unmarshal.qtpl:154
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:154
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:155
	} else {
//line goqtpl/unmarshal.qtpl:155
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:156
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:156
		qw422016.N().S(` = `)
//line goqtpl/unmarshal.qtpl:156
		qw422016.N().S(ConvertToGoType(RawType))
//line goqtpl/unmarshal.qtpl:156
		qw422016.N().S(`(v`)
//line goqtpl/unmarshal.qtpl:156
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:156
		qw422016.N().S(`)
    `)
//line goqtpl/unmarshal.qtpl:157
	}
//line goqtpl/unmarshal.qtpl:157
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:158
}

//line goqtpl/unmarshal.qtpl:158
func WriteDeSerializeIntegerStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:158
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:158
	StreamDeSerializeIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:158
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:158
}

//line goqtpl/unmarshal.qtpl:158
func DeSerializeIntegerStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/unmarshal.qtpl:158
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:158
	WriteDeSerializeIntegerStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:158
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:158
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:158
	return qs422016
//line goqtpl/unmarshal.qtpl:158
}

//line goqtpl/unmarshal.qtpl:160
func StreamDeSerializeFloatStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:160
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:161
	if Size == 4 {
//line goqtpl/unmarshal.qtpl:161
		qw422016.N().S(`
        var v`)
//line goqtpl/unmarshal.qtpl:162
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:162
		qw422016.N().S(` uint32
		`)
//line goqtpl/unmarshal.qtpl:163
		StreamDeSerializeIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloatStatic")), "u32", Offset, Size)
//line goqtpl/unmarshal.qtpl:163
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:164
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:164
		qw422016.N().S(` = math.Float32frombits(v`)
//line goqtpl/unmarshal.qtpl:164
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:164
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:165
	} else if Size == 8 {
//line goqtpl/unmarshal.qtpl:165
		qw422016.N().S(`
		var v`)
//line goqtpl/unmarshal.qtpl:166
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:166
		qw422016.N().S(` uint64
		`)
//line goqtpl/unmarshal.qtpl:167
		StreamDeSerializeIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloatStatic")), "u64", Offset, Size)
//line goqtpl/unmarshal.qtpl:167
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:168
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:168
		qw422016.N().S(` = math.Float64frombits(v`)
//line goqtpl/unmarshal.qtpl:168
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:168
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:169
	}
//line goqtpl/unmarshal.qtpl:169
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:170
}

//line goqtpl/unmarshal.qtpl:170
func WriteDeSerializeFloatStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:170
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:170
	StreamDeSerializeFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:170
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:170
}

//line goqtpl/unmarshal.qtpl:170
func DeSerializeFloatStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/unmarshal.qtpl:170
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:170
	WriteDeSerializeFloatStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:170
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:170
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:170
	return qs422016
//line goqtpl/unmarshal.qtpl:170
}
