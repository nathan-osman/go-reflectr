package reflectr

import (
	"errors"
	"testing"
)

var (
	errTest = errors.New("test error")
)

type testStruct struct{}

func (t testStruct) VMethod()  {}
func (t *testStruct) PMethod() {}

func (t *testStruct) Params(string, int, bool) {}

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
	if err := Struct(&testStruct{}).Method("VMethod").Error(); err != nil {
		t.Fatal(err)
	}
	if err := Struct(&testStruct{}).Method("PMethod").Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadParam(t *testing.T) {
	if err := Struct(&testStruct{}).Param(0, "").Error(); err != errMethodNotSelected {
		t.Fatalf("%v != %v", err, errMethodNotSelected)
	}
	if err := Struct(&testStruct{}).Method("Params").Param(3, "").Error(); err != errInvalidParamOffset {
		t.Fatalf("%v != %v", err, errInvalidParamOffset)
	}
	if err := Struct(&testStruct{}).Method("Params").Param(0, 42).Error(); err != errParamType {
		t.Fatalf("%v != %v", err, errParamType)
	}
}

func TestParam(t *testing.T) {
	if err := Struct(&testStruct{}).Method("Params").Param(0, "").Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadParams(t *testing.T) {
	if err := Struct(&testStruct{}).Params("", 0, false).Error(); err != errMethodNotSelected {
		t.Fatalf("%v != %v", err, errMethodNotSelected)
	}
	if err := Struct(&testStruct{}).Method("Params").Params("", 0).Error(); err != errParamCount {
		t.Fatalf("%v != %v", err, errParamCount)
	}
	if err := Struct(&testStruct{}).Method("Params").Params("", 0, 0).Error(); err != errParamType {
		t.Fatalf("%v != %v", err, errParamType)
	}
}

func TestParams(t *testing.T) {
	if err := Struct(&testStruct{}).Method("Params").Params("", 0, false).Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadReturn(t *testing.T) {
	if err := Struct(&testStruct{}).Return(0, "").Error(); err != errMethodNotSelected {
		t.Fatalf("%v != %v", err, errMethodNotSelected)
	}
	if err := Struct(&testStruct{}).Method("Returns").Return(3, "").Error(); err != errInvalidReturnOffset {
		t.Fatalf("%v != %v", err, errInvalidReturnOffset)
	}
	if err := Struct(&testStruct{}).Method("Returns").Return(0, 42).Error(); err != errReturnType {
		t.Fatalf("%v != %v", err, errReturnType)
	}
}

func TestReturn(t *testing.T) {
	if err := Struct(&testStruct{}).Method("Returns").Return(0, "").Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadReturns(t *testing.T) {
	if err := Struct(&testStruct{}).Returns("", 0, false).Error(); err != errMethodNotSelected {
		t.Fatalf("%v != %v", err, errMethodNotSelected)
	}
	if err := Struct(&testStruct{}).Method("Returns").Returns("", 0).Error(); err != errReturnCount {
		t.Fatalf("%v != %v", err, errReturnCount)
	}
	if err := Struct(&testStruct{}).Method("Returns").Returns("", 0, 0).Error(); err != errReturnType {
		t.Fatalf("%v != %v", err, errReturnType)
	}
}

func TestReturns(t *testing.T) {
	if err := Struct(&testStruct{}).Method("Returns").Returns("", 0, false).Error(); err != nil {
		t.Fatal(err)
	}
}

func TestError(t *testing.T) {
	s := Struct(&testStruct{})
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
