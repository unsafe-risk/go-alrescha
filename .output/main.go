package main

import (
	"io"
	"math"
)

var _ = math.Float64bits
var _ = io.EOF

type Person struct {
	Name string

	Age uint8

	Address Address

	Phone PhoneNumber

	Email string

	Acls [16]string

	LastConnection float64

	Float64s [16]float64

	LastUpdated float64
}

func (v *Person) GetName() string {
	return v.Name
}

func (v *Person) SetName(value string) {
	v.Name = value
}

func (v *Person) GetAge() uint8 {
	return v.Age
}

func (v *Person) SetAge(value uint8) {
	v.Age = value
}

func (v *Person) GetAddress() Address {
	return v.Address
}

func (v *Person) SetAddress(value Address) {
	v.Address = value
}

func (v *Person) GetPhone() PhoneNumber {
	return v.Phone
}

func (v *Person) SetPhone(value PhoneNumber) {
	v.Phone = value
}

func (v *Person) GetEmail() string {
	return v.Email
}

func (v *Person) SetEmail(value string) {
	v.Email = value
}

func (v *Person) GetAcls() [16]string {
	return v.Acls
}

func (v *Person) SetAcls(value [16]string) {
	v.Acls = value
}

func (v *Person) GetLastConnection() float64 {
	return v.LastConnection
}

func (v *Person) SetLastConnection(value float64) {
	v.LastConnection = value
}

func (v *Person) GetFloat64s() [16]float64 {
	return v.Float64s
}

func (v *Person) SetFloat64s(value [16]float64) {
	v.Float64s = value
}

func (v *Person) GetLastUpdated() float64 {
	return v.LastUpdated
}

func (v *Person) SetLastUpdated(value float64) {
	v.LastUpdated = value
}



type Address struct {
	Street string

	City string

	State string

	Zip uint16
}

func (v *Address) GetStreet() string {
	return v.Street
}

func (v *Address) SetStreet(value string) {
	v.Street = value
}

func (v *Address) GetCity() string {
	return v.City
}

func (v *Address) SetCity(value string) {
	v.City = value
}

func (v *Address) GetState() string {
	return v.State
}

func (v *Address) SetState(value string) {
	v.State = value
}

func (v *Address) GetZip() uint16 {
	return v.Zip
}

func (v *Address) SetZip(value uint16) {
	v.Zip = value
}



type PhoneNumber struct {
	Number string

	Type string
}

func (v *PhoneNumber) GetNumber() string {
	return v.Number
}

func (v *PhoneNumber) SetNumber(value string) {
	v.Number = value
}

func (v *PhoneNumber) GetType() string {
	return v.Type
}

func (v *PhoneNumber) SetType(value string) {
	v.Type = value
}



func (v *Person) wt(w io.Writer) {

	var staticBuffer [147]byte

	// 0 : Person.LastConnection
	// Type : f64
	// Size : 8
	// Offset : 0

	v6706486543488242042 := math.Float64bits(v.LastConnection)

	// Size : 8, Offset : 0, VarName : v6706486543488242042

	staticBuffer[7] = byte(v6706486543488242042)

	staticBuffer[6] = byte(v6706486543488242042 >> 8)

	staticBuffer[5] = byte(v6706486543488242042 >> 16)

	staticBuffer[4] = byte(v6706486543488242042 >> 24)

	staticBuffer[3] = byte(v6706486543488242042 >> 32)

	staticBuffer[2] = byte(v6706486543488242042 >> 40)

	staticBuffer[1] = byte(v6706486543488242042 >> 48)

	staticBuffer[0] = byte(v6706486543488242042 >> 56)

	// 1 : Person.LastUpdated
	// Type : f64
	// Size : 8
	// Offset : 8

	v4528304775663089888 := math.Float64bits(v.LastUpdated)

	// Size : 8, Offset : 8, VarName : v4528304775663089888

	staticBuffer[15] = byte(v4528304775663089888)

	staticBuffer[14] = byte(v4528304775663089888 >> 8)

	staticBuffer[13] = byte(v4528304775663089888 >> 16)

	staticBuffer[12] = byte(v4528304775663089888 >> 24)

	staticBuffer[11] = byte(v4528304775663089888 >> 32)

	staticBuffer[10] = byte(v4528304775663089888 >> 40)

	staticBuffer[9] = byte(v4528304775663089888 >> 48)

	staticBuffer[8] = byte(v4528304775663089888 >> 56)

	// 2 : Person.Float64s[0]
	// Type : f64
	// Size : 8
	// Offset : 16

	v12470977835207192315 := math.Float64bits(v.Float64s[0])

	// Size : 8, Offset : 16, VarName : v12470977835207192315

	staticBuffer[23] = byte(v12470977835207192315)

	staticBuffer[22] = byte(v12470977835207192315 >> 8)

	staticBuffer[21] = byte(v12470977835207192315 >> 16)

	staticBuffer[20] = byte(v12470977835207192315 >> 24)

	staticBuffer[19] = byte(v12470977835207192315 >> 32)

	staticBuffer[18] = byte(v12470977835207192315 >> 40)

	staticBuffer[17] = byte(v12470977835207192315 >> 48)

	staticBuffer[16] = byte(v12470977835207192315 >> 56)

	// 3 : Person.Float64s[1]
	// Type : f64
	// Size : 8
	// Offset : 24

	v5138655428976818856 := math.Float64bits(v.Float64s[1])

	// Size : 8, Offset : 24, VarName : v5138655428976818856

	staticBuffer[31] = byte(v5138655428976818856)

	staticBuffer[30] = byte(v5138655428976818856 >> 8)

	staticBuffer[29] = byte(v5138655428976818856 >> 16)

	staticBuffer[28] = byte(v5138655428976818856 >> 24)

	staticBuffer[27] = byte(v5138655428976818856 >> 32)

	staticBuffer[26] = byte(v5138655428976818856 >> 40)

	staticBuffer[25] = byte(v5138655428976818856 >> 48)

	staticBuffer[24] = byte(v5138655428976818856 >> 56)

	// 4 : Person.Float64s[2]
	// Type : f64
	// Size : 8
	// Offset : 32

	v1122683454161749385 := math.Float64bits(v.Float64s[2])

	// Size : 8, Offset : 32, VarName : v1122683454161749385

	staticBuffer[39] = byte(v1122683454161749385)

	staticBuffer[38] = byte(v1122683454161749385 >> 8)

	staticBuffer[37] = byte(v1122683454161749385 >> 16)

	staticBuffer[36] = byte(v1122683454161749385 >> 24)

	staticBuffer[35] = byte(v1122683454161749385 >> 32)

	staticBuffer[34] = byte(v1122683454161749385 >> 40)

	staticBuffer[33] = byte(v1122683454161749385 >> 48)

	staticBuffer[32] = byte(v1122683454161749385 >> 56)

	// 5 : Person.Float64s[3]
	// Type : f64
	// Size : 8
	// Offset : 40

	v17130242275333474189 := math.Float64bits(v.Float64s[3])

	// Size : 8, Offset : 40, VarName : v17130242275333474189

	staticBuffer[47] = byte(v17130242275333474189)

	staticBuffer[46] = byte(v17130242275333474189 >> 8)

	staticBuffer[45] = byte(v17130242275333474189 >> 16)

	staticBuffer[44] = byte(v17130242275333474189 >> 24)

	staticBuffer[43] = byte(v17130242275333474189 >> 32)

	staticBuffer[42] = byte(v17130242275333474189 >> 40)

	staticBuffer[41] = byte(v17130242275333474189 >> 48)

	staticBuffer[40] = byte(v17130242275333474189 >> 56)

	// 6 : Person.Float64s[4]
	// Type : f64
	// Size : 8
	// Offset : 48

	v13394140246851720790 := math.Float64bits(v.Float64s[4])

	// Size : 8, Offset : 48, VarName : v13394140246851720790

	staticBuffer[55] = byte(v13394140246851720790)

	staticBuffer[54] = byte(v13394140246851720790 >> 8)

	staticBuffer[53] = byte(v13394140246851720790 >> 16)

	staticBuffer[52] = byte(v13394140246851720790 >> 24)

	staticBuffer[51] = byte(v13394140246851720790 >> 32)

	staticBuffer[50] = byte(v13394140246851720790 >> 40)

	staticBuffer[49] = byte(v13394140246851720790 >> 48)

	staticBuffer[48] = byte(v13394140246851720790 >> 56)

	// 7 : Person.Float64s[5]
	// Type : f64
	// Size : 8
	// Offset : 56

	v3376172112261645491 := math.Float64bits(v.Float64s[5])

	// Size : 8, Offset : 56, VarName : v3376172112261645491

	staticBuffer[63] = byte(v3376172112261645491)

	staticBuffer[62] = byte(v3376172112261645491 >> 8)

	staticBuffer[61] = byte(v3376172112261645491 >> 16)

	staticBuffer[60] = byte(v3376172112261645491 >> 24)

	staticBuffer[59] = byte(v3376172112261645491 >> 32)

	staticBuffer[58] = byte(v3376172112261645491 >> 40)

	staticBuffer[57] = byte(v3376172112261645491 >> 48)

	staticBuffer[56] = byte(v3376172112261645491 >> 56)

	// 8 : Person.Float64s[6]
	// Type : f64
	// Size : 8
	// Offset : 64

	v9423497948159733789 := math.Float64bits(v.Float64s[6])

	// Size : 8, Offset : 64, VarName : v9423497948159733789

	staticBuffer[71] = byte(v9423497948159733789)

	staticBuffer[70] = byte(v9423497948159733789 >> 8)

	staticBuffer[69] = byte(v9423497948159733789 >> 16)

	staticBuffer[68] = byte(v9423497948159733789 >> 24)

	staticBuffer[67] = byte(v9423497948159733789 >> 32)

	staticBuffer[66] = byte(v9423497948159733789 >> 40)

	staticBuffer[65] = byte(v9423497948159733789 >> 48)

	staticBuffer[64] = byte(v9423497948159733789 >> 56)

	// 9 : Person.Float64s[7]
	// Type : f64
	// Size : 8
	// Offset : 72

	v10937475156283519705 := math.Float64bits(v.Float64s[7])

	// Size : 8, Offset : 72, VarName : v10937475156283519705

	staticBuffer[79] = byte(v10937475156283519705)

	staticBuffer[78] = byte(v10937475156283519705 >> 8)

	staticBuffer[77] = byte(v10937475156283519705 >> 16)

	staticBuffer[76] = byte(v10937475156283519705 >> 24)

	staticBuffer[75] = byte(v10937475156283519705 >> 32)

	staticBuffer[74] = byte(v10937475156283519705 >> 40)

	staticBuffer[73] = byte(v10937475156283519705 >> 48)

	staticBuffer[72] = byte(v10937475156283519705 >> 56)

	// 10 : Person.Float64s[8]
	// Type : f64
	// Size : 8
	// Offset : 80

	v7644229136777653739 := math.Float64bits(v.Float64s[8])

	// Size : 8, Offset : 80, VarName : v7644229136777653739

	staticBuffer[87] = byte(v7644229136777653739)

	staticBuffer[86] = byte(v7644229136777653739 >> 8)

	staticBuffer[85] = byte(v7644229136777653739 >> 16)

	staticBuffer[84] = byte(v7644229136777653739 >> 24)

	staticBuffer[83] = byte(v7644229136777653739 >> 32)

	staticBuffer[82] = byte(v7644229136777653739 >> 40)

	staticBuffer[81] = byte(v7644229136777653739 >> 48)

	staticBuffer[80] = byte(v7644229136777653739 >> 56)

	// 11 : Person.Float64s[9]
	// Type : f64
	// Size : 8
	// Offset : 88

	v4824756505142059494 := math.Float64bits(v.Float64s[9])

	// Size : 8, Offset : 88, VarName : v4824756505142059494

	staticBuffer[95] = byte(v4824756505142059494)

	staticBuffer[94] = byte(v4824756505142059494 >> 8)

	staticBuffer[93] = byte(v4824756505142059494 >> 16)

	staticBuffer[92] = byte(v4824756505142059494 >> 24)

	staticBuffer[91] = byte(v4824756505142059494 >> 32)

	staticBuffer[90] = byte(v4824756505142059494 >> 40)

	staticBuffer[89] = byte(v4824756505142059494 >> 48)

	staticBuffer[88] = byte(v4824756505142059494 >> 56)

	// 12 : Person.Float64s[10]
	// Type : f64
	// Size : 8
	// Offset : 96

	v2530353830496907173 := math.Float64bits(v.Float64s[10])

	// Size : 8, Offset : 96, VarName : v2530353830496907173

	staticBuffer[103] = byte(v2530353830496907173)

	staticBuffer[102] = byte(v2530353830496907173 >> 8)

	staticBuffer[101] = byte(v2530353830496907173 >> 16)

	staticBuffer[100] = byte(v2530353830496907173 >> 24)

	staticBuffer[99] = byte(v2530353830496907173 >> 32)

	staticBuffer[98] = byte(v2530353830496907173 >> 40)

	staticBuffer[97] = byte(v2530353830496907173 >> 48)

	staticBuffer[96] = byte(v2530353830496907173 >> 56)

	// 13 : Person.Float64s[11]
	// Type : f64
	// Size : 8
	// Offset : 104

	v6230011054625933431 := math.Float64bits(v.Float64s[11])

	// Size : 8, Offset : 104, VarName : v6230011054625933431

	staticBuffer[111] = byte(v6230011054625933431)

	staticBuffer[110] = byte(v6230011054625933431 >> 8)

	staticBuffer[109] = byte(v6230011054625933431 >> 16)

	staticBuffer[108] = byte(v6230011054625933431 >> 24)

	staticBuffer[107] = byte(v6230011054625933431 >> 32)

	staticBuffer[106] = byte(v6230011054625933431 >> 40)

	staticBuffer[105] = byte(v6230011054625933431 >> 48)

	staticBuffer[104] = byte(v6230011054625933431 >> 56)

	// 14 : Person.Float64s[12]
	// Type : f64
	// Size : 8
	// Offset : 112

	v12629137263417362073 := math.Float64bits(v.Float64s[12])

	// Size : 8, Offset : 112, VarName : v12629137263417362073

	staticBuffer[119] = byte(v12629137263417362073)

	staticBuffer[118] = byte(v12629137263417362073 >> 8)

	staticBuffer[117] = byte(v12629137263417362073 >> 16)

	staticBuffer[116] = byte(v12629137263417362073 >> 24)

	staticBuffer[115] = byte(v12629137263417362073 >> 32)

	staticBuffer[114] = byte(v12629137263417362073 >> 40)

	staticBuffer[113] = byte(v12629137263417362073 >> 48)

	staticBuffer[112] = byte(v12629137263417362073 >> 56)

	// 15 : Person.Float64s[13]
	// Type : f64
	// Size : 8
	// Offset : 120

	v10307908525389598458 := math.Float64bits(v.Float64s[13])

	// Size : 8, Offset : 120, VarName : v10307908525389598458

	staticBuffer[127] = byte(v10307908525389598458)

	staticBuffer[126] = byte(v10307908525389598458 >> 8)

	staticBuffer[125] = byte(v10307908525389598458 >> 16)

	staticBuffer[124] = byte(v10307908525389598458 >> 24)

	staticBuffer[123] = byte(v10307908525389598458 >> 32)

	staticBuffer[122] = byte(v10307908525389598458 >> 40)

	staticBuffer[121] = byte(v10307908525389598458 >> 48)

	staticBuffer[120] = byte(v10307908525389598458 >> 56)

	// 16 : Person.Float64s[14]
	// Type : f64
	// Size : 8
	// Offset : 128

	v1635293563065644244 := math.Float64bits(v.Float64s[14])

	// Size : 8, Offset : 128, VarName : v1635293563065644244

	staticBuffer[135] = byte(v1635293563065644244)

	staticBuffer[134] = byte(v1635293563065644244 >> 8)

	staticBuffer[133] = byte(v1635293563065644244 >> 16)

	staticBuffer[132] = byte(v1635293563065644244 >> 24)

	staticBuffer[131] = byte(v1635293563065644244 >> 32)

	staticBuffer[130] = byte(v1635293563065644244 >> 40)

	staticBuffer[129] = byte(v1635293563065644244 >> 48)

	staticBuffer[128] = byte(v1635293563065644244 >> 56)

	// 17 : Person.Float64s[15]
	// Type : f64
	// Size : 8
	// Offset : 136

	v8076638462475212619 := math.Float64bits(v.Float64s[15])

	// Size : 8, Offset : 136, VarName : v8076638462475212619

	staticBuffer[143] = byte(v8076638462475212619)

	staticBuffer[142] = byte(v8076638462475212619 >> 8)

	staticBuffer[141] = byte(v8076638462475212619 >> 16)

	staticBuffer[140] = byte(v8076638462475212619 >> 24)

	staticBuffer[139] = byte(v8076638462475212619 >> 32)

	staticBuffer[138] = byte(v8076638462475212619 >> 40)

	staticBuffer[137] = byte(v8076638462475212619 >> 48)

	staticBuffer[136] = byte(v8076638462475212619 >> 56)

	// 18 : Person.Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 144

	// Size : 2, Offset : 144, VarName : v.Address.Zip

	staticBuffer[145] = byte(v.Address.Zip)

	staticBuffer[144] = byte(v.Address.Zip >> 8)

	// 19 : Person.Age
	// Type : u8
	// Size : 1
	// Offset : 146

	// Size : 1, Offset : 146, VarName : v.Age

	staticBuffer[146] = byte(v.Age)

	w.Write(staticBuffer[:])

	// 20 : Person.Name
	// Type : str
	// Size : Variable
	// Offset : 0

	v11849227170113041263 := uint32(len(v.Name))

	// Size : 4, VarName : v11849227170113041263
	w.Write([]byte{

		byte(v11849227170113041263 >> 24),

		byte(v11849227170113041263 >> 16),

		byte(v11849227170113041263 >> 8),

		byte(v11849227170113041263),
	})

	w.Write([]byte(v.Name))

	// 21 : Person.Email
	// Type : str
	// Size : Variable
	// Offset : 0

	v3858853229782848418 := uint32(len(v.Email))

	// Size : 4, VarName : v3858853229782848418
	w.Write([]byte{

		byte(v3858853229782848418 >> 24),

		byte(v3858853229782848418 >> 16),

		byte(v3858853229782848418 >> 8),

		byte(v3858853229782848418),
	})

	w.Write([]byte(v.Email))

	// 22 : Person.Address.Street
	// Type : str
	// Size : Variable
	// Offset : 0

	v8409318322652509725 := uint32(len(v.Address.Street))

	// Size : 4, VarName : v8409318322652509725
	w.Write([]byte{

		byte(v8409318322652509725 >> 24),

		byte(v8409318322652509725 >> 16),

		byte(v8409318322652509725 >> 8),

		byte(v8409318322652509725),
	})

	w.Write([]byte(v.Address.Street))

	// 23 : Person.Address.City
	// Type : str
	// Size : Variable
	// Offset : 0

	v15591972827899983791 := uint32(len(v.Address.City))

	// Size : 4, VarName : v15591972827899983791
	w.Write([]byte{

		byte(v15591972827899983791 >> 24),

		byte(v15591972827899983791 >> 16),

		byte(v15591972827899983791 >> 8),

		byte(v15591972827899983791),
	})

	w.Write([]byte(v.Address.City))

	// 24 : Person.Address.State
	// Type : str
	// Size : Variable
	// Offset : 0

	v3453535492155706582 := uint32(len(v.Address.State))

	// Size : 4, VarName : v3453535492155706582
	w.Write([]byte{

		byte(v3453535492155706582 >> 24),

		byte(v3453535492155706582 >> 16),

		byte(v3453535492155706582 >> 8),

		byte(v3453535492155706582),
	})

	w.Write([]byte(v.Address.State))

	// 25 : Person.Phone.Number
	// Type : str
	// Size : Variable
	// Offset : 0

	v13507399295017256024 := uint32(len(v.Phone.Number))

	// Size : 4, VarName : v13507399295017256024
	w.Write([]byte{

		byte(v13507399295017256024 >> 24),

		byte(v13507399295017256024 >> 16),

		byte(v13507399295017256024 >> 8),

		byte(v13507399295017256024),
	})

	w.Write([]byte(v.Phone.Number))

	// 26 : Person.Phone.Type
	// Type : str
	// Size : Variable
	// Offset : 0

	v4972214672458315571 := uint32(len(v.Phone.Type))

	// Size : 4, VarName : v4972214672458315571
	w.Write([]byte{

		byte(v4972214672458315571 >> 24),

		byte(v4972214672458315571 >> 16),

		byte(v4972214672458315571 >> 8),

		byte(v4972214672458315571),
	})

	w.Write([]byte(v.Phone.Type))

	// 27 : Person.Acls[0]
	// Type : str
	// Size : Variable
	// Offset : 0

	v6836198594955968347 := uint32(len(v.Acls[0]))

	// Size : 4, VarName : v6836198594955968347
	w.Write([]byte{

		byte(v6836198594955968347 >> 24),

		byte(v6836198594955968347 >> 16),

		byte(v6836198594955968347 >> 8),

		byte(v6836198594955968347),
	})

	w.Write([]byte(v.Acls[0]))

	// 28 : Person.Acls[1]
	// Type : str
	// Size : Variable
	// Offset : 0

	v10426624325524632933 := uint32(len(v.Acls[1]))

	// Size : 4, VarName : v10426624325524632933
	w.Write([]byte{

		byte(v10426624325524632933 >> 24),

		byte(v10426624325524632933 >> 16),

		byte(v10426624325524632933 >> 8),

		byte(v10426624325524632933),
	})

	w.Write([]byte(v.Acls[1]))

	// 29 : Person.Acls[2]
	// Type : str
	// Size : Variable
	// Offset : 0

	v8205407213944833598 := uint32(len(v.Acls[2]))

	// Size : 4, VarName : v8205407213944833598
	w.Write([]byte{

		byte(v8205407213944833598 >> 24),

		byte(v8205407213944833598 >> 16),

		byte(v8205407213944833598 >> 8),

		byte(v8205407213944833598),
	})

	w.Write([]byte(v.Acls[2]))

	// 30 : Person.Acls[3]
	// Type : str
	// Size : Variable
	// Offset : 0

	v15017092115168473635 := uint32(len(v.Acls[3]))

	// Size : 4, VarName : v15017092115168473635
	w.Write([]byte{

		byte(v15017092115168473635 >> 24),

		byte(v15017092115168473635 >> 16),

		byte(v15017092115168473635 >> 8),

		byte(v15017092115168473635),
	})

	w.Write([]byte(v.Acls[3]))

	// 31 : Person.Acls[4]
	// Type : str
	// Size : Variable
	// Offset : 0

	v5816500411862272684 := uint32(len(v.Acls[4]))

	// Size : 4, VarName : v5816500411862272684
	w.Write([]byte{

		byte(v5816500411862272684 >> 24),

		byte(v5816500411862272684 >> 16),

		byte(v5816500411862272684 >> 8),

		byte(v5816500411862272684),
	})

	w.Write([]byte(v.Acls[4]))

	// 32 : Person.Acls[5]
	// Type : str
	// Size : Variable
	// Offset : 0

	v14723241295911999351 := uint32(len(v.Acls[5]))

	// Size : 4, VarName : v14723241295911999351
	w.Write([]byte{

		byte(v14723241295911999351 >> 24),

		byte(v14723241295911999351 >> 16),

		byte(v14723241295911999351 >> 8),

		byte(v14723241295911999351),
	})

	w.Write([]byte(v.Acls[5]))

	// 33 : Person.Acls[6]
	// Type : str
	// Size : Variable
	// Offset : 0

	v1935244031808836844 := uint32(len(v.Acls[6]))

	// Size : 4, VarName : v1935244031808836844
	w.Write([]byte{

		byte(v1935244031808836844 >> 24),

		byte(v1935244031808836844 >> 16),

		byte(v1935244031808836844 >> 8),

		byte(v1935244031808836844),
	})

	w.Write([]byte(v.Acls[6]))

	// 34 : Person.Acls[7]
	// Type : str
	// Size : Variable
	// Offset : 0

	v9997594027417585017 := uint32(len(v.Acls[7]))

	// Size : 4, VarName : v9997594027417585017
	w.Write([]byte{

		byte(v9997594027417585017 >> 24),

		byte(v9997594027417585017 >> 16),

		byte(v9997594027417585017 >> 8),

		byte(v9997594027417585017),
	})

	w.Write([]byte(v.Acls[7]))

	// 35 : Person.Acls[8]
	// Type : str
	// Size : Variable
	// Offset : 0

	v7247673547066349949 := uint32(len(v.Acls[8]))

	// Size : 4, VarName : v7247673547066349949
	w.Write([]byte{

		byte(v7247673547066349949 >> 24),

		byte(v7247673547066349949 >> 16),

		byte(v7247673547066349949 >> 8),

		byte(v7247673547066349949),
	})

	w.Write([]byte(v.Acls[8]))

	// 36 : Person.Acls[9]
	// Type : str
	// Size : Variable
	// Offset : 0

	v4280783872168215183 := uint32(len(v.Acls[9]))

	// Size : 4, VarName : v4280783872168215183
	w.Write([]byte{

		byte(v4280783872168215183 >> 24),

		byte(v4280783872168215183 >> 16),

		byte(v4280783872168215183 >> 8),

		byte(v4280783872168215183),
	})

	w.Write([]byte(v.Acls[9]))

	// 37 : Person.Acls[10]
	// Type : str
	// Size : Variable
	// Offset : 0

	v6002717496732854062 := uint32(len(v.Acls[10]))

	// Size : 4, VarName : v6002717496732854062
	w.Write([]byte{

		byte(v6002717496732854062 >> 24),

		byte(v6002717496732854062 >> 16),

		byte(v6002717496732854062 >> 8),

		byte(v6002717496732854062),
	})

	w.Write([]byte(v.Acls[10]))

	// 38 : Person.Acls[11]
	// Type : str
	// Size : Variable
	// Offset : 0

	v11257125322235012634 := uint32(len(v.Acls[11]))

	// Size : 4, VarName : v11257125322235012634
	w.Write([]byte{

		byte(v11257125322235012634 >> 24),

		byte(v11257125322235012634 >> 16),

		byte(v11257125322235012634 >> 8),

		byte(v11257125322235012634),
	})

	w.Write([]byte(v.Acls[11]))

	// 39 : Person.Acls[12]
	// Type : str
	// Size : Variable
	// Offset : 0

	v16173330515772890826 := uint32(len(v.Acls[12]))

	// Size : 4, VarName : v16173330515772890826
	w.Write([]byte{

		byte(v16173330515772890826 >> 24),

		byte(v16173330515772890826 >> 16),

		byte(v16173330515772890826 >> 8),

		byte(v16173330515772890826),
	})

	w.Write([]byte(v.Acls[12]))

	// 40 : Person.Acls[13]
	// Type : str
	// Size : Variable
	// Offset : 0

	v15321225512058676956 := uint32(len(v.Acls[13]))

	// Size : 4, VarName : v15321225512058676956
	w.Write([]byte{

		byte(v15321225512058676956 >> 24),

		byte(v15321225512058676956 >> 16),

		byte(v15321225512058676956 >> 8),

		byte(v15321225512058676956),
	})

	w.Write([]byte(v.Acls[13]))

	// 41 : Person.Acls[14]
	// Type : str
	// Size : Variable
	// Offset : 0

	v6885681762954042091 := uint32(len(v.Acls[14]))

	// Size : 4, VarName : v6885681762954042091
	w.Write([]byte{

		byte(v6885681762954042091 >> 24),

		byte(v6885681762954042091 >> 16),

		byte(v6885681762954042091 >> 8),

		byte(v6885681762954042091),
	})

	w.Write([]byte(v.Acls[14]))

	// 42 : Person.Acls[15]
	// Type : str
	// Size : Variable
	// Offset : 0

	v5484361586777816866 := uint32(len(v.Acls[15]))

	// Size : 4, VarName : v5484361586777816866
	w.Write([]byte{

		byte(v5484361586777816866 >> 24),

		byte(v5484361586777816866 >> 16),

		byte(v5484361586777816866 >> 8),

		byte(v5484361586777816866),
	})

	w.Write([]byte(v.Acls[15]))

}


func (v *Address) wt(w io.Writer) {

	var staticBuffer [2]byte

	// 0 : Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 0

	// Size : 2, Offset : 0, VarName : v.Zip

	staticBuffer[1] = byte(v.Zip)

	staticBuffer[0] = byte(v.Zip >> 8)

	w.Write(staticBuffer[:])

	// 1 : Address.Street
	// Type : str
	// Size : Variable
	// Offset : 0

	v18249806280722448238 := uint32(len(v.Street))

	// Size : 4, VarName : v18249806280722448238
	w.Write([]byte{

		byte(v18249806280722448238 >> 24),

		byte(v18249806280722448238 >> 16),

		byte(v18249806280722448238 >> 8),

		byte(v18249806280722448238),
	})

	w.Write([]byte(v.Street))

	// 2 : Address.City
	// Type : str
	// Size : Variable
	// Offset : 0

	v13944058276368568816 := uint32(len(v.City))

	// Size : 4, VarName : v13944058276368568816
	w.Write([]byte{

		byte(v13944058276368568816 >> 24),

		byte(v13944058276368568816 >> 16),

		byte(v13944058276368568816 >> 8),

		byte(v13944058276368568816),
	})

	w.Write([]byte(v.City))

	// 3 : Address.State
	// Type : str
	// Size : Variable
	// Offset : 0

	v2365195072861399743 := uint32(len(v.State))

	// Size : 4, VarName : v2365195072861399743
	w.Write([]byte{

		byte(v2365195072861399743 >> 24),

		byte(v2365195072861399743 >> 16),

		byte(v2365195072861399743 >> 8),

		byte(v2365195072861399743),
	})

	w.Write([]byte(v.State))

}


func (v *PhoneNumber) wt(w io.Writer) {

	// 0 : PhoneNumber.Number
	// Type : str
	// Size : Variable
	// Offset : 0

	v11589120410089113089 := uint32(len(v.Number))

	// Size : 4, VarName : v11589120410089113089
	w.Write([]byte{

		byte(v11589120410089113089 >> 24),

		byte(v11589120410089113089 >> 16),

		byte(v11589120410089113089 >> 8),

		byte(v11589120410089113089),
	})

	w.Write([]byte(v.Number))

	// 1 : PhoneNumber.Type
	// Type : str
	// Size : Variable
	// Offset : 0

	v10568967721202894828 := uint32(len(v.Type))

	// Size : 4, VarName : v10568967721202894828
	w.Write([]byte{

		byte(v10568967721202894828 >> 24),

		byte(v10568967721202894828 >> 16),

		byte(v10568967721202894828 >> 8),

		byte(v10568967721202894828),
	})

	w.Write([]byte(v.Type))

}

