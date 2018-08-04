package reflectr

import (
	"reflect"
)

var typeType = reflect.TypeOf((*reflect.Type)(nil)).Elem()

func typeAwareComparison(t reflect.Type, v interface{}) bool {
	vType := reflect.TypeOf(v)
	if vType.Kind() == reflect.Ptr && vType.Implements(typeType) {
		vType = v.(reflect.Type)
	}
	if t.Kind() == reflect.Interface {
		return vType.Implements(t)
	}
	return t == vType
}
