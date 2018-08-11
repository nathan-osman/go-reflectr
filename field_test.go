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

func TestBadAddr(t *testing.T) {
	if _, err := Struct(&testFieldStruct{}).
		Addr(); err != errFieldNotSelected {
		t.Fatalf("%v != %v", err, errFieldNotSelected)
	}
	if _, err := Struct(testFieldStruct{}).
		Field("Field1").
		Addr(); err != errFieldReadOnly {
		t.Fatalf("%v != %v", err, errFieldReadOnly)
	}
}

func TestAddr(t *testing.T) {
	var (
		f = &testFieldStruct{}
		s = Struct(f)
	)
	if v, err := s.Field("Field1").Addr(); err != nil {
		t.Fatal(err)
	} else if *v.(*string) = strTest; f.Field1 != strTest {
		t.Fatalf("%v != %v", f.Field1, strTest)
	}
	if v, err := s.Field("Field2").Addr(); err != nil {
		t.Fatal(err)
	} else if *v.(*int) = intTest; f.Field2 != intTest {
		t.Fatalf("%v != %v", f.Field2, intTest)
	}
	if v, err := s.Field("Field3").Addr(); err != nil {
		t.Fatal(err)
	} else if *v.(*error) = errTest; f.Field3 != errTest {
		t.Fatalf("%v != %v", f.Field3, errTest)
	}
}
