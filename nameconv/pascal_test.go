package nameconv_test

import (
	"alrescha-go/nameconv"
	"testing"
)

func TestA(t *testing.T) {
	r := nameconv.Snake2Pascal("name_space")
	t.Log(r)
	if r == "NameSpace" {
		t.Log("TestA passed")
	} else {
		t.Error("TestA failed")
	}
}

func TestB(t *testing.T) {
	r := nameconv.Snake2Pascal("ursa_major")
	t.Log(r)
	if r == "UrsaMajor" {
		t.Log("TestA passed")
	} else {
		t.Error("TestA failed")
	}
}

func TestC(t *testing.T) {
	r := nameconv.Snake2Pascal("_do_convert")
	t.Log(r)
	if r == "DoConvert" {
		t.Log("TestA passed")
	} else {
		t.Error("TestA failed")
	}
}

func TestD(t *testing.T) {
	r := nameconv.Snake2Pascal("_why_convert_this")
	t.Log(r)
	if r == "WhyConvertThis" {
		t.Log("TestA passed")
	} else {
		t.Error("TestA failed")
	}
}
