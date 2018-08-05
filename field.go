package reflectr

import (
	"errors"
	"reflect"
)

var (
	errFieldDoesNotExist = errors.New("field does not exist")
	errFieldNotSelected  = errors.New("field not selected")

	errFieldType       = errors.New("field type mismatch")
	errFieldUnexported = errors.New("field is unexported")
)

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

// SetValue sets the value of the currently selected field.
func (s *StructMeta) SetValue(v interface{}) *StructMeta {
	s.Type(v)
	if s.err != nil {
		return s
	}
	if !s.field.CanSet() {
		s.err = errFieldUnexported
		return s
	}
	s.field.Set(reflect.ValueOf(v))
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
