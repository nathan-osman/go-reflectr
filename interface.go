package reflectr

import (
	"reflect"
)

// Interface accepts a nil pointer to an interface and returns a value that can be passed as a type to the introspection methods.
func Interface(v interface{}) interface{} {
	return reflect.TypeOf(v).Elem()
}

// ErrorType may be specified wherever a type variable is expected.
var ErrorType = Interface((*error)(nil))
