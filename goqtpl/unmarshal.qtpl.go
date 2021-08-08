// Code generated by qtc from "unmarshal.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line goqtpl/unmarshal.qtpl:1
package goqtpl

//line goqtpl/unmarshal.qtpl:1
import "github.com/unsafe-risk/go-alrescha/parser"

//line goqtpl/unmarshal.qtpl:2
import "fmt"

//line goqtpl/unmarshal.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line goqtpl/unmarshal.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line goqtpl/unmarshal.qtpl:4
func StreamMakeUnmarshal(qw422016 *qt422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/unmarshal.qtpl:4
	qw422016.N().S(`
func (v *`)
//line goqtpl/unmarshal.qtpl:5
	qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:5
	qw422016.N().S(`) rf(r io.Reader) {
    `)
//line goqtpl/unmarshal.qtpl:6
	if sumSize(fields) > 0 {
//line goqtpl/unmarshal.qtpl:6
		qw422016.N().S(`
    var staticBuffer [`)
//line goqtpl/unmarshal.qtpl:7
		qw422016.N().D(sumSize(fields))
//line goqtpl/unmarshal.qtpl:7
		qw422016.N().S(`]byte
    r.Read(staticBuffer[:])
    `)
//line goqtpl/unmarshal.qtpl:9
	}
//line goqtpl/unmarshal.qtpl:9
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:10
	for i, field := range fields {
//line goqtpl/unmarshal.qtpl:10
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:11
		if field.IsFixed && !field.IsList {
//line goqtpl/unmarshal.qtpl:11
			qw422016.N().S(`
    // `)
//line goqtpl/unmarshal.qtpl:12
			qw422016.N().D(i)
//line goqtpl/unmarshal.qtpl:12
			qw422016.N().S(` : `)
//line goqtpl/unmarshal.qtpl:12
			qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:12
			qw422016.E().S(field.Path)
//line goqtpl/unmarshal.qtpl:12
			qw422016.N().S(`
    // Type : `)
//line goqtpl/unmarshal.qtpl:13
			qw422016.E().S(field.RawType)
//line goqtpl/unmarshal.qtpl:13
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:14
			if field.IsFixed {
//line goqtpl/unmarshal.qtpl:14
				qw422016.N().S(`// Size : `)
//line goqtpl/unmarshal.qtpl:14
				qw422016.N().D(field.Size)
//line goqtpl/unmarshal.qtpl:14
			} else {
//line goqtpl/unmarshal.qtpl:14
				qw422016.N().S(`// Size : Variable`)
//line goqtpl/unmarshal.qtpl:14
			}
//line goqtpl/unmarshal.qtpl:14
			qw422016.N().S(`
    // Offset : `)
//line goqtpl/unmarshal.qtpl:15
			qw422016.N().D(field.Offset)
//line goqtpl/unmarshal.qtpl:15
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:16
			StreamDeSerializeStatic(qw422016, "v"+field.Path, field.RawType, field.Offset, field.Size)
//line goqtpl/unmarshal.qtpl:16
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:17
		}
//line goqtpl/unmarshal.qtpl:17
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:18
	}
//line goqtpl/unmarshal.qtpl:18
	qw422016.N().S(`

    `)
//line goqtpl/unmarshal.qtpl:20
	for i, field := range fields {
//line goqtpl/unmarshal.qtpl:20
		qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:21
		if field.IsList {
//line goqtpl/unmarshal.qtpl:21
			qw422016.N().S(`
	// `)
//line goqtpl/unmarshal.qtpl:22
			qw422016.N().D(i)
//line goqtpl/unmarshal.qtpl:22
			qw422016.N().S(` : `)
//line goqtpl/unmarshal.qtpl:22
			qw422016.E().S(name)
//line goqtpl/unmarshal.qtpl:22
			qw422016.E().S(field.Path)
//line goqtpl/unmarshal.qtpl:22
			qw422016.N().S(`
    // Type : `)
//line goqtpl/unmarshal.qtpl:23
			qw422016.E().S(field.RawType)
//line goqtpl/unmarshal.qtpl:23
			qw422016.N().S(` (List)
	// Size : Variable
	`)
//line goqtpl/unmarshal.qtpl:25
			StreamDeSerializeList(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/unmarshal.qtpl:25
			qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:26
		}
//line goqtpl/unmarshal.qtpl:26
		qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:27
		if !field.IsFixed {
//line goqtpl/unmarshal.qtpl:27
			qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:28
			if isBytes(field.RawType) && !field.IsList {
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
				qw422016.N().S(`
    // Size : Variable
	`)
//line goqtpl/unmarshal.qtpl:32
				StreamDeSerializeBytes(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/unmarshal.qtpl:32
				qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:33
			}
//line goqtpl/unmarshal.qtpl:33
			qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:34
		}
//line goqtpl/unmarshal.qtpl:34
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:35
	}
//line goqtpl/unmarshal.qtpl:35
	qw422016.N().S(`
}
`)
//line goqtpl/unmarshal.qtpl:37
}

//line goqtpl/unmarshal.qtpl:37
func WriteMakeUnmarshal(qq422016 qtio422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/unmarshal.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:37
	StreamMakeUnmarshal(qw422016, name, fields)
//line goqtpl/unmarshal.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:37
}

//line goqtpl/unmarshal.qtpl:37
func MakeUnmarshal(name string, fields []*parser.GenerateField) string {
//line goqtpl/unmarshal.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:37
	WriteMakeUnmarshal(qb422016, name, fields)
//line goqtpl/unmarshal.qtpl:37
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:37
	return qs422016
//line goqtpl/unmarshal.qtpl:37
}

//line goqtpl/unmarshal.qtpl:39
func StreamDeSerializeList(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:39
	qw422016.N().S(`
	var v`)
//line goqtpl/unmarshal.qtpl:40
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:40
	qw422016.N().S(` uint32
	`)
//line goqtpl/unmarshal.qtpl:41
	StreamDeSerializeInteger(qw422016, fmt.Sprintf("v%d", xh(VarName+RawType+"DeSerializeList")), "u32", 4)
//line goqtpl/unmarshal.qtpl:41
	qw422016.N().S(`
	// == List ==
	// List VarName : `)
//line goqtpl/unmarshal.qtpl:43
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:43
	qw422016.N().S(`
	// List VarType : `)
//line goqtpl/unmarshal.qtpl:44
	qw422016.E().S(RawType)
//line goqtpl/unmarshal.qtpl:44
	qw422016.N().S(`
    if len(VarName) < v`)
//line goqtpl/unmarshal.qtpl:45
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:45
	qw422016.N().S(` {
        `)
//line goqtpl/unmarshal.qtpl:46
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:46
	qw422016.N().S(` = make([]`)
//line goqtpl/unmarshal.qtpl:46
	qw422016.N().S(ConvertToGoType(RawType))
//line goqtpl/unmarshal.qtpl:46
	qw422016.N().S(`, v`)
//line goqtpl/unmarshal.qtpl:46
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:46
	qw422016.N().S(`)
    }
	for i := 0; i < v`)
//line goqtpl/unmarshal.qtpl:48
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeList"))
//line goqtpl/unmarshal.qtpl:48
	qw422016.N().S(`; i++ {
		`)
//line goqtpl/unmarshal.qtpl:49
	if isBytes(RawType) {
//line goqtpl/unmarshal.qtpl:49
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:50
		StreamDeSerializeBytes(qw422016, VarName+"[i]", RawType)
//line goqtpl/unmarshal.qtpl:50
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:51
	} else if isInteger(RawType) {
//line goqtpl/unmarshal.qtpl:51
		qw422016.N().S(`
        `)
//line goqtpl/unmarshal.qtpl:52
		StreamDeSerializeInteger(qw422016, VarName+"[i]", RawType, varSize(RawType))
//line goqtpl/unmarshal.qtpl:52
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:53
	} else if isFloat(RawType) {
//line goqtpl/unmarshal.qtpl:53
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:54
		StreamDeSerializeFloat(qw422016, VarName+"[i]", RawType, varSize(RawType))
//line goqtpl/unmarshal.qtpl:54
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:55
	} else {
//line goqtpl/unmarshal.qtpl:55
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:56
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:56
		qw422016.N().S(`[i].rf(r)
		`)
//line goqtpl/unmarshal.qtpl:57
	}
//line goqtpl/unmarshal.qtpl:57
	qw422016.N().S(`
	}
`)
//line goqtpl/unmarshal.qtpl:59
}

//line goqtpl/unmarshal.qtpl:59
func WriteDeSerializeList(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:59
	StreamDeSerializeList(qw422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:59
}

//line goqtpl/unmarshal.qtpl:59
func DeSerializeList(VarName string, RawType string) string {
//line goqtpl/unmarshal.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:59
	WriteDeSerializeList(qb422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:59
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:59
	return qs422016
//line goqtpl/unmarshal.qtpl:59
}

//line goqtpl/unmarshal.qtpl:61
func StreamDeSerializeBytes(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:61
	qw422016.N().S(`
	var v`)
//line goqtpl/unmarshal.qtpl:62
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:62
	qw422016.N().S(` uint32
	`)
//line goqtpl/unmarshal.qtpl:63
	StreamDeSerializeInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeBytes")), "u32", 4)
//line goqtpl/unmarshal.qtpl:63
	qw422016.N().S(`

    var Buffer`)
//line goqtpl/unmarshal.qtpl:65
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:65
	qw422016.N().S(` []byte = make([]byte, v`)
//line goqtpl/unmarshal.qtpl:65
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:65
	qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:66
	if RawType == "bytes" {
//line goqtpl/unmarshal.qtpl:66
		qw422016.N().S(`
		r.Read(Buffer`)
//line goqtpl/unmarshal.qtpl:67
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:67
		qw422016.N().S(`)
        `)
//line goqtpl/unmarshal.qtpl:68
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:68
		qw422016.N().S(` = Buffer`)
//line goqtpl/unmarshal.qtpl:68
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:68
		qw422016.N().S(`
	`)
//line goqtpl/unmarshal.qtpl:69
	} else if RawType == "str" {
//line goqtpl/unmarshal.qtpl:69
		qw422016.N().S(`
		r.Read(Buffer`)
//line goqtpl/unmarshal.qtpl:70
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:70
		qw422016.N().S(`)
        `)
//line goqtpl/unmarshal.qtpl:71
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:71
		qw422016.N().S(` = string(Buffer`)
//line goqtpl/unmarshal.qtpl:71
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeBytes"))
//line goqtpl/unmarshal.qtpl:71
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:72
	}
//line goqtpl/unmarshal.qtpl:72
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:73
}

//line goqtpl/unmarshal.qtpl:73
func WriteDeSerializeBytes(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/unmarshal.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:73
	StreamDeSerializeBytes(qw422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:73
}

//line goqtpl/unmarshal.qtpl:73
func DeSerializeBytes(VarName string, RawType string) string {
//line goqtpl/unmarshal.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:73
	WriteDeSerializeBytes(qb422016, VarName, RawType)
//line goqtpl/unmarshal.qtpl:73
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:73
	return qs422016
//line goqtpl/unmarshal.qtpl:73
}

//line goqtpl/unmarshal.qtpl:75
func StreamDeSerializeStatic(qw422016 *qt422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:75
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:76
	if isInteger(RawType) {
//line goqtpl/unmarshal.qtpl:76
		qw422016.N().S(`
        `)
//line goqtpl/unmarshal.qtpl:77
		StreamDeSerializeIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:77
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:78
	} else if isFloat(RawType) {
//line goqtpl/unmarshal.qtpl:78
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:79
		StreamDeSerializeFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:79
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:80
	} else {
//line goqtpl/unmarshal.qtpl:80
		qw422016.N().S(`
	// Failed to serialize `)
//line goqtpl/unmarshal.qtpl:81
		qw422016.E().S(RawType)
//line goqtpl/unmarshal.qtpl:81
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:82
	}
//line goqtpl/unmarshal.qtpl:82
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:83
}

//line goqtpl/unmarshal.qtpl:83
func WriteDeSerializeStatic(qq422016 qtio422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:83
	StreamDeSerializeStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:83
}

//line goqtpl/unmarshal.qtpl:83
func DeSerializeStatic(VarName, RawType string, Offset, Size int) string {
//line goqtpl/unmarshal.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:83
	WriteDeSerializeStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:83
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:83
	return qs422016
//line goqtpl/unmarshal.qtpl:83
}

//line goqtpl/unmarshal.qtpl:85
func StreamDeSerializeInteger(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:85
	qw422016.N().S(`
    // Size : `)
//line goqtpl/unmarshal.qtpl:86
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:86
	qw422016.N().S(`, VarName : `)
//line goqtpl/unmarshal.qtpl:86
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:86
	qw422016.N().S(`
    var v`)
//line goqtpl/unmarshal.qtpl:87
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:87
	qw422016.N().S(` `)
//line goqtpl/unmarshal.qtpl:87
	qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:87
	qw422016.N().S(`
    var Buffer`)
//line goqtpl/unmarshal.qtpl:88
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:88
	qw422016.N().S(` [`)
//line goqtpl/unmarshal.qtpl:88
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:88
	qw422016.N().S(`]byte
    r.Read(Buffer`)
//line goqtpl/unmarshal.qtpl:89
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:89
	qw422016.N().S(`[:])
    `)
//line goqtpl/unmarshal.qtpl:90
	for i := 0; i < Size; i++ {
//line goqtpl/unmarshal.qtpl:90
		qw422016.N().S(`
    v`)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().S(` |= `)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().S(`(Buffer`)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().S(`[`)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().D(Size - 1 - i)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().S(`]) << `)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().D(8 * i)
//line goqtpl/unmarshal.qtpl:91
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:92
	}
//line goqtpl/unmarshal.qtpl:92
	qw422016.N().S(`

    `)
//line goqtpl/unmarshal.qtpl:94
	if isUint(RawType) {
//line goqtpl/unmarshal.qtpl:94
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:95
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:95
		qw422016.N().S(` = v`)
//line goqtpl/unmarshal.qtpl:95
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:95
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:96
	} else {
//line goqtpl/unmarshal.qtpl:96
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:97
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().S(` = `)
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().S(ConvertToGoType(RawType))
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().S(`(v`)
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeInteger"))
//line goqtpl/unmarshal.qtpl:97
		qw422016.N().S(`)
    `)
//line goqtpl/unmarshal.qtpl:98
	}
//line goqtpl/unmarshal.qtpl:98
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:99
}

//line goqtpl/unmarshal.qtpl:99
func WriteDeSerializeInteger(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:99
	StreamDeSerializeInteger(qw422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:99
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:99
}

//line goqtpl/unmarshal.qtpl:99
func DeSerializeInteger(VarName string, RawType string, Size int) string {
//line goqtpl/unmarshal.qtpl:99
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:99
	WriteDeSerializeInteger(qb422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:99
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:99
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:99
	return qs422016
//line goqtpl/unmarshal.qtpl:99
}

//line goqtpl/unmarshal.qtpl:101
func StreamDeSerializeFloat(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:101
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:102
	if Size == 4 {
//line goqtpl/unmarshal.qtpl:102
		qw422016.N().S(`
        var v`)
//line goqtpl/unmarshal.qtpl:103
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:103
		qw422016.N().S(` uint32
		`)
//line goqtpl/unmarshal.qtpl:104
		StreamDeSerializeInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloat")), "u32", Size)
//line goqtpl/unmarshal.qtpl:104
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:105
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:105
		qw422016.N().S(` = math.Float32frombits(v`)
//line goqtpl/unmarshal.qtpl:105
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:105
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:106
	} else if Size == 8 {
//line goqtpl/unmarshal.qtpl:106
		qw422016.N().S(`
		var v`)
//line goqtpl/unmarshal.qtpl:107
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:107
		qw422016.N().S(` uint64
		`)
//line goqtpl/unmarshal.qtpl:108
		StreamDeSerializeInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloat")), "u64", Size)
//line goqtpl/unmarshal.qtpl:108
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:109
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:109
		qw422016.N().S(` = math.Float64frombits(v`)
//line goqtpl/unmarshal.qtpl:109
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloat"))
//line goqtpl/unmarshal.qtpl:109
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:110
	}
//line goqtpl/unmarshal.qtpl:110
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:111
}

//line goqtpl/unmarshal.qtpl:111
func WriteDeSerializeFloat(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/unmarshal.qtpl:111
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:111
	StreamDeSerializeFloat(qw422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:111
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:111
}

//line goqtpl/unmarshal.qtpl:111
func DeSerializeFloat(VarName string, RawType string, Size int) string {
//line goqtpl/unmarshal.qtpl:111
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:111
	WriteDeSerializeFloat(qb422016, VarName, RawType, Size)
//line goqtpl/unmarshal.qtpl:111
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:111
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:111
	return qs422016
//line goqtpl/unmarshal.qtpl:111
}

//line goqtpl/unmarshal.qtpl:113
func StreamDeSerializeIntegerStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:113
	qw422016.N().S(`
    // Size : `)
//line goqtpl/unmarshal.qtpl:114
	qw422016.N().D(Size)
//line goqtpl/unmarshal.qtpl:114
	qw422016.N().S(`, Offset : `)
//line goqtpl/unmarshal.qtpl:114
	qw422016.N().D(Offset)
//line goqtpl/unmarshal.qtpl:114
	qw422016.N().S(`, VarName : `)
//line goqtpl/unmarshal.qtpl:114
	qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:114
	qw422016.N().S(`
    var v`)
//line goqtpl/unmarshal.qtpl:115
	qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:115
	qw422016.N().S(` `)
//line goqtpl/unmarshal.qtpl:115
	qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:115
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:116
	for i := 0; i < Size; i++ {
//line goqtpl/unmarshal.qtpl:116
		qw422016.N().S(`
    v`)
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().S(` |= `)
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().S(getUintType(RawType))
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().S(`(staticBuffer[`)
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().D(Offset + Size - 1 - i)
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().S(`]) << `)
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().D(8 * i)
//line goqtpl/unmarshal.qtpl:117
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:118
	}
//line goqtpl/unmarshal.qtpl:118
	qw422016.N().S(`

    `)
//line goqtpl/unmarshal.qtpl:120
	if isUint(RawType) {
//line goqtpl/unmarshal.qtpl:120
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:121
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:121
		qw422016.N().S(` = v`)
//line goqtpl/unmarshal.qtpl:121
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:121
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:122
	} else {
//line goqtpl/unmarshal.qtpl:122
		qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:123
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:123
		qw422016.N().S(` = `)
//line goqtpl/unmarshal.qtpl:123
		qw422016.N().S(ConvertToGoType(RawType))
//line goqtpl/unmarshal.qtpl:123
		qw422016.N().S(`(v`)
//line goqtpl/unmarshal.qtpl:123
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeIntegerStatic"))
//line goqtpl/unmarshal.qtpl:123
		qw422016.N().S(`)
    `)
//line goqtpl/unmarshal.qtpl:124
	}
//line goqtpl/unmarshal.qtpl:124
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:125
}

//line goqtpl/unmarshal.qtpl:125
func WriteDeSerializeIntegerStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:125
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:125
	StreamDeSerializeIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:125
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:125
}

//line goqtpl/unmarshal.qtpl:125
func DeSerializeIntegerStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/unmarshal.qtpl:125
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:125
	WriteDeSerializeIntegerStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:125
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:125
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:125
	return qs422016
//line goqtpl/unmarshal.qtpl:125
}

//line goqtpl/unmarshal.qtpl:127
func StreamDeSerializeFloatStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:127
	qw422016.N().S(`
    `)
//line goqtpl/unmarshal.qtpl:128
	if Size == 4 {
//line goqtpl/unmarshal.qtpl:128
		qw422016.N().S(`
        var v`)
//line goqtpl/unmarshal.qtpl:129
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:129
		qw422016.N().S(` uint32
		`)
//line goqtpl/unmarshal.qtpl:130
		StreamDeSerializeIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloatStatic")), "u32", Offset, Size)
//line goqtpl/unmarshal.qtpl:130
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:131
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:131
		qw422016.N().S(` = math.Float32frombits(v`)
//line goqtpl/unmarshal.qtpl:131
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:131
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:132
	} else if Size == 8 {
//line goqtpl/unmarshal.qtpl:132
		qw422016.N().S(`
		var v`)
//line goqtpl/unmarshal.qtpl:133
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:133
		qw422016.N().S(` uint64
		`)
//line goqtpl/unmarshal.qtpl:134
		StreamDeSerializeIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"DeSerializeFloatStatic")), "u64", Offset, Size)
//line goqtpl/unmarshal.qtpl:134
		qw422016.N().S(`
		`)
//line goqtpl/unmarshal.qtpl:135
		qw422016.E().S(VarName)
//line goqtpl/unmarshal.qtpl:135
		qw422016.N().S(` = math.Float64frombits(v`)
//line goqtpl/unmarshal.qtpl:135
		qw422016.N().DUL(xh(VarName + RawType + "DeSerializeFloatStatic"))
//line goqtpl/unmarshal.qtpl:135
		qw422016.N().S(`)
	`)
//line goqtpl/unmarshal.qtpl:136
	}
//line goqtpl/unmarshal.qtpl:136
	qw422016.N().S(`
`)
//line goqtpl/unmarshal.qtpl:137
}

//line goqtpl/unmarshal.qtpl:137
func WriteDeSerializeFloatStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/unmarshal.qtpl:137
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/unmarshal.qtpl:137
	StreamDeSerializeFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:137
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/unmarshal.qtpl:137
}

//line goqtpl/unmarshal.qtpl:137
func DeSerializeFloatStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/unmarshal.qtpl:137
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/unmarshal.qtpl:137
	WriteDeSerializeFloatStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/unmarshal.qtpl:137
	qs422016 := string(qb422016.B)
//line goqtpl/unmarshal.qtpl:137
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/unmarshal.qtpl:137
	return qs422016
//line goqtpl/unmarshal.qtpl:137
}
