package reflectx

import "reflect"

func IsPointer(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Ptr
}
