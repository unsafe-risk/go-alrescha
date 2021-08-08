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

func varSize(RawType string) int {
	switch RawType {
	case "i8":
		return 1
	case "i16":
		return 2
	case "i32":
		return 4
	case "i64":
		return 8
	case "u8":
		return 1
	case "u16":
		return 2
	case "u32":
		return 4
	case "u64":
		return 8
	case "f32":
		return 4
	case "f64":
		return 8
	}
	return 0
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
		if fields[i].IsFixed && !fields[i].IsList {
			size += fields[i].Size
		}
	}
	return size
}

//line goqtpl/marshal.qtpl:134
func StreamSerializeStatic(qw422016 *qt422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:134
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:135
	if isInteger(RawType) {
//line goqtpl/marshal.qtpl:135
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:136
		StreamSerializeIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:136
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:137
	} else if isFloat(RawType) {
//line goqtpl/marshal.qtpl:137
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:138
		StreamSerializeFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:138
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:139
	} else {
//line goqtpl/marshal.qtpl:139
		qw422016.N().S(`
	// Failed to serialize `)
//line goqtpl/marshal.qtpl:140
		qw422016.E().S(RawType)
//line goqtpl/marshal.qtpl:140
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:141
	}
//line goqtpl/marshal.qtpl:141
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:142
}

//line goqtpl/marshal.qtpl:142
func WriteSerializeStatic(qq422016 qtio422016.Writer, VarName, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:142
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:142
	StreamSerializeStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:142
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:142
}

//line goqtpl/marshal.qtpl:142
func SerializeStatic(VarName, RawType string, Offset, Size int) string {
//line goqtpl/marshal.qtpl:142
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:142
	WriteSerializeStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:142
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:142
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:142
	return qs422016
//line goqtpl/marshal.qtpl:142
}

//line goqtpl/marshal.qtpl:144
func StreamMakeMarshal(qw422016 *qt422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/marshal.qtpl:144
	qw422016.N().S(`
func (v *`)
//line goqtpl/marshal.qtpl:145
	qw422016.E().S(name)
//line goqtpl/marshal.qtpl:145
	qw422016.N().S(`) wt(w io.Writer) {
    `)
//line goqtpl/marshal.qtpl:146
	if sumSize(fields) > 0 {
//line goqtpl/marshal.qtpl:146
		qw422016.N().S(`
    var staticBuffer [`)
//line goqtpl/marshal.qtpl:147
		qw422016.N().D(sumSize(fields))
//line goqtpl/marshal.qtpl:147
		qw422016.N().S(`]byte
    `)
//line goqtpl/marshal.qtpl:148
	}
//line goqtpl/marshal.qtpl:148
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:149
	for i, field := range fields {
//line goqtpl/marshal.qtpl:149
		qw422016.N().S(`

    `)
//line goqtpl/marshal.qtpl:151
		if field.IsFixed && !field.IsList {
//line goqtpl/marshal.qtpl:151
			qw422016.N().S(`
    // `)
//line goqtpl/marshal.qtpl:152
			qw422016.N().D(i)
//line goqtpl/marshal.qtpl:152
			qw422016.N().S(` : `)
//line goqtpl/marshal.qtpl:152
			qw422016.E().S(name)
//line goqtpl/marshal.qtpl:152
			qw422016.E().S(field.Path)
//line goqtpl/marshal.qtpl:152
			qw422016.N().S(`
    // Type : `)
//line goqtpl/marshal.qtpl:153
			qw422016.E().S(field.RawType)
//line goqtpl/marshal.qtpl:153
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:154
			if field.IsFixed {
//line goqtpl/marshal.qtpl:154
				qw422016.N().S(`// Size : `)
//line goqtpl/marshal.qtpl:154
				qw422016.N().D(field.Size)
//line goqtpl/marshal.qtpl:154
			} else {
//line goqtpl/marshal.qtpl:154
				qw422016.N().S(`// Size : Variable`)
//line goqtpl/marshal.qtpl:154
			}
//line goqtpl/marshal.qtpl:154
			qw422016.N().S(`
    // Offset : `)
//line goqtpl/marshal.qtpl:155
			qw422016.N().D(field.Offset)
//line goqtpl/marshal.qtpl:155
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:156
			StreamSerializeStatic(qw422016, "v"+field.Path, field.RawType, field.Offset, field.Size)
//line goqtpl/marshal.qtpl:156
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:157
		}
//line goqtpl/marshal.qtpl:157
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:158
	}
//line goqtpl/marshal.qtpl:158
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:159
	if sumSize(fields) > 0 {
//line goqtpl/marshal.qtpl:159
		qw422016.N().S(`
    w.Write(staticBuffer[:])
    `)
//line goqtpl/marshal.qtpl:161
	}
//line goqtpl/marshal.qtpl:161
	qw422016.N().S(`

    `)
//line goqtpl/marshal.qtpl:163
	for i, field := range fields {
//line goqtpl/marshal.qtpl:163
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:164
		if field.IsList {
//line goqtpl/marshal.qtpl:164
			qw422016.N().S(`
	// `)
//line goqtpl/marshal.qtpl:165
			qw422016.N().D(i)
//line goqtpl/marshal.qtpl:165
			qw422016.N().S(` : `)
//line goqtpl/marshal.qtpl:165
			qw422016.E().S(name)
//line goqtpl/marshal.qtpl:165
			qw422016.E().S(field.Path)
//line goqtpl/marshal.qtpl:165
			qw422016.N().S(`
    // Type : `)
//line goqtpl/marshal.qtpl:166
			qw422016.E().S(field.RawType)
//line goqtpl/marshal.qtpl:166
			qw422016.N().S(` (List)
	// Size : Variable
	`)
//line goqtpl/marshal.qtpl:168
			StreamList(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/marshal.qtpl:168
			qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:169
		}
//line goqtpl/marshal.qtpl:169
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:170
		if !field.IsFixed {
//line goqtpl/marshal.qtpl:170
			qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:171
			if isBytes(field.RawType) && !field.IsList {
//line goqtpl/marshal.qtpl:171
				qw422016.N().S(`
	// `)
//line goqtpl/marshal.qtpl:172
				qw422016.N().D(i)
//line goqtpl/marshal.qtpl:172
				qw422016.N().S(` : `)
//line goqtpl/marshal.qtpl:172
				qw422016.E().S(name)
//line goqtpl/marshal.qtpl:172
				qw422016.E().S(field.Path)
//line goqtpl/marshal.qtpl:172
				qw422016.N().S(`
    // Type : `)
//line goqtpl/marshal.qtpl:173
				qw422016.E().S(field.RawType)
//line goqtpl/marshal.qtpl:173
				qw422016.N().S(`
    // Size : Variable
	`)
//line goqtpl/marshal.qtpl:175
				StreamBytes(qw422016, "v"+field.Path, field.RawType)
//line goqtpl/marshal.qtpl:175
				qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:176
			}
//line goqtpl/marshal.qtpl:176
			qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:177
		}
//line goqtpl/marshal.qtpl:177
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:178
	}
//line goqtpl/marshal.qtpl:178
	qw422016.N().S(`
}
`)
//line goqtpl/marshal.qtpl:180
}

//line goqtpl/marshal.qtpl:180
func WriteMakeMarshal(qq422016 qtio422016.Writer, name string, fields []*parser.GenerateField) {
//line goqtpl/marshal.qtpl:180
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:180
	StreamMakeMarshal(qw422016, name, fields)
//line goqtpl/marshal.qtpl:180
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:180
}

//line goqtpl/marshal.qtpl:180
func MakeMarshal(name string, fields []*parser.GenerateField) string {
//line goqtpl/marshal.qtpl:180
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:180
	WriteMakeMarshal(qb422016, name, fields)
//line goqtpl/marshal.qtpl:180
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:180
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:180
	return qs422016
//line goqtpl/marshal.qtpl:180
}

//line goqtpl/marshal.qtpl:182
func StreamSerializeIntegerStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:182
	qw422016.N().S(`
    // Size : `)
//line goqtpl/marshal.qtpl:183
	qw422016.N().D(Size)
//line goqtpl/marshal.qtpl:183
	qw422016.N().S(`, Offset : `)
//line goqtpl/marshal.qtpl:183
	qw422016.N().D(Offset)
//line goqtpl/marshal.qtpl:183
	qw422016.N().S(`, VarName : `)
//line goqtpl/marshal.qtpl:183
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:183
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:184
	for i := 0; i < Size; i++ {
//line goqtpl/marshal.qtpl:184
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:185
		if isUint(RawType) {
//line goqtpl/marshal.qtpl:185
			qw422016.N().S(`
        staticBuffer[`)
//line goqtpl/marshal.qtpl:186
			qw422016.N().D(Offset + Size - 1 - i)
//line goqtpl/marshal.qtpl:186
			qw422016.N().S(`] = byte(`)
//line goqtpl/marshal.qtpl:186
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:186
			if i != 0 {
//line goqtpl/marshal.qtpl:186
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:186
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:186
			}
//line goqtpl/marshal.qtpl:186
			qw422016.N().S(`)
        `)
//line goqtpl/marshal.qtpl:187
		} else {
//line goqtpl/marshal.qtpl:187
			qw422016.N().S(`
        staticBuffer[`)
//line goqtpl/marshal.qtpl:188
			qw422016.N().D(Offset + Size - 1 - i)
//line goqtpl/marshal.qtpl:188
			qw422016.N().S(`] = byte(`)
//line goqtpl/marshal.qtpl:188
			qw422016.N().S(getUintType(RawType))
//line goqtpl/marshal.qtpl:188
			qw422016.N().S(`(`)
//line goqtpl/marshal.qtpl:188
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:188
			if i != 0 {
//line goqtpl/marshal.qtpl:188
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:188
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:188
			}
//line goqtpl/marshal.qtpl:188
			qw422016.N().S(`))
        `)
//line goqtpl/marshal.qtpl:189
		}
//line goqtpl/marshal.qtpl:189
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:190
	}
//line goqtpl/marshal.qtpl:190
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:191
}

//line goqtpl/marshal.qtpl:191
func WriteSerializeIntegerStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:191
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:191
	StreamSerializeIntegerStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:191
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:191
}

//line goqtpl/marshal.qtpl:191
func SerializeIntegerStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/marshal.qtpl:191
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:191
	WriteSerializeIntegerStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:191
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:191
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:191
	return qs422016
//line goqtpl/marshal.qtpl:191
}

//line goqtpl/marshal.qtpl:193
func StreamSerializeFloatStatic(qw422016 *qt422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:193
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:194
	if Size == 4 {
//line goqtpl/marshal.qtpl:194
		qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:195
		qw422016.N().DUL(xh(VarName + RawType + "SerializeFloatStatic"))
//line goqtpl/marshal.qtpl:195
		qw422016.N().S(` := math.Float32bits(`)
//line goqtpl/marshal.qtpl:195
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:195
		qw422016.N().S(`)
		`)
//line goqtpl/marshal.qtpl:196
		StreamSerializeIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"SerializeFloatStatic")), "u32", Offset, Size)
//line goqtpl/marshal.qtpl:196
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:197
	} else if Size == 8 {
//line goqtpl/marshal.qtpl:197
		qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:198
		qw422016.N().DUL(xh(VarName + RawType + "SerializeFloatStatic"))
//line goqtpl/marshal.qtpl:198
		qw422016.N().S(` := math.Float64bits(`)
//line goqtpl/marshal.qtpl:198
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:198
		qw422016.N().S(`)
		`)
//line goqtpl/marshal.qtpl:199
		StreamSerializeIntegerStatic(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"SerializeFloatStatic")), "u64", Offset, Size)
//line goqtpl/marshal.qtpl:199
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:200
	}
//line goqtpl/marshal.qtpl:200
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:201
}

//line goqtpl/marshal.qtpl:201
func WriteSerializeFloatStatic(qq422016 qtio422016.Writer, VarName string, RawType string, Offset, Size int) {
//line goqtpl/marshal.qtpl:201
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:201
	StreamSerializeFloatStatic(qw422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:201
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:201
}

//line goqtpl/marshal.qtpl:201
func SerializeFloatStatic(VarName string, RawType string, Offset, Size int) string {
//line goqtpl/marshal.qtpl:201
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:201
	WriteSerializeFloatStatic(qb422016, VarName, RawType, Offset, Size)
//line goqtpl/marshal.qtpl:201
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:201
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:201
	return qs422016
//line goqtpl/marshal.qtpl:201
}

//line goqtpl/marshal.qtpl:204
func StreamInteger(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/marshal.qtpl:204
	qw422016.N().S(`
    // Size : `)
//line goqtpl/marshal.qtpl:205
	qw422016.N().D(Size)
//line goqtpl/marshal.qtpl:205
	qw422016.N().S(`, VarName : `)
//line goqtpl/marshal.qtpl:205
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:205
	qw422016.N().S(`
    w.Write([]byte{
	`)
//line goqtpl/marshal.qtpl:207
	for i := Size - 1; i >= 0; i-- {
//line goqtpl/marshal.qtpl:207
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:208
		if isUint(RawType) {
//line goqtpl/marshal.qtpl:208
			qw422016.N().S(`
        byte(`)
//line goqtpl/marshal.qtpl:209
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:209
			if i != 0 {
//line goqtpl/marshal.qtpl:209
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:209
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:209
			}
//line goqtpl/marshal.qtpl:209
			qw422016.N().S(`),
        `)
//line goqtpl/marshal.qtpl:210
		} else {
//line goqtpl/marshal.qtpl:210
			qw422016.N().S(`
        byte(`)
//line goqtpl/marshal.qtpl:211
			qw422016.N().S(getUintType(RawType))
//line goqtpl/marshal.qtpl:211
			qw422016.N().S(`(`)
//line goqtpl/marshal.qtpl:211
			qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:211
			if i != 0 {
//line goqtpl/marshal.qtpl:211
				qw422016.N().S(` >> `)
//line goqtpl/marshal.qtpl:211
				qw422016.N().D(i * 8)
//line goqtpl/marshal.qtpl:211
			}
//line goqtpl/marshal.qtpl:211
			qw422016.N().S(`)),
        `)
//line goqtpl/marshal.qtpl:212
		}
//line goqtpl/marshal.qtpl:212
		qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:213
	}
//line goqtpl/marshal.qtpl:213
	qw422016.N().S(`
	})
`)
//line goqtpl/marshal.qtpl:215
}

//line goqtpl/marshal.qtpl:215
func WriteInteger(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/marshal.qtpl:215
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:215
	StreamInteger(qw422016, VarName, RawType, Size)
//line goqtpl/marshal.qtpl:215
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:215
}

//line goqtpl/marshal.qtpl:215
func Integer(VarName string, RawType string, Size int) string {
//line goqtpl/marshal.qtpl:215
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:215
	WriteInteger(qb422016, VarName, RawType, Size)
//line goqtpl/marshal.qtpl:215
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:215
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:215
	return qs422016
//line goqtpl/marshal.qtpl:215
}

//line goqtpl/marshal.qtpl:217
func StreamFloat(qw422016 *qt422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/marshal.qtpl:217
	qw422016.N().S(`
    `)
//line goqtpl/marshal.qtpl:218
	if Size == 4 {
//line goqtpl/marshal.qtpl:218
		qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:219
		qw422016.N().DUL(xh(VarName + RawType + "Float"))
//line goqtpl/marshal.qtpl:219
		qw422016.N().S(` := math.Float32bits(`)
//line goqtpl/marshal.qtpl:219
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:219
		qw422016.N().S(`)
		`)
//line goqtpl/marshal.qtpl:220
		StreamInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"Float")), "u32", Size)
//line goqtpl/marshal.qtpl:220
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:221
	} else if Size == 8 {
//line goqtpl/marshal.qtpl:221
		qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:222
		qw422016.N().DUL(xh(VarName + RawType + "Float"))
//line goqtpl/marshal.qtpl:222
		qw422016.N().S(` := math.Float64bits(`)
//line goqtpl/marshal.qtpl:222
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:222
		qw422016.N().S(`)
		`)
//line goqtpl/marshal.qtpl:223
		StreamInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"Float")), "u64", Size)
//line goqtpl/marshal.qtpl:223
		qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:224
	}
//line goqtpl/marshal.qtpl:224
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:225
}

//line goqtpl/marshal.qtpl:225
func WriteFloat(qq422016 qtio422016.Writer, VarName string, RawType string, Size int) {
//line goqtpl/marshal.qtpl:225
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:225
	StreamFloat(qw422016, VarName, RawType, Size)
//line goqtpl/marshal.qtpl:225
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:225
}

//line goqtpl/marshal.qtpl:225
func Float(VarName string, RawType string, Size int) string {
//line goqtpl/marshal.qtpl:225
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:225
	WriteFloat(qb422016, VarName, RawType, Size)
//line goqtpl/marshal.qtpl:225
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:225
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:225
	return qs422016
//line goqtpl/marshal.qtpl:225
}

//line goqtpl/marshal.qtpl:227
func StreamBytes(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/marshal.qtpl:227
	qw422016.N().S(`
		v`)
//line goqtpl/marshal.qtpl:228
	qw422016.N().DUL(xh(VarName + RawType + "Bytes"))
//line goqtpl/marshal.qtpl:228
	qw422016.N().S(` := uint32(len(`)
//line goqtpl/marshal.qtpl:228
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:228
	qw422016.N().S(`))
		`)
//line goqtpl/marshal.qtpl:229
	StreamInteger(qw422016, "v"+fmt.Sprint(xh(VarName+RawType+"Bytes")), "u32", 4)
//line goqtpl/marshal.qtpl:229
	qw422016.N().S(`
	`)
//line goqtpl/marshal.qtpl:230
	if RawType == "bytes" {
//line goqtpl/marshal.qtpl:230
		qw422016.N().S(`
		w.Write(VarName)
	`)
//line goqtpl/marshal.qtpl:232
	} else if RawType == "str" {
//line goqtpl/marshal.qtpl:232
		qw422016.N().S(`
		w.Write(`)
//line goqtpl/marshal.qtpl:233
		qw422016.N().S(fmt.Sprintf("[]byte(%v)", VarName))
//line goqtpl/marshal.qtpl:233
		qw422016.N().S(`)
	`)
//line goqtpl/marshal.qtpl:234
	}
//line goqtpl/marshal.qtpl:234
	qw422016.N().S(`
`)
//line goqtpl/marshal.qtpl:235
}

//line goqtpl/marshal.qtpl:235
func WriteBytes(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/marshal.qtpl:235
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:235
	StreamBytes(qw422016, VarName, RawType)
//line goqtpl/marshal.qtpl:235
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:235
}

//line goqtpl/marshal.qtpl:235
func Bytes(VarName string, RawType string) string {
//line goqtpl/marshal.qtpl:235
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:235
	WriteBytes(qb422016, VarName, RawType)
//line goqtpl/marshal.qtpl:235
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:235
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:235
	return qs422016
//line goqtpl/marshal.qtpl:235
}

//line goqtpl/marshal.qtpl:237
func StreamList(qw422016 *qt422016.Writer, VarName string, RawType string) {
//line goqtpl/marshal.qtpl:237
	qw422016.N().S(`
	v`)
//line goqtpl/marshal.qtpl:238
	qw422016.N().DUL(xh(VarName + RawType + "List"))
//line goqtpl/marshal.qtpl:238
	qw422016.N().S(` := uint32(len(`)
//line goqtpl/marshal.qtpl:238
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:238
	qw422016.N().S(`))
	`)
//line goqtpl/marshal.qtpl:239
	StreamInteger(qw422016, fmt.Sprintf("v%d", xh(VarName+RawType+"List")), "u32", 4)
//line goqtpl/marshal.qtpl:239
	qw422016.N().S(`
	// == List ==
	// List VarName : `)
//line goqtpl/marshal.qtpl:241
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:241
	qw422016.N().S(`
	// List VarType : `)
//line goqtpl/marshal.qtpl:242
	qw422016.E().S(RawType)
//line goqtpl/marshal.qtpl:242
	qw422016.N().S(`
	for i := range `)
//line goqtpl/marshal.qtpl:243
	qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:243
	qw422016.N().S(` {
		`)
//line goqtpl/marshal.qtpl:244
	if isBytes(RawType) {
//line goqtpl/marshal.qtpl:244
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:245
		StreamBytes(qw422016, VarName+"[i]", RawType)
//line goqtpl/marshal.qtpl:245
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:246
	} else if isInteger(RawType) {
//line goqtpl/marshal.qtpl:246
		qw422016.N().S(`
        `)
//line goqtpl/marshal.qtpl:247
		StreamInteger(qw422016, VarName+"[i]", RawType, varSize(RawType))
//line goqtpl/marshal.qtpl:247
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:248
	} else if isFloat(RawType) {
//line goqtpl/marshal.qtpl:248
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:249
		StreamFloat(qw422016, VarName+"[i]", RawType, varSize(RawType))
//line goqtpl/marshal.qtpl:249
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:250
	} else {
//line goqtpl/marshal.qtpl:250
		qw422016.N().S(`
		`)
//line goqtpl/marshal.qtpl:251
		qw422016.E().S(VarName)
//line goqtpl/marshal.qtpl:251
		qw422016.N().S(`[i].wt(w)
		`)
//line goqtpl/marshal.qtpl:252
	}
//line goqtpl/marshal.qtpl:252
	qw422016.N().S(`
	}
`)
//line goqtpl/marshal.qtpl:254
}

//line goqtpl/marshal.qtpl:254
func WriteList(qq422016 qtio422016.Writer, VarName string, RawType string) {
//line goqtpl/marshal.qtpl:254
	qw422016 := qt422016.AcquireWriter(qq422016)
//line goqtpl/marshal.qtpl:254
	StreamList(qw422016, VarName, RawType)
//line goqtpl/marshal.qtpl:254
	qt422016.ReleaseWriter(qw422016)
//line goqtpl/marshal.qtpl:254
}

//line goqtpl/marshal.qtpl:254
func List(VarName string, RawType string) string {
//line goqtpl/marshal.qtpl:254
	qb422016 := qt422016.AcquireByteBuffer()
//line goqtpl/marshal.qtpl:254
	WriteList(qb422016, VarName, RawType)
//line goqtpl/marshal.qtpl:254
	qs422016 := string(qb422016.B)
//line goqtpl/marshal.qtpl:254
	qt422016.ReleaseByteBuffer(qb422016)
//line goqtpl/marshal.qtpl:254
	return qs422016
//line goqtpl/marshal.qtpl:254
}
