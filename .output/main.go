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

	Phone []PhoneNumber

	Email string

	Acls []string

	LastConnection float64

	LastUpdated float32
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

func (v *Person) GetPhone() []PhoneNumber {
	return v.Phone
}

func (v *Person) SetPhone(value []PhoneNumber) {
	v.Phone = value
}

func (v *Person) GetEmail() string {
	return v.Email
}

func (v *Person) SetEmail(value string) {
	v.Email = value
}

func (v *Person) GetAcls() []string {
	return v.Acls
}

func (v *Person) SetAcls(value []string) {
	v.Acls = value
}

func (v *Person) GetLastConnection() float64 {
	return v.LastConnection
}

func (v *Person) SetLastConnection(value float64) {
	v.LastConnection = value
}

func (v *Person) GetLastUpdated() float32 {
	return v.LastUpdated
}

func (v *Person) SetLastUpdated(value float32) {
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

func (v *Person) wt(w io.Writer) error {
	var err error

	var staticBuffer [15]byte

	// 0 : Person.LastConnection
	// Type : f64
	// Size : 8
	// Offset : 0

	v2886259293251357746 := math.Float64bits(v.LastConnection)

	// Size : 8, Offset : 0, VarName : v2886259293251357746

	staticBuffer[7] = byte(v2886259293251357746)

	staticBuffer[6] = byte(v2886259293251357746 >> 8)

	staticBuffer[5] = byte(v2886259293251357746 >> 16)

	staticBuffer[4] = byte(v2886259293251357746 >> 24)

	staticBuffer[3] = byte(v2886259293251357746 >> 32)

	staticBuffer[2] = byte(v2886259293251357746 >> 40)

	staticBuffer[1] = byte(v2886259293251357746 >> 48)

	staticBuffer[0] = byte(v2886259293251357746 >> 56)

	// 1 : Person.LastUpdated
	// Type : f32
	// Size : 4
	// Offset : 8

	v11300417155691566972 := math.Float32bits(v.LastUpdated)

	// Size : 4, Offset : 8, VarName : v11300417155691566972

	staticBuffer[11] = byte(v11300417155691566972)

	staticBuffer[10] = byte(v11300417155691566972 >> 8)

	staticBuffer[9] = byte(v11300417155691566972 >> 16)

	staticBuffer[8] = byte(v11300417155691566972 >> 24)

	// 2 : Person.Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 12

	// Size : 2, Offset : 12, VarName : v.Address.Zip

	staticBuffer[13] = byte(v.Address.Zip)

	staticBuffer[12] = byte(v.Address.Zip >> 8)

	// 3 : Person.Age
	// Type : u8
	// Size : 1
	// Offset : 14

	// Size : 1, Offset : 14, VarName : v.Age

	staticBuffer[14] = byte(v.Age)

	_, err = w.Write(staticBuffer[:])
	if err != nil {
		return err
	}

	// 4 : Person.Name
	// Type : str
	// Size : Variable

	v390690336615699554 := uint32(len(v.Name))

	// Size : 4, VarName : v390690336615699554
	_, err = w.Write([]byte{

		byte(v390690336615699554 >> 24),

		byte(v390690336615699554 >> 16),

		byte(v390690336615699554 >> 8),

		byte(v390690336615699554),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Name))
	if err != nil {
		return err
	}

	// 5 : Person.Phone
	// Type : phone_number (List)
	// Size : Variable

	v17713352639180781079 := uint32(len(v.Phone))

	// Size : 4, VarName : v17713352639180781079
	_, err = w.Write([]byte{

		byte(v17713352639180781079 >> 24),

		byte(v17713352639180781079 >> 16),

		byte(v17713352639180781079 >> 8),

		byte(v17713352639180781079),
	})
	if err != nil {
		return err
	}

	// == List ==
	// List VarName : v.Phone
	// List VarType : phone_number
	for i := range v.Phone {

		v.Phone[i].wt(w)

	}

	// 6 : Person.Email
	// Type : str
	// Size : Variable

	v8750545231426907830 := uint32(len(v.Email))

	// Size : 4, VarName : v8750545231426907830
	_, err = w.Write([]byte{

		byte(v8750545231426907830 >> 24),

		byte(v8750545231426907830 >> 16),

		byte(v8750545231426907830 >> 8),

		byte(v8750545231426907830),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Email))
	if err != nil {
		return err
	}

	// 7 : Person.Acls
	// Type : str (List)
	// Size : Variable

	v9636894019407427379 := uint32(len(v.Acls))

	// Size : 4, VarName : v9636894019407427379
	_, err = w.Write([]byte{

		byte(v9636894019407427379 >> 24),

		byte(v9636894019407427379 >> 16),

		byte(v9636894019407427379 >> 8),

		byte(v9636894019407427379),
	})
	if err != nil {
		return err
	}

	// == List ==
	// List VarName : v.Acls
	// List VarType : str
	for i := range v.Acls {

		v4615704643516027163 := uint32(len(v.Acls[i]))

		// Size : 4, VarName : v4615704643516027163
		_, err = w.Write([]byte{

			byte(v4615704643516027163 >> 24),

			byte(v4615704643516027163 >> 16),

			byte(v4615704643516027163 >> 8),

			byte(v4615704643516027163),
		})
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(v.Acls[i]))
		if err != nil {
			return err
		}

	}

	// 8 : Person.Address.Street
	// Type : str
	// Size : Variable

	v16952515822426794351 := uint32(len(v.Address.Street))

	// Size : 4, VarName : v16952515822426794351
	_, err = w.Write([]byte{

		byte(v16952515822426794351 >> 24),

		byte(v16952515822426794351 >> 16),

		byte(v16952515822426794351 >> 8),

		byte(v16952515822426794351),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Address.Street))
	if err != nil {
		return err
	}

	// 9 : Person.Address.City
	// Type : str
	// Size : Variable

	v7891621465354055997 := uint32(len(v.Address.City))

	// Size : 4, VarName : v7891621465354055997
	_, err = w.Write([]byte{

		byte(v7891621465354055997 >> 24),

		byte(v7891621465354055997 >> 16),

		byte(v7891621465354055997 >> 8),

		byte(v7891621465354055997),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Address.City))
	if err != nil {
		return err
	}

	// 10 : Person.Address.State
	// Type : str
	// Size : Variable

	v10308970753390518818 := uint32(len(v.Address.State))

	// Size : 4, VarName : v10308970753390518818
	_, err = w.Write([]byte{

		byte(v10308970753390518818 >> 24),

		byte(v10308970753390518818 >> 16),

		byte(v10308970753390518818 >> 8),

		byte(v10308970753390518818),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Address.State))
	if err != nil {
		return err
	}

	return err
}

func (v *Person) rf(r io.Reader) error {
	var err error

	var staticBuffer [15]byte
	_, err = r.Read(staticBuffer[:])
	if err != nil {
		return err
	}

	// 0 : Person.LastConnection
	// Type : f64
	// Size : 8
	// Offset : 0

	var v13771331064356464926 uint64

	// Size : 8, Offset : 0, VarName : v13771331064356464926
	var v15811822008244594417 uint64

	v15811822008244594417 |= uint64(staticBuffer[7]) << 0

	v15811822008244594417 |= uint64(staticBuffer[6]) << 8

	v15811822008244594417 |= uint64(staticBuffer[5]) << 16

	v15811822008244594417 |= uint64(staticBuffer[4]) << 24

	v15811822008244594417 |= uint64(staticBuffer[3]) << 32

	v15811822008244594417 |= uint64(staticBuffer[2]) << 40

	v15811822008244594417 |= uint64(staticBuffer[1]) << 48

	v15811822008244594417 |= uint64(staticBuffer[0]) << 56

	v13771331064356464926 = v15811822008244594417

	v.LastConnection = math.Float64frombits(v13771331064356464926)

	// 1 : Person.LastUpdated
	// Type : f32
	// Size : 4
	// Offset : 8

	var v9675428200009801379 uint32

	// Size : 4, Offset : 8, VarName : v9675428200009801379
	var v14373855068129063513 uint32

	v14373855068129063513 |= uint32(staticBuffer[11]) << 0

	v14373855068129063513 |= uint32(staticBuffer[10]) << 8

	v14373855068129063513 |= uint32(staticBuffer[9]) << 16

	v14373855068129063513 |= uint32(staticBuffer[8]) << 24

	v9675428200009801379 = v14373855068129063513

	v.LastUpdated = math.Float32frombits(v9675428200009801379)

	// 2 : Person.Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 12

	// Size : 2, Offset : 12, VarName : v.Address.Zip
	var v17653918525122512866 uint16

	v17653918525122512866 |= uint16(staticBuffer[13]) << 0

	v17653918525122512866 |= uint16(staticBuffer[12]) << 8

	v.Address.Zip = v17653918525122512866

	// 3 : Person.Age
	// Type : u8
	// Size : 1
	// Offset : 14

	// Size : 1, Offset : 14, VarName : v.Age
	var v14139199562580759745 uint8

	v14139199562580759745 |= uint8(staticBuffer[14]) << 0

	v.Age = v14139199562580759745

	// 4 : Person.Name
	// Type : str
	// Size : Variable

	var v15850420683481224815 uint32

	// Size : 4, VarName : v15850420683481224815
	var v12536694065450811718 uint32
	var Buffer12536694065450811718 [4]byte
	_, err = r.Read(Buffer12536694065450811718[:])
	if err != nil {
		return err
	}

	v12536694065450811718 |= uint32(Buffer12536694065450811718[3]) << 0

	v12536694065450811718 |= uint32(Buffer12536694065450811718[2]) << 8

	v12536694065450811718 |= uint32(Buffer12536694065450811718[1]) << 16

	v12536694065450811718 |= uint32(Buffer12536694065450811718[0]) << 24

	v15850420683481224815 = v12536694065450811718

	var Buffer15850420683481224815 []byte = make([]byte, v15850420683481224815)

	_, err = r.Read(Buffer15850420683481224815)
	if err != nil {
		return err
	}
	v.Name = string(Buffer15850420683481224815)

	// 5 : Person.Phone
	// Type : phone_number (List)
	// Size : Variable

	var v1415735174609993392 uint32

	// Size : 4, VarName : v1415735174609993392
	var v16399176508491325905 uint32
	var Buffer16399176508491325905 [4]byte
	_, err = r.Read(Buffer16399176508491325905[:])
	if err != nil {
		return err
	}

	v16399176508491325905 |= uint32(Buffer16399176508491325905[3]) << 0

	v16399176508491325905 |= uint32(Buffer16399176508491325905[2]) << 8

	v16399176508491325905 |= uint32(Buffer16399176508491325905[1]) << 16

	v16399176508491325905 |= uint32(Buffer16399176508491325905[0]) << 24

	v1415735174609993392 = v16399176508491325905

	// == List ==
	// List VarName : v.Phone
	// List VarType : phone_number
	v16414938421894519138 := int(v1415735174609993392)
	if len(v.Phone) < v16414938421894519138 {
		v.Phone = make([]PhoneNumber, v1415735174609993392)
	}
	for i := 0; i < v16414938421894519138; i++ {

		v.Phone[i].rf(r)

	}

	// 6 : Person.Email
	// Type : str
	// Size : Variable

	var v13628614141203813367 uint32

	// Size : 4, VarName : v13628614141203813367
	var v13903310095902045678 uint32
	var Buffer13903310095902045678 [4]byte
	_, err = r.Read(Buffer13903310095902045678[:])
	if err != nil {
		return err
	}

	v13903310095902045678 |= uint32(Buffer13903310095902045678[3]) << 0

	v13903310095902045678 |= uint32(Buffer13903310095902045678[2]) << 8

	v13903310095902045678 |= uint32(Buffer13903310095902045678[1]) << 16

	v13903310095902045678 |= uint32(Buffer13903310095902045678[0]) << 24

	v13628614141203813367 = v13903310095902045678

	var Buffer13628614141203813367 []byte = make([]byte, v13628614141203813367)

	_, err = r.Read(Buffer13628614141203813367)
	if err != nil {
		return err
	}
	v.Email = string(Buffer13628614141203813367)

	// 7 : Person.Acls
	// Type : str (List)
	// Size : Variable

	var v9529777818909013126 uint32

	// Size : 4, VarName : v9529777818909013126
	var v13529413298993747006 uint32
	var Buffer13529413298993747006 [4]byte
	_, err = r.Read(Buffer13529413298993747006[:])
	if err != nil {
		return err
	}

	v13529413298993747006 |= uint32(Buffer13529413298993747006[3]) << 0

	v13529413298993747006 |= uint32(Buffer13529413298993747006[2]) << 8

	v13529413298993747006 |= uint32(Buffer13529413298993747006[1]) << 16

	v13529413298993747006 |= uint32(Buffer13529413298993747006[0]) << 24

	v9529777818909013126 = v13529413298993747006

	// == List ==
	// List VarName : v.Acls
	// List VarType : str
	v2938392810170085288 := int(v9529777818909013126)
	if len(v.Acls) < v2938392810170085288 {
		v.Acls = make([]string, v9529777818909013126)
	}
	for i := 0; i < v2938392810170085288; i++ {

		var v9507258741087743219 uint32

		// Size : 4, VarName : v9507258741087743219
		var v10610551275313405629 uint32
		var Buffer10610551275313405629 [4]byte
		_, err = r.Read(Buffer10610551275313405629[:])
		if err != nil {
			return err
		}

		v10610551275313405629 |= uint32(Buffer10610551275313405629[3]) << 0

		v10610551275313405629 |= uint32(Buffer10610551275313405629[2]) << 8

		v10610551275313405629 |= uint32(Buffer10610551275313405629[1]) << 16

		v10610551275313405629 |= uint32(Buffer10610551275313405629[0]) << 24

		v9507258741087743219 = v10610551275313405629

		var Buffer9507258741087743219 []byte = make([]byte, v9507258741087743219)

		_, err = r.Read(Buffer9507258741087743219)
		if err != nil {
			return err
		}
		v.Acls[i] = string(Buffer9507258741087743219)

	}

	// 8 : Person.Address.Street
	// Type : str
	// Size : Variable

	var v7260754940691539846 uint32

	// Size : 4, VarName : v7260754940691539846
	var v3535989811736617975 uint32
	var Buffer3535989811736617975 [4]byte
	_, err = r.Read(Buffer3535989811736617975[:])
	if err != nil {
		return err
	}

	v3535989811736617975 |= uint32(Buffer3535989811736617975[3]) << 0

	v3535989811736617975 |= uint32(Buffer3535989811736617975[2]) << 8

	v3535989811736617975 |= uint32(Buffer3535989811736617975[1]) << 16

	v3535989811736617975 |= uint32(Buffer3535989811736617975[0]) << 24

	v7260754940691539846 = v3535989811736617975

	var Buffer7260754940691539846 []byte = make([]byte, v7260754940691539846)

	_, err = r.Read(Buffer7260754940691539846)
	if err != nil {
		return err
	}
	v.Address.Street = string(Buffer7260754940691539846)

	// 9 : Person.Address.City
	// Type : str
	// Size : Variable

	var v7235065057497506816 uint32

	// Size : 4, VarName : v7235065057497506816
	var v16821742904349168785 uint32
	var Buffer16821742904349168785 [4]byte
	_, err = r.Read(Buffer16821742904349168785[:])
	if err != nil {
		return err
	}

	v16821742904349168785 |= uint32(Buffer16821742904349168785[3]) << 0

	v16821742904349168785 |= uint32(Buffer16821742904349168785[2]) << 8

	v16821742904349168785 |= uint32(Buffer16821742904349168785[1]) << 16

	v16821742904349168785 |= uint32(Buffer16821742904349168785[0]) << 24

	v7235065057497506816 = v16821742904349168785

	var Buffer7235065057497506816 []byte = make([]byte, v7235065057497506816)

	_, err = r.Read(Buffer7235065057497506816)
	if err != nil {
		return err
	}
	v.Address.City = string(Buffer7235065057497506816)

	// 10 : Person.Address.State
	// Type : str
	// Size : Variable

	var v11189147022576886905 uint32

	// Size : 4, VarName : v11189147022576886905
	var v8086879863351031716 uint32
	var Buffer8086879863351031716 [4]byte
	_, err = r.Read(Buffer8086879863351031716[:])
	if err != nil {
		return err
	}

	v8086879863351031716 |= uint32(Buffer8086879863351031716[3]) << 0

	v8086879863351031716 |= uint32(Buffer8086879863351031716[2]) << 8

	v8086879863351031716 |= uint32(Buffer8086879863351031716[1]) << 16

	v8086879863351031716 |= uint32(Buffer8086879863351031716[0]) << 24

	v11189147022576886905 = v8086879863351031716

	var Buffer11189147022576886905 []byte = make([]byte, v11189147022576886905)

	_, err = r.Read(Buffer11189147022576886905)
	if err != nil {
		return err
	}
	v.Address.State = string(Buffer11189147022576886905)

	return err
}

func (v *Address) wt(w io.Writer) error {
	var err error

	var staticBuffer [2]byte

	// 0 : Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 0

	// Size : 2, Offset : 0, VarName : v.Zip

	staticBuffer[1] = byte(v.Zip)

	staticBuffer[0] = byte(v.Zip >> 8)

	_, err = w.Write(staticBuffer[:])
	if err != nil {
		return err
	}

	// 1 : Address.Street
	// Type : str
	// Size : Variable

	v17903199866470059336 := uint32(len(v.Street))

	// Size : 4, VarName : v17903199866470059336
	_, err = w.Write([]byte{

		byte(v17903199866470059336 >> 24),

		byte(v17903199866470059336 >> 16),

		byte(v17903199866470059336 >> 8),

		byte(v17903199866470059336),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Street))
	if err != nil {
		return err
	}

	// 2 : Address.City
	// Type : str
	// Size : Variable

	v6796448128511569708 := uint32(len(v.City))

	// Size : 4, VarName : v6796448128511569708
	_, err = w.Write([]byte{

		byte(v6796448128511569708 >> 24),

		byte(v6796448128511569708 >> 16),

		byte(v6796448128511569708 >> 8),

		byte(v6796448128511569708),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.City))
	if err != nil {
		return err
	}

	// 3 : Address.State
	// Type : str
	// Size : Variable

	v5459143683193012010 := uint32(len(v.State))

	// Size : 4, VarName : v5459143683193012010
	_, err = w.Write([]byte{

		byte(v5459143683193012010 >> 24),

		byte(v5459143683193012010 >> 16),

		byte(v5459143683193012010 >> 8),

		byte(v5459143683193012010),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.State))
	if err != nil {
		return err
	}

	return err
}

func (v *Address) rf(r io.Reader) error {
	var err error

	var staticBuffer [2]byte
	_, err = r.Read(staticBuffer[:])
	if err != nil {
		return err
	}

	// 0 : Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 0

	// Size : 2, Offset : 0, VarName : v.Zip
	var v10711880795426572879 uint16

	v10711880795426572879 |= uint16(staticBuffer[1]) << 0

	v10711880795426572879 |= uint16(staticBuffer[0]) << 8

	v.Zip = v10711880795426572879

	// 1 : Address.Street
	// Type : str
	// Size : Variable

	var v13950793840491554029 uint32

	// Size : 4, VarName : v13950793840491554029
	var v15102258315840966967 uint32
	var Buffer15102258315840966967 [4]byte
	_, err = r.Read(Buffer15102258315840966967[:])
	if err != nil {
		return err
	}

	v15102258315840966967 |= uint32(Buffer15102258315840966967[3]) << 0

	v15102258315840966967 |= uint32(Buffer15102258315840966967[2]) << 8

	v15102258315840966967 |= uint32(Buffer15102258315840966967[1]) << 16

	v15102258315840966967 |= uint32(Buffer15102258315840966967[0]) << 24

	v13950793840491554029 = v15102258315840966967

	var Buffer13950793840491554029 []byte = make([]byte, v13950793840491554029)

	_, err = r.Read(Buffer13950793840491554029)
	if err != nil {
		return err
	}
	v.Street = string(Buffer13950793840491554029)

	// 2 : Address.City
	// Type : str
	// Size : Variable

	var v13444170626975976623 uint32

	// Size : 4, VarName : v13444170626975976623
	var v10514028885871222810 uint32
	var Buffer10514028885871222810 [4]byte
	_, err = r.Read(Buffer10514028885871222810[:])
	if err != nil {
		return err
	}

	v10514028885871222810 |= uint32(Buffer10514028885871222810[3]) << 0

	v10514028885871222810 |= uint32(Buffer10514028885871222810[2]) << 8

	v10514028885871222810 |= uint32(Buffer10514028885871222810[1]) << 16

	v10514028885871222810 |= uint32(Buffer10514028885871222810[0]) << 24

	v13444170626975976623 = v10514028885871222810

	var Buffer13444170626975976623 []byte = make([]byte, v13444170626975976623)

	_, err = r.Read(Buffer13444170626975976623)
	if err != nil {
		return err
	}
	v.City = string(Buffer13444170626975976623)

	// 3 : Address.State
	// Type : str
	// Size : Variable

	var v8716053268009330450 uint32

	// Size : 4, VarName : v8716053268009330450
	var v17631054214499992798 uint32
	var Buffer17631054214499992798 [4]byte
	_, err = r.Read(Buffer17631054214499992798[:])
	if err != nil {
		return err
	}

	v17631054214499992798 |= uint32(Buffer17631054214499992798[3]) << 0

	v17631054214499992798 |= uint32(Buffer17631054214499992798[2]) << 8

	v17631054214499992798 |= uint32(Buffer17631054214499992798[1]) << 16

	v17631054214499992798 |= uint32(Buffer17631054214499992798[0]) << 24

	v8716053268009330450 = v17631054214499992798

	var Buffer8716053268009330450 []byte = make([]byte, v8716053268009330450)

	_, err = r.Read(Buffer8716053268009330450)
	if err != nil {
		return err
	}
	v.State = string(Buffer8716053268009330450)

	return err
}

func (v *PhoneNumber) wt(w io.Writer) error {
	var err error

	// 0 : PhoneNumber.Number
	// Type : str
	// Size : Variable

	v5370204494138738806 := uint32(len(v.Number))

	// Size : 4, VarName : v5370204494138738806
	_, err = w.Write([]byte{

		byte(v5370204494138738806 >> 24),

		byte(v5370204494138738806 >> 16),

		byte(v5370204494138738806 >> 8),

		byte(v5370204494138738806),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Number))
	if err != nil {
		return err
	}

	// 1 : PhoneNumber.Type
	// Type : str
	// Size : Variable

	v10993859955262981653 := uint32(len(v.Type))

	// Size : 4, VarName : v10993859955262981653
	_, err = w.Write([]byte{

		byte(v10993859955262981653 >> 24),

		byte(v10993859955262981653 >> 16),

		byte(v10993859955262981653 >> 8),

		byte(v10993859955262981653),
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(v.Type))
	if err != nil {
		return err
	}

	return err
}

func (v *PhoneNumber) rf(r io.Reader) error {
	var err error

	// 0 : PhoneNumber.Number
	// Type : str
	// Size : Variable

	var v209701120546619398 uint32

	// Size : 4, VarName : v209701120546619398
	var v8476012184825072548 uint32
	var Buffer8476012184825072548 [4]byte
	_, err = r.Read(Buffer8476012184825072548[:])
	if err != nil {
		return err
	}

	v8476012184825072548 |= uint32(Buffer8476012184825072548[3]) << 0

	v8476012184825072548 |= uint32(Buffer8476012184825072548[2]) << 8

	v8476012184825072548 |= uint32(Buffer8476012184825072548[1]) << 16

	v8476012184825072548 |= uint32(Buffer8476012184825072548[0]) << 24

	v209701120546619398 = v8476012184825072548

	var Buffer209701120546619398 []byte = make([]byte, v209701120546619398)

	_, err = r.Read(Buffer209701120546619398)
	if err != nil {
		return err
	}
	v.Number = string(Buffer209701120546619398)

	// 1 : PhoneNumber.Type
	// Type : str
	// Size : Variable

	var v16632206138779083865 uint32

	// Size : 4, VarName : v16632206138779083865
	var v15302526146674420988 uint32
	var Buffer15302526146674420988 [4]byte
	_, err = r.Read(Buffer15302526146674420988[:])
	if err != nil {
		return err
	}

	v15302526146674420988 |= uint32(Buffer15302526146674420988[3]) << 0

	v15302526146674420988 |= uint32(Buffer15302526146674420988[2]) << 8

	v15302526146674420988 |= uint32(Buffer15302526146674420988[1]) << 16

	v15302526146674420988 |= uint32(Buffer15302526146674420988[0]) << 24

	v16632206138779083865 = v15302526146674420988

	var Buffer16632206138779083865 []byte = make([]byte, v16632206138779083865)

	_, err = r.Read(Buffer16632206138779083865)
	if err != nil {
		return err
	}
	v.Type = string(Buffer16632206138779083865)

	return err
}
// Sun Aug  8 02:45:02 UTC 2021
