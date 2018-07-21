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

	errInvalidParamOffset = errors.New("invalid parameter offset")
	errParamType          = errors.New("parameter type mismatch")
	errParamCount         = errors.New("parameter count mismatch")

	errInvalidReturnOffset = errors.New("invalid return offset")
	errReturnType          = errors.New("return type mismatch")
	errReturnCount         = errors.New("return count mismatch")
)

// Interface accepts a nil pointer to an interface and returns a value that can be passed as a type to the introspection methods.
func Interface(v interface{}) interface{} {
	return reflect.TypeOf(v).Elem()
}

var (
	typeType = reflect.TypeOf((*reflect.Type)(nil)).Elem()

	// ErrorType may be specified wherever a type variable is expected.
	ErrorType = Interface((*error)(nil))
)

func typeAwareComparison(t reflect.Type, v interface{}) bool {
	vType := reflect.TypeOf(v)
	if vType.Kind() == reflect.Ptr && vType.Implements(typeType) {
		return t.Implements(v.(reflect.Type))
	}
	return t == vType
}

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

// Param ensures the parameter (specified by offset) matches the provided type.
func (s *StructMeta) Param(i int, v interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.selType != stMethod {
		s.err = errMethodNotSelected
		return s
	}
	methodType := s.method.Type
	if i+1 >= methodType.NumIn() {
		s.err = errInvalidParamOffset
		return s
	}
	if !typeAwareComparison(methodType.In(i+1), v) {
		s.err = errParamType
	}
	return s
}

// Params ensures the parameters match the provided types.
func (s *StructMeta) Params(v ...interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.selType != stMethod {
		s.err = errMethodNotSelected
		return s
	}
	var (
		methodType = s.method.Type
		paramCount = methodType.NumIn() - 1
	)
	if len(v) != paramCount {
		s.err = errParamCount
		return s
	}
	for i := 0; i < paramCount; i++ {
		if !typeAwareComparison(methodType.In(i+1), v[i]) {
			s.err = errParamType
			break
		}
	}
	return s
}

// Return ensures the parameter (specified by offset) matches the provided type.
func (s *StructMeta) Return(i int, v interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.selType != stMethod {
		s.err = errMethodNotSelected
		return s
	}
	methodType := s.method.Type
	if i >= methodType.NumOut() {
		s.err = errInvalidReturnOffset
		return s
	}
	if !typeAwareComparison(methodType.Out(i), v) {
		s.err = errReturnType
	}
	return s
}

// Returns ensures that the return types match the provided types.
func (s *StructMeta) Returns(v ...interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.selType != stMethod {
		s.err = errMethodNotSelected
		return s
	}
	var (
		methodType  = s.method.Type
		returnCount = methodType.NumOut()
	)
	if len(v) != returnCount {
		s.err = errReturnCount
		return s
	}
	for i := 0; i < returnCount; i++ {
		if !typeAwareComparison(methodType.Out(i), v[i]) {
			s.err = errReturnType
			break
		}
	}
	return s
}

// Error returns an error that occurred (if any).
func (s *StructMeta) Error() error {
	return s.err
}
