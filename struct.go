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
	errMustBeStruct = errors.New("parameter must be a struct or pointer to struct")

	errMethodDoesNotExist = errors.New("method does not exist")
	errMethodNotSelected  = errors.New("method not selected")

	errReturnsCount = errors.New("return value count mismatch")
	errReturnsType  = errors.New("return value type mismatch")
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
		s.selType = stMethod
		s.method = m
	} else {
		s.err = errMethodDoesNotExist
	}
	return s
}

// Returns ensures that the return types of a method matche the provided types.
func (s *StructMeta) Returns(v ...interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.selType != stMethod {
		s.err = errMethodNotSelected
		return s
	}
	var (
		methodType = s.method.Type
		paramCount = methodType.NumOut()
	)
	if len(v) != paramCount {
		s.err = errReturnsCount
		return s
	}
	for i := 0; i < paramCount; i++ {
		if reflect.TypeOf(v[i]) != methodType.Out(i) {
			s.err = errReturnsType
			return s
		}
	}
	return s
}

// Error returns an error that occurred (if any).
func (s *StructMeta) Error() error {
	return s.err
}
