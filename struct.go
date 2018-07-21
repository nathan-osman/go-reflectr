package reflectr

import (
	"errors"
	"reflect"
)

type selType int

const (
	stStruct selType = iota
	stMethod
	stParam
)

var (
	errMustBeStruct       = errors.New("parameter must be a struct or pointer to struct")
	errMethodDoesNotExist = errors.New("method does not exist")
)

// StructMeta provides methods for struct introspection.
type StructMeta struct {
	selType     selType
	structType  reflect.Type
	structValue reflect.Value
	method      reflect.Method
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
		err = errMustBeStruct
	}
	return &StructMeta{
		selType:     stStruct,
		structType:  reflect.TypeOf(v),
		structValue: reflect.ValueOf(v),
		err:         err,
	}
}

// Method ensures that the specified method exists and selects it.
func (s *StructMeta) Method(name string) *StructMeta {
	if s.err != nil {
		return s
	}
	m, ok := s.structType.MethodByName(name)
	if ok {
		s.method = m
	} else {
		s.err = errMethodDoesNotExist
	}
	return s
}

// Error returns an error that occurred (if any).
func (s *StructMeta) Error() error {
	return s.err
}
