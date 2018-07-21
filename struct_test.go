package reflectr

import (
	"testing"
)

type testStruct struct{}

func (t testStruct) VMethod()  {}
func (t *testStruct) PMethod() {}

func TestBadStruct(t *testing.T) {
	if err := Struct(42).Error(); err != errMustBeStruct {
		t.Fatalf("%v != %v", err, errMustBeStruct)
	}
}

func TestStruct(t *testing.T) {
	if err := Struct(testStruct{}).Error(); err != nil {
		t.Fatal(err)
	}
	if err := Struct(&testStruct{}).Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadMethod(t *testing.T) {
	if err := Struct(&testStruct{}).Method("42").Error(); err != errMethodDoesNotExist {
		t.Fatalf("%v != %v", err, errMethodDoesNotExist)
	}
}

func TestMethod(t *testing.T) {
	if err := Struct(testStruct{}).Method("VMethod").Error(); err != nil {
		t.Fatal(err)
	}
	for _, m := range []string{"VMethod", "PMethod"} {
		if err := Struct(&testStruct{}).Method(m).Error(); err != nil {
			t.Fatal(err)
		}
	}
}
