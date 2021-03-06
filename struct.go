package reflectr

import (
	"errors"
	"reflect"
)

var (
	// ErrMustBeStruct indicates that an invalid type was passed to Struct().
	ErrMustBeStruct = errors.New("parameter must be a struct or pointer to struct")
)

// StructMeta provides methods for struct introspection.
type StructMeta struct {
	structType  reflect.Type
	structValue reflect.Value
	method      reflect.Method
	field       reflect.Value
	err         error
}

// Struct creates a StructMeta from the provided struct.
func Struct(v interface{}) *StructMeta {
	var (
		structType = reflect.TypeOf(v)
		err        error
	)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	if structType.Kind() != reflect.Struct {
		err = ErrMustBeStruct
	}
	return &StructMeta{
		structType:  reflect.TypeOf(v),
		structValue: reflect.ValueOf(v),
		err:         err,
	}
}

// IsPtr determines whether a struct or a pointer to struct was provided to the Struct function.
func (s *StructMeta) IsPtr() bool {
	return s.structType.Kind() == reflect.Ptr
}

// Error returns an error that occurred (if any).
func (s *StructMeta) Error() error {
	return s.err
}
