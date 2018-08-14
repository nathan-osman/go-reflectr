package reflectr

import (
	"errors"
	"reflect"
)

var (
	// ErrMethodDoesNotExist indicates that the selected method does not exist.
	ErrMethodDoesNotExist = errors.New("method does not exist")

	// ErrMethodNotSelected indicates that no method was selected.
	ErrMethodNotSelected = errors.New("method not selected")

	// ErrInvalidParamOffset indicates that the supplied parameter offset is invalid.
	ErrInvalidParamOffset = errors.New("invalid parameter offset")

	// ErrParamType indicates that the parameter type does not match.
	ErrParamType = errors.New("parameter type mismatch")

	// ErrParamCount indicates that an incorrect number of parameter types were supplied.
	ErrParamCount = errors.New("parameter count mismatch")

	// ErrInvalidReturnOffset indicates that the supplied return offset is invalid.
	ErrInvalidReturnOffset = errors.New("invalid return offset")

	// ErrReturnType indicates that the return type does not match.
	ErrReturnType = errors.New("return type mismatch")

	// ErrReturnCount indicates that an incorrect number of return types were supplied.
	ErrReturnCount = errors.New("return count mismatch")
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
		s.err = ErrMethodDoesNotExist
	}
	return s
}

// Param ensures the parameter (specified by offset) matches the provided type.
func (s *StructMeta) Param(i int, v interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.method.Type == nil {
		s.err = ErrMethodNotSelected
		return s
	}
	methodType := s.method.Type
	if i+1 >= methodType.NumIn() {
		s.err = ErrInvalidParamOffset
		return s
	}
	if !typeAwareComparison(methodType.In(i+1), v) {
		s.err = ErrParamType
	}
	return s
}

// Params ensures the parameters match the provided types.
func (s *StructMeta) Params(v ...interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.method.Type == nil {
		s.err = ErrMethodNotSelected
		return s
	}
	var (
		methodType = s.method.Type
		paramCount = methodType.NumIn() - 1
	)
	if len(v) != paramCount {
		s.err = ErrParamCount
		return s
	}
	for i := 0; i < paramCount; i++ {
		if !typeAwareComparison(methodType.In(i+1), v[i]) {
			s.err = ErrParamType
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
		s.err = ErrMethodNotSelected
		return s
	}
	methodType := s.method.Type
	if i >= methodType.NumOut() {
		s.err = ErrInvalidReturnOffset
		return s
	}
	if !typeAwareComparison(methodType.Out(i), v) {
		s.err = ErrReturnType
	}
	return s
}

// Returns ensures that the return types match the provided types.
func (s *StructMeta) Returns(v ...interface{}) *StructMeta {
	if s.err != nil {
		return s
	}
	if s.method.Type == nil {
		s.err = ErrMethodNotSelected
		return s
	}
	var (
		methodType  = s.method.Type
		returnCount = methodType.NumOut()
	)
	if len(v) != returnCount {
		s.err = ErrReturnCount
		return s
	}
	for i := 0; i < returnCount; i++ {
		if !typeAwareComparison(methodType.Out(i), v[i]) {
			s.err = ErrReturnType
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
