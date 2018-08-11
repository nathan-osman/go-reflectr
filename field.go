package reflectr

import (
	"errors"
	"reflect"
)

var (
	errFieldDoesNotExist = errors.New("field does not exist")
	errFieldNotSelected  = errors.New("field not selected")

	errFieldType     = errors.New("field type mismatch")
	errFieldReadOnly = errors.New("field is read-only")
)

// Fields returns a list of field names in the struct.
func (s *StructMeta) Fields() []string {
	t := s.structType
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fields := []string{}
	for i := 0; i < t.NumField(); i++ {
		if len(t.Field(i).PkgPath) == 0 {
			fields = append(fields, t.Field(i).Name)
		}
	}
	return fields
}

// Field ensures that the specified field exists and selects it.
func (s *StructMeta) Field(name string) *StructMeta {
	if s.err != nil {
		return s
	}
	f := s.structValue
	if f.Kind() == reflect.Ptr {
		f = f.Elem()
	}
	f = f.FieldByName(name)
	if f.IsValid() {
		s.field = f
	} else {
		s.err = errFieldDoesNotExist
	}
	return s
}

// Type ensures that the selected field matches the type of the supplied parameter.
func (s *StructMeta) Type(v interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if !s.field.IsValid() {
		s.err = errFieldNotSelected
		return s
	}
	if !typeAwareComparison(s.field.Type(), v) {
		s.err = errFieldType
	}
	return s
}

// Value retrieves the value of the selected field.
func (s *StructMeta) Value() (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	if !s.field.IsValid() {
		return nil, errFieldNotSelected
	}
	return s.field.Interface(), nil
}

// Addr retrieves a pointer to the selected field.
// This method will return an error if the field is read-only.
func (s *StructMeta) Addr() (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	if !s.field.IsValid() {
		return nil, errFieldNotSelected
	}
	if !s.field.CanAddr() {
		return nil, errFieldReadOnly
	}
	return s.field.Addr().Interface(), nil
}
