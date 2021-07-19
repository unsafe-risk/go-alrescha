// Code generated by qtc from "marshal.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line goqtpl/marshal.qtpl:1
package goqtpl

//line goqtpl/marshal.qtpl:1
import "github.com/unsafe-risk/go-alrescha/parser"

//line goqtpl/marshal.qtpl:2
import "github.com/cespare/xxhash"

//line goqtpl/marshal.qtpl:3
import "fmt"

//line goqtpl/marshal.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line goqtpl/marshal.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line goqtpl/marshal.qtpl:6
func xh(v interface{}) uint64 {
	return xxhash.Sum64String(fmt.Sprintf("%v", v))
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

func IntInfo(RawType string) (isInteger bool, isUint bool, uIntType string) {
	switch RawType {
	case "i8":
		return true, false, "uint8"
	case "i16":
		return true, false, "uint16"
	case "i32":
		return true, false, "uint32"
	case "i64":
		return true, false, "uint64"
	case "u8":
		return true, true, "uint8"
	case "u16":
		return true, true, "uint16"
	case "u32":
		return true, true, "uint32"
	case "u64":
		return true, true, "uint64"
	}
	return false, false, RawType
}

func isInteger(RawType string) bool {
	isInteger, _, _ := IntInfo(RawType)
	return isInteger
}

func isUint(RawType string) bool {
	_, isuint, _ := IntInfo(RawType)
	return isuint
}

func isFloat(RawType string) bool {
	switch RawType {
	case "f32":
		return true
	case "f64":
		return true
	}
	return false
}

func isBytes(RawType string) bool {
	switch RawType {
	case "bytes":
		return true
	case "str":
		return true
	}
	return false
}

func getUintType(RawType string) string {
	_, _, uIntType := IntInfo(RawType)
	return uIntType
}

func sumSize(fields []*parser.GenerateField) int {
	var size int
	for i := range fields {
		if fields[i].IsFixed {
			size += fields[i].Size
		}
	}
	return size
}

//line goqtpl/marshal.qtpl:108
func StreamSerializeStatic(qw422016 *qt422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:108
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:109
	if isInteger(RawType) {
//line goqtpl/marshal.qtpl:109
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:110
		StreamIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:110
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:111
	} else if isFloat(RawType) {
//line goqtpl/marshal.qtpl:111
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:112
		StreamFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:112
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:113
	} else {
//line goqtpl/marshal.qtpl:113
		qw422016.N().S(`
	// Failed to serialize `)
//line goqtpl/marshal.qtpl:114
		qw422016.E().S(RawType)
//line goqtpl/marshal.qtpl:114
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:115
	}
//line goqtpl/marshal.qtpl:115
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:116
}

//line goqtpl/marshal.qtpl:116
func WriteSerializeStatic(qq422016 qtio422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:116
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:116
	StreamSerializeStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:116
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:116
}

//line goqtpl/marshal.qtpl:116
func SerializeStatic(VarName, RawType string, Offset, Size int) string {
//line goqtpl/marshal.qtpl:116
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:116
	WriteSerializeStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:116
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:116
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:116
	return qs422016
//line goqtpl/marshal.qtpl:116
}

//line goqtpl/marshal.qtpl:118
func StreamMakeMarshal(qw422016 *qt422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/marshal.qtpl:118
	qw422016.N().S(`
func (v *`)
//line goqtpl/marshal.qtpl:119
	qw422016.E().S(name)
//line goqtpl/marshal.qtpl:119
	qw422016.N().S(`) WriteTo(w io.Writer) {
    `)
//line goqtpl/marshal.qtpl:120
	if sumSize(fields) > 0 {
//line goqtpl/marshal.qtpl:120
		qw422016.N().S(`
    var staticBuffer [`)
//line goqtpl/marshal.qtpl:121
		qw422016.N().D(sumSize(fields))
//line goqtpl/marshal.qtpl:121
		qw422016.N().S(`]byte
    `)
//line goqtpl/marshal.qtpl:122
	}
//line goqtpl/marshal.qtpl:122
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:123
	for i, field := range fields {
//line goqtpl/marshal.qtpl:123
		qw422016.N().S(`

    `)
//line goqtpl/marshal.qtpl:125
		if field.IsFixed {
//line goqtpl/marshal.qtpl:125
			qw422016.N().S(`
    // `)
//line goqtpl/marshal.qtpl:126
			qw422016.N().D(i)
//line goqtpl/marshal.qtpl:126
			qw422016.N().S(` : `)
//line goqtpl/marshal.qtpl:126
			qw422016.E().S(name)
//line goqtpl/marshal.qtpl:126
			qw422016.E().S(field.Path)
//line goqtpl/marshal.qtpl:126
			qw422016.N().S(`
    // Type : `)
//line goqtpl/marshal.qtpl:127
			qw422016.E().S(field.RawType)
//line goqtpl/marshal.qtpl:127
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:128
			if field.IsFixed {
//line goqtpl/marshal.qtpl:128
				qw422016.N().S(`// Size : `)
//line goqtpl/marshal.qtpl:128
				qw422016.N().D(field.Size)
//line goqtpl/marshal.qtpl:128
			} else {
//line goqtpl/marshal.qtpl:128
				qw422016.N().S(`// Size : Variable`)
//line goqtpl/marshal.qtpl:128
			}
//line goqtpl/marshal.qtpl:128
			qw422016.N().S(`
    // Offset : `)
//line goqtpl/marshal.qtpl:129
			qw422016.N().D(field.Offset)
//line goqtpl/marshal.qtpl:129
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:130
			StreamSerializeStatic(qw422016, "v"+field.Path, field.RawType, field.Offset, field.Size)
//line goqtpl/marshal.qtpl:130
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:131
		}
//line goqtpl/marshal.qtpl:131
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:132
	}
//line goqtpl/marshal.qtpl:132
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:133
	if sumSize(fields) > 0 {
//line goqtpl/marshal.qtpl:133
		qw422016.N().S(`
    w.Write(staticBuffer[:])
    `)
//line goqtpl/marshal.qtpl:135
	}
//line goqtpl/marshal.qtpl:135
	qw422016.N().S(`

    `)
//line goqtpl/marshal.qtpl:137
	for i, field := range fields {
//line goqtpl/marshal.qtpl:137
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:138
		if !field.IsFixed {
//line goqtpl/marshal.qtpl:138
			qw422016.N().S(`
    // `)
//line goqtpl/marshal.qtpl:139
			qw422016.N().D(i)
//line goqtpl/marshal.qtpl:139
			qw422016.N().S(` : `)
//line goqtpl/marshal.qtpl:139
			qw422016.E().S(name)
//line goqtpl/marshal.qtpl:139
			qw422016.E().S(field.Path)
//line goqtpl/marshal.qtpl:139
			qw422016.N().S(`
    // Type : `)
//line goqtpl/marshal.qtpl:140
			qw422016.E().S(field.RawType)
//line goqtpl/marshal.qtpl:140
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:141
			if field.IsFixed {
//line goqtpl/marshal.qtpl:141
				qw422016.N().S(`// Size : `)
//line goqtpl/marshal.qtpl:141
				qw422016.N().D(field.Size)
//line goqtpl/marshal.qtpl:141
			} else {
//line goqtpl/marshal.qtpl:141
				qw422016.N().S(`// Size : Variable`)
//line goqtpl/marshal.qtpl:141
			}
//line goqtpl/marshal.qtpl:141
			qw422016.N().S(`
    // Offset : `)
//line goqtpl/marshal.qtpl:142
			qw422016.N().D(field.Offset)
//line goqtpl/marshal.qtpl:142
			qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:143
			if isBytes(field.RawType) {
//line goqtpl/marshal.qtpl:143
				qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:144
				StreamBytes(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/marshal.qtpl:144
				qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:145
			}
//line goqtpl/marshal.qtpl:145
			qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:146
		}
//line goqtpl/marshal.qtpl:146
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:147
	}
//line goqtpl/marshal.qtpl:147
	qw422016.N().S(`
}
`)
//line goqtpl/marshal.qtpl:149
}

//line goqtpl/marshal.qtpl:149
func WriteMakeMarshal(qq422016 qtio422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/marshal.qtpl:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:149
	StreamMakeMarshal(qw422016, name, fields)
//line goqtpl/marshal.qtpl:149
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:149
}

//line goqtpl/marshal.qtpl:149
func MakeMarshal(name string, fields []*parser.GenerateField) string {
//line goqtpl/marshal.qtpl:149
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:149
	WriteMakeMarshal(qb422016, name, fields)
//line goqtpl/marshal.qtpl:149
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:149
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:149
	return qs422016
//line goqtpl/marshal.qtpl:149
}

//line goqtpl/marshal.qtpl:151
func StreamIntegerStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:151
	qw422016.N().S(`
    // Size : `)
//line goqtpl/marshal.qtpl:152
	qw422016.N().D(Size)
//line goqtpl/marshal.qtpl:152
	qw422016.N().S(`, Offset : `)
//line goqtpl/marshal.qtpl:152
	qw422016.N().D(Offset)
//line goqtpl/marshal.qtpl:152
	qw422016.N().S(`, VarName : `)
//line goqtpl/marshal.qtpl:152
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:152
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:153
	for i := 0; i < Size; i++ {
//line goqtpl/marshal.qtpl:153
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:154
		if isUint(RawType) {
//line goqtpl/marshal.qtpl:154
			qw422016.N().S(`
        staticBuffer[`)
//line goqtpl/marshal.qtpl:155
			qw422016.N().D(Offset + Size - 1 - i)
//line goqtpl/marshal.qtpl:155
			qw422016.N().S(`] = byte(`)
//line goqtpl/marshal.qtpl:155
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:155
			if i != 0 {
//line goqtpl/marshal.qtpl:155
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:155
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:155
			}
//line goqtpl/marshal.qtpl:155
			qw422016.N().S(`)
        `)
//line goqtpl/marshal.qtpl:156
		} else {
//line goqtpl/marshal.qtpl:156
			qw422016.N().S(`
        staticBuffer[`)
//line goqtpl/marshal.qtpl:157
			qw422016.N().D(Offset + Size - 1 - i)
//line goqtpl/marshal.qtpl:157
			qw422016.N().S(`] = byte(`)
//line goqtpl/marshal.qtpl:157
			qw422016.N().S(getUintType(RawType))
//line goqtpl/marshal.qtpl:157
			qw422016.N().S(`(`)
//line goqtpl/marshal.qtpl:157
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:157
			if i != 0 {
//line goqtpl/marshal.qtpl:157
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:157
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:157
			}
//line goqtpl/marshal.qtpl:157
			qw422016.N().S(`))
        `)
//line goqtpl/marshal.qtpl:158
		}
//line goqtpl/marshal.qtpl:158
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:159
	}
//line goqtpl/marshal.qtpl:159
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:160
}

//line goqtpl/marshal.qtpl:160
func WriteIntegerStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:160
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:160
	StreamIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:160
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:160
}

//line goqtpl/marshal.qtpl:160
func IntegerStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/marshal.qtpl:160
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:160
	WriteIntegerStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:160
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:160
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:160
	return qs422016
//line goqtpl/marshal.qtpl:160
}

//line goqtpl/marshal.qtpl:162
func StreamFloatStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:162
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:163
	if Size == 4 {
//line goqtpl/marshal.qtpl:163
		qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:164
		qw422016.N().DUL(xh(VarName + RawType))
//line goqtpl/marshal.qtpl:164
		qw422016.N().S(` := math.Float32bits(`)
//line goqtpl/marshal.qtpl:164
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:164
		qw422016.N().S(`)
		`)
//line goqtpl/marshal.qtpl:165
		StreamIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType)), "u32", Offset, Size)
//line goqtpl/marshal.qtpl:165
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:166
	} else if Size == 8 {
//line goqtpl/marshal.qtpl:166
		qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:167
		qw422016.N().DUL(xh(VarName + RawType))
//line goqtpl/marshal.qtpl:167
		qw422016.N().S(` := math.Float64bits(`)
//line goqtpl/marshal.qtpl:167
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:167
		qw422016.N().S(`)
		`)
//line goqtpl/marshal.qtpl:168
		StreamIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType)), "u64", Offset, Size)
//line goqtpl/marshal.qtpl:168
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:169
	}
//line goqtpl/marshal.qtpl:169
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:170
}

//line goqtpl/marshal.qtpl:170
func WriteFloatStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:170
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:170
	StreamFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:170
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:170
}

//line goqtpl/marshal.qtpl:170
func FloatStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/marshal.qtpl:170
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:170
	WriteFloatStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:170
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:170
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:170
	return qs422016
//line goqtpl/marshal.qtpl:170
}

//line goqtpl/marshal.qtpl:173
func StreamInteger(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/marshal.qtpl:173
	qw422016.N().S(`
    // Size : `)
//line goqtpl/marshal.qtpl:174
	qw422016.N().D(Size)
//line goqtpl/marshal.qtpl:174
	qw422016.N().S(`, VarName : `)
//line goqtpl/marshal.qtpl:174
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:174
	qw422016.N().S(`
    w.Write([]byte{
	`)
//line goqtpl/marshal.qtpl:176
	for i := Size - 1; i >= 0; i-- {
//line goqtpl/marshal.qtpl:176
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:177
		if isUint(RawType) {
//line goqtpl/marshal.qtpl:177
			qw422016.N().S(`
        byte(`)
//line goqtpl/marshal.qtpl:178
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:178
			if i != 0 {
//line goqtpl/marshal.qtpl:178
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:178
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:178
			}
//line goqtpl/marshal.qtpl:178
			qw422016.N().S(`),
        `)
//line goqtpl/marshal.qtpl:179
		} else {
//line goqtpl/marshal.qtpl:179
			qw422016.N().S(`
        byte(`)
//line goqtpl/marshal.qtpl:180
			qw422016.N().S(getUintType(RawType))
//line goqtpl/marshal.qtpl:180
			qw422016.N().S(`(`)
//line goqtpl/marshal.qtpl:180
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:180
			if i != 0 {
//line goqtpl/marshal.qtpl:180
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:180
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:180
			}
//line goqtpl/marshal.qtpl:180
			qw422016.N().S(`)),
        `)
//line goqtpl/marshal.qtpl:181
		}
//line goqtpl/marshal.qtpl:181
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:182
	}
//line goqtpl/marshal.qtpl:182
	qw422016.N().S(`
	})
`)
//line goqtpl/marshal.qtpl:184
}

//line goqtpl/marshal.qtpl:184
func WriteInteger(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/marshal.qtpl:184
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:184
	StreamInteger(qw422016, VarName, RawType, Size)
//line goqtpl/marshal.qtpl:184
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:184
}

//line goqtpl/marshal.qtpl:184
func Integer(VarName string, RawType string, Size int) string {
//line goqtpl/marshal.qtpl:184
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:184
	WriteInteger(qb422016, VarName, RawType, Size)
//line goqtpl/marshal.qtpl:184
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:184
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:184
	return qs422016
//line goqtpl/marshal.qtpl:184
}

//line goqtpl/marshal.qtpl:186
func StreamBytes(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/marshal.qtpl:186
	qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:187
	qw422016.N().DUL(xh(VarName + RawType))
//line goqtpl/marshal.qtpl:187
	qw422016.N().S(` := uint32(len(`)
//line goqtpl/marshal.qtpl:187
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:187
	qw422016.N().S(`))
		`)
//line goqtpl/marshal.qtpl:188
	StreamInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType)), "u32", 4)
//line goqtpl/marshal.qtpl:188
	qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:189
	if RawType == "bytes" {
//line goqtpl/marshal.qtpl:189
		qw422016.N().S(`
		w.Write(VarName)
	`)
//line goqtpl/marshal.qtpl:191
	} else if RawType == "str" {
//line goqtpl/marshal.qtpl:191
		qw422016.N().S(`
		w.Write(`)
//line goqtpl/marshal.qtpl:192
		qw422016.N().S(fmt.Sprintf("[]byte(%v)", VarName))
//line goqtpl/marshal.qtpl:192
		qw422016.N().S(`)
	`)
//line goqtpl/marshal.qtpl:193
	}
//line goqtpl/marshal.qtpl:193
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:194
}

//line goqtpl/marshal.qtpl:194
func WriteBytes(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/marshal.qtpl:194
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:194
	StreamBytes(qw422016, VarName, RawType)
//line goqtpl/marshal.qtpl:194
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:194
}

//line goqtpl/marshal.qtpl:194
func Bytes(VarName string, RawType string) string {
//line goqtpl/marshal.qtpl:194
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:194
	WriteBytes(qb422016, VarName, RawType)
//line goqtpl/marshal.qtpl:194
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:194
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:194
	return qs422016
//line goqtpl/marshal.qtpl:194
}
