package reflectr

import (
	"testing"
)

type testBadMethodStruct struct{}

func TestBadMethod(t *testing.T) {
	if err := Struct(&testBadMethodStruct{}).
		Method("42").
		Error(); err != ErrMethodDoesNotExist {
		t.Fatalf("%v != %v", err, ErrMethodDoesNotExist)
	}
}

type testMethodStruct struct{}

func (testMethodStruct) VMethod()  {}
func (*testMethodStruct) PMethod() {}

func TestMethod(t *testing.T) {
	if err := Struct(testMethodStruct{}).
		Method("VMethod").
		Error(); err != nil {
		t.Fatal(err)
	}
	if err := Struct(&testMethodStruct{}).
		Method("VMethod").
		Error(); err != nil {
		t.Fatal(err)
	}
	if err := Struct(&testMethodStruct{}).
		Method("PMethod").
		Error(); err != nil {
		t.Fatal(err)
	}
}

type testParamStruct struct{}

func (*testParamStruct) Method(string, int, error) {}

func TestBadParamStruct(t *testing.T) {
	if err := Struct(&testParamStruct{}).
		Param(0, "").
		Error(); err != ErrMethodNotSelected {
		t.Fatalf("%v != %v", err, ErrMethodNotSelected)
	}
	if err := Struct(&testParamStruct{}).
		Method("Method").
		Param(3, "").
		Error(); err != ErrInvalidParamOffset {
		t.Fatalf("%v != %v", err, ErrInvalidParamOffset)
	}
	if err := Struct(&testParamStruct{}).
		Method("Method").
		Param(0, 0).
		Error(); err != ErrParamType {
		t.Fatalf("%v != %v", err, ErrParamType)
	}
}

func TestParam(t *testing.T) {
	if err := Struct(&testParamStruct{}).
		Method("Method").
		Param(0, "").
		Param(1, 0).
		Param(2, ErrorType).
		Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadParams(t *testing.T) {
	if err := Struct(&testParamStruct{}).
		Params("", 0, false).
		Error(); err != ErrMethodNotSelected {
		t.Fatalf("%v != %v", err, ErrMethodNotSelected)
	}
	if err := Struct(&testParamStruct{}).
		Method("Method").
		Params("", 0).
		Error(); err != ErrParamCount {
		t.Fatalf("%v != %v", err, ErrParamCount)
	}
	if err := Struct(&testParamStruct{}).
		Method("Method").
		Params("", 0, 0).
		Error(); err != ErrParamType {
		t.Fatalf("%v != %v", err, ErrParamType)
	}
}

func TestParams(t *testing.T) {
	if err := Struct(&testParamStruct{}).
		Method("Method").
		Params("", 0, ErrorType).
		Error(); err != nil {
		t.Fatal(err)
	}
}

type testReturnStruct struct{}

func (*testReturnStruct) Method() (string, int, error) {
	return "", 0, nil
}

func TestBadReturn(t *testing.T) {
	if err := Struct(&testReturnStruct{}).
		Return(0, "").
		Error(); err != ErrMethodNotSelected {
		t.Fatalf("%v != %v", err, ErrMethodNotSelected)
	}
	if err := Struct(&testReturnStruct{}).
		Method("Method").
		Return(3, "").
		Error(); err != ErrInvalidReturnOffset {
		t.Fatalf("%v != %v", err, ErrInvalidReturnOffset)
	}
	if err := Struct(&testReturnStruct{}).
		Method("Method").
		Return(0, 0).
		Error(); err != ErrReturnType {
		t.Fatalf("%v != %v", err, ErrReturnType)
	}
}

func TestReturn(t *testing.T) {
	if err := Struct(&testReturnStruct{}).
		Method("Method").
		Return(0, "").
		Return(1, 0).
		Return(2, ErrorType).
		Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadReturns(t *testing.T) {
	if err := Struct(&testReturnStruct{}).
		Returns("", 0, false).
		Error(); err != ErrMethodNotSelected {
		t.Fatalf("%v != %v", err, ErrMethodNotSelected)
	}
	if err := Struct(&testReturnStruct{}).
		Method("Method").
		Returns("", 0).Error(); err != ErrReturnCount {
		t.Fatalf("%v != %v", err, ErrReturnCount)
	}
	if err := Struct(&testReturnStruct{}).
		Method("Method").
		Returns("", 0, 0).Error(); err != ErrReturnType {
		t.Fatalf("%v != %v", err, ErrReturnType)
	}
}

func TestReturns(t *testing.T) {
	if err := Struct(&testReturnStruct{}).
		Method("Method").
		Returns("", 0, ErrorType).
		Error(); err != nil {
		t.Fatal(err)
	}
}

type testCallStruct struct{}

func (*testCallStruct) Method(s string, i int, e error) (string, int, error) {
	return s, i, e
}

func TestBadCall(t *testing.T) {
	if _, err := Struct(&testCallStruct{}).
		Method("Method").
		Call(); err != ErrParamCount {
		t.Fatalf("%v != %v", err, ErrParamCount)
	}
}

func TestCall(t *testing.T) {
	r, err := Struct(&testCallStruct{}).
		Method("Method").
		Params("", 0, ErrorType).
		Returns("", 0, ErrorType).
		Call(strTest, intTest, errTest)
	if err != nil {
		t.Fatal(err)
	}
	if strVal := r[0].(string); strVal != strTest {
		t.Fatalf("%v != %v", strVal, strTest)
	}
	if intVal := r[1].(int); intVal != intTest {
		t.Fatalf("%v != %v", intVal, intTest)
	}
	if errVal := r[2].(error); errVal != errTest {
		t.Fatalf("%v != %v", errVal, errTest)
	}
}
