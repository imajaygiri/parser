package helpers

import (
	"fmt"
	"reflect"
)

// [ExpectType](unknownType) -> return cast_to_T(r)
func ExpectType[T any](r any) T {
	expectedType := reflect.TypeFor[T]()
	recievedType := reflect.TypeOf(r)

	if expectedType == recievedType {
		return r.(T)
	}

	panic(fmt.Sprintf("Expected --> [%s] but recieved --> [%s] instead\n", expectedType, recievedType))
}
