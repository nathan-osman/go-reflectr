package reflectr

import (
	"errors"
	"reflect"
)

var (
	errMethodDoesNotExist = errors.New("method does not exist")
	errMethodNotSelected  = errors.New("method not selected")

	errInvalidParamOffset = errors.New("invalid parameter offset")
	errParamType          = errors.New("parameter type mismatch")
	errParamCount         = errors.New("parameter count mismatch")

	errInvalidReturnOffset = errors.New("invalid return offset")
	errReturnType          = errors.New("return type mismatch")
	errReturnCount         = errors.New("return count mismatch")
)

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

// Param ensures the parameter (specified by offset) matches the provided type.
func (s *StructMeta) Param(i int, v interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.method.Type == nil {
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
	if s.method.Type == nil {
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
	if s.method.Type == nil {
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
	if s.method.Type == nil {
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

// Call invokes the selected method, checking to ensure parameter types match the method.
func (s *StructMeta) Call(v ...interface{}) ([]interface{}, error) {
	s.Params(v...)
	if s.err != nil {
		return nil, s.err
	}
	params := make([]reflect.Value, len(v))
	for i, p := range v {
		params[i] = reflect.ValueOf(p)
	}
	var (
		rVals = s.structValue.MethodByName(s.method.Name).Call(params)
		ret   = make([]interface{}, len(rVals))
	)
	for i, r := range rVals {
		ret[i] = r.Interface()
	}
	return ret, nil
}
