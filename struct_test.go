package reflectr

import (
	"errors"
	"testing"
)

const (
	strTest = "test"
	intTest = 42
)

var (
	errTest = errors.New("test")
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

func TestIsPtr(t *testing.T) {
	if Struct(testStructStruct{}).IsPtr() == true {
		t.Fatal("IsPtr returned true")
	}
	if Struct(&testStructStruct{}).IsPtr() == false {
		t.Fatal("IsPtr returned false")
	}
}

func TestError(t *testing.T) {
	s := Struct(&testStructStruct{})
	s.err = errTest
	if _, err := s.
		Method("42").
		Param(3, "").
		Params().
		Return(3, "").
		Returns().
		Field("42").
		Type("").
		SetValue("").
		Value(); err != errTest {
		t.Fatalf("%v != %v", err, errTest)
	}
}
