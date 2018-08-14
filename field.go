package reflectr

import (
	"errors"
	"reflect"
)

var (
	// ErrFieldDoesNotExist indicates that the selected field does not exist.
	ErrFieldDoesNotExist = errors.New("field does not exist")

	// ErrFieldNotSelected indicates that no field was selected.
	ErrFieldNotSelected = errors.New("field not selected")

	// ErrFieldType indicates that the field type does not match.
	ErrFieldType = errors.New("field type mismatch")

	// ErrFieldReadOnly indicates that the field is read-only.
	ErrFieldReadOnly = errors.New("field is read-only")
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
		s.err = ErrFieldDoesNotExist
	}
	return s
}

// Type ensures that the selected field matches the type of the supplied parameter.
func (s *StructMeta) Type(v interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if !s.field.IsValid() {
		s.err = ErrFieldNotSelected
		return s
	}
	if !typeAwareComparison(s.field.Type(), v) {
		s.err = ErrFieldType
	}
	return s
}

// Value retrieves the value of the selected field.
func (s *StructMeta) Value() (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	if !s.field.IsValid() {
		return nil, ErrFieldNotSelected
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
		return nil, ErrFieldNotSelected
	}
	if !s.field.CanAddr() {
		return nil, ErrFieldReadOnly
	}
	return s.field.Addr().Interface(), nil
}
