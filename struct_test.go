package reflectr

import (
	"testing"
)

type testStruct struct{}

func (t testStruct) VMethod()  {}
func (t *testStruct) PMethod() {}

func (t *testStruct) Returns() (string, int, bool) {
	return "", 0, false
}

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

func TestBadReturns(t *testing.T) {
	if err := Struct(&testStruct{}).Returns("", 0, false).Error(); err != errMethodNotSelected {
		t.Fatalf("%v != %v", err, errMethodNotSelected)
	}
	if err := Struct(&testStruct{}).Method("Returns").Returns("", 0).Error(); err != errReturnsCount {
		t.Fatalf("%v != %v", err, errReturnsCount)
	}
	if err := Struct(&testStruct{}).Method("Returns").Returns("", 0, 0).Error(); err != errReturnsType {
		t.Fatalf("%v != %v", err, errReturnsType)
	}
}

func TestReturns(t *testing.T) {
	if err := Struct(&testStruct{}).Method("Returns").Returns("", 0, false).Error(); err != nil {
		t.Fatal(err)
	}
}
