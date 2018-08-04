package reflectr

import (
	"testing"
)

type testStructStruct struct{}

func TestBadStruct(t *testing.T) {
	if err := Struct(42).Error(); err != errMustBeStruct {
		t.Fatalf("%v != %v", err, errMustBeStruct)
	}
}

func TestStruct(t *testing.T) {
	if err := Struct(testStructStruct{}).Error(); err != nil {
		t.Fatal(err)
	}
	if err := Struct(&testStructStruct{}).Error(); err != nil {
		t.Fatal(err)
	}
}

func TestError(t *testing.T) {
	s := Struct(&testStructStruct{})
	s.err = errTest
	if err := s.
		Method("42").
		Param(3, "").
		Params().
		Return(3, "").
		Returns().
		Error(); err != errTest {
		t.Fatalf("%v != %v", err, errTest)
	}
}
