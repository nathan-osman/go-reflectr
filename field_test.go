package reflectr

import (
	"reflect"
	"testing"
)

type testFieldStruct struct {
	field0 string
	Field1 string
	Field2 int
	Field3 error
}

func TestFields(t *testing.T) {
	var (
		sFields = Struct(&testFieldStruct{}).Fields()
		cFields = []string{"Field1", "Field2", "Field3"}
	)
	if !reflect.DeepEqual(sFields, cFields) {
		t.Fatalf("%v != %v", sFields, cFields)
	}
}

func TestBadField(t *testing.T) {
	if err := Struct(&testFieldStruct{}).
		Field("42").
		Error(); err != errFieldDoesNotExist {
		t.Fatalf("%v != %v", err, errFieldDoesNotExist)
	}
}

func TestField(t *testing.T) {
	if err := Struct(&testFieldStruct{}).
		Field("Field1").
		Field("Field2").
		Field("Field3").
		Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadType(t *testing.T) {
	if err := Struct(&testFieldStruct{}).
		Type(0).
		Error(); err != errFieldNotSelected {
		t.Fatalf("%v != %v", err, errFieldNotSelected)
	}
	if err := Struct(&testFieldStruct{}).
		Field("Field1").
		Type(0).
		Error(); err != errFieldType {
		t.Fatalf("%v != %v", err, errFieldType)
	}
}

func TestType(t *testing.T) {
	if err := Struct(&testFieldStruct{}).
		Field("Field1").
		Type("").
		Field("Field2").
		Type(0).
		Field("Field3").
		Type(ErrorType).
		Error(); err != nil {
		t.Fatal(err)
	}
}

func TestBadSetValue(t *testing.T) {
	if err := Struct(&testFieldStruct{}).
		Field("field0").
		SetValue("").
		Error(); err != errFieldReadOnly {
		t.Fatalf("%v != %v", err, errFieldReadOnly)
	}
}

func TestSetValue(t *testing.T) {
	s := &testFieldStruct{}
	if err := Struct(s).
		Field("Field1").
		SetValue(strTest).
		Field("Field2").
		SetValue(intTest).
		Field("Field3").
		SetValue(errTest).
		Error(); err != nil {
		t.Fatal(err)
	}
	if s.Field1 != strTest {
		t.Fatalf("%v != %v", s.Field1, strTest)
	}
	if s.Field2 != intTest {
		t.Fatalf("%v != %v", s.Field2, intTest)
	}
	if s.Field3 != errTest {
		t.Fatalf("%v != %v", s.Field3, errTest)
	}
}

func TestBadValue(t *testing.T) {
	if _, err := Struct(&testFieldStruct{}).
		Value(); err != errFieldNotSelected {
		t.Fatalf("%v != %v", err, errFieldNotSelected)
	}
}

func TestValue(t *testing.T) {
	s := Struct(&testFieldStruct{
		Field1: strTest,
		Field2: intTest,
		Field3: errTest,
	})
	if v, err := s.Field("Field1").Value(); err != nil {
		t.Fatal(err)
	} else if v.(string) != strTest {
		t.Fatalf("%v != %v", v.(string), strTest)
	}
	if v, err := s.Field("Field2").Value(); err != nil {
		t.Fatal(err)
	} else if v.(int) != intTest {
		t.Fatalf("%v != %v", v.(int), intTest)
	}
	if v, err := s.Field("Field3").Value(); err != nil {
		t.Fatal(err)
	} else if v.(error) != errTest {
		t.Fatalf("%v != %v", v.(error), errTest)
	}
}
