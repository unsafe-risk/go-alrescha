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

	var staticBuffer [19]byte

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

	// 2 : Person.Address.Zip
	// Type : u16
	// Size : 2
	// Offset : 16

	// Size : 2, Offset : 16, VarName : v.Address.Zip

	staticBuffer[17] = byte(v.Address.Zip)

	staticBuffer[16] = byte(v.Address.Zip >> 8)

	// 3 : Person.Age
	// Type : u8
	// Size : 1
	// Offset : 18

	// Size : 1, Offset : 18, VarName : v.Age

	staticBuffer[18] = byte(v.Age)

	w.Write(staticBuffer[:])

	// 4 : Person.Name
	// Type : str
	// Size : Variable

	v11849227170113041263 := uint32(len(v.Name))

	// Size : 4, VarName : v11849227170113041263
	w.Write([]byte{

		byte(v11849227170113041263 >> 24),

		byte(v11849227170113041263 >> 16),

		byte(v11849227170113041263 >> 8),

		byte(v11849227170113041263),
	})

	w.Write([]byte(v.Name))

	// 5 : Person.Phone
	// Type : phone_number (List)
	// Size : Variable

	v12293774587795065081 := uint32(len(v.Phone))

	// Size : 4, VarName : v12293774587795065081
	w.Write([]byte{

		byte(v12293774587795065081 >> 24),

		byte(v12293774587795065081 >> 16),

		byte(v12293774587795065081 >> 8),

		byte(v12293774587795065081),
	})

	// == List ==
	// List VarName : v.Phone
	// List VarType : phone_number
	for i := range v.Phone {

		v.Phone[i].wt(w)

	}

	// 6 : Person.Email
	// Type : str
	// Size : Variable

	v3858853229782848418 := uint32(len(v.Email))

	// Size : 4, VarName : v3858853229782848418
	w.Write([]byte{

		byte(v3858853229782848418 >> 24),

		byte(v3858853229782848418 >> 16),

		byte(v3858853229782848418 >> 8),

		byte(v3858853229782848418),
	})

	w.Write([]byte(v.Email))

	// 7 : Person.Acls
	// Type : str (List)
	// Size : Variable

	v11719081467165893116 := uint32(len(v.Acls))

	// Size : 4, VarName : v11719081467165893116
	w.Write([]byte{

		byte(v11719081467165893116 >> 24),

		byte(v11719081467165893116 >> 16),

		byte(v11719081467165893116 >> 8),

		byte(v11719081467165893116),
	})

	// == List ==
	// List VarName : v.Acls
	// List VarType : str
	for i := range v.Acls {

		v86177424608865347 := uint32(len(v.Acls[i]))

		// Size : 4, VarName : v86177424608865347
		w.Write([]byte{

			byte(v86177424608865347 >> 24),

			byte(v86177424608865347 >> 16),

			byte(v86177424608865347 >> 8),

			byte(v86177424608865347),
		})

		w.Write([]byte(v.Acls[i]))

	}

	// 8 : Person.Address.Street
	// Type : str
	// Size : Variable

	v8409318322652509725 := uint32(len(v.Address.Street))

	// Size : 4, VarName : v8409318322652509725
	w.Write([]byte{

		byte(v8409318322652509725 >> 24),

		byte(v8409318322652509725 >> 16),

		byte(v8409318322652509725 >> 8),

		byte(v8409318322652509725),
	})

	w.Write([]byte(v.Address.Street))

	// 9 : Person.Address.City
	// Type : str
	// Size : Variable

	v15591972827899983791 := uint32(len(v.Address.City))

	// Size : 4, VarName : v15591972827899983791
	w.Write([]byte{

		byte(v15591972827899983791 >> 24),

		byte(v15591972827899983791 >> 16),

		byte(v15591972827899983791 >> 8),

		byte(v15591972827899983791),
	})

	w.Write([]byte(v.Address.City))

	// 10 : Person.Address.State
	// Type : str
	// Size : Variable

	v3453535492155706582 := uint32(len(v.Address.State))

	// Size : 4, VarName : v3453535492155706582
	w.Write([]byte{

		byte(v3453535492155706582 >> 24),

		byte(v3453535492155706582 >> 16),

		byte(v3453535492155706582 >> 8),

		byte(v3453535492155706582),
	})

	w.Write([]byte(v.Address.State))

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
// Tue Jul 27 23:10:02 UTC 2021
