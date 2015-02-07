// go-search is a simple package to perform a search in a slice, accepting
// multiple types (even structs)
//
// Author: Magd Kudama
package search

import (
	"errors"
	"reflect"
)

var (
	// ErrNotSlice returned when the input value is not a slice
	ErrNotSlice = errors.New("Input must be a slice")

	// ErrInvalidType returned when the input is of type interface or pointer
	ErrInvalidType = errors.New("Input can't be a pointer / interface{}")

	// ErrNotSameType returned when the input value and the element to compare don't have the same type
	ErrNotSameType = errors.New("Input and element to compare are not the same type")
)

// T is just an alias to the interface{} type
type T interface{}

// Search is the only function in this package. It's just a simple function
// that receives a slice and an element to search for as an input
// query result evaluating method.
//
// It returns three parameters: a boolean indicating if the value was found,
// an int with the position found (-1 otherwise) and an error value, in case
// an error occurred while doing the search
//
// Example:
//     values := []string{"hello", "go"}
//     Search(values, "go")
func Search(values T, search T) (bool, int, error) {
	typeOf := reflect.TypeOf(values)

	if typeOf.Kind() != reflect.Slice {
		return false, -1, ErrNotSlice
	}

	kind := typeOf.Elem().Kind()

	// I should probably accept a Ptr, but meh!
	if kind == reflect.Ptr || kind == reflect.Interface {
		return false, -1, ErrInvalidType
	}

	if kind != reflect.TypeOf(search).Kind() {
		return false, -1, ErrNotSameType
	}

	inputValues := reflect.ValueOf(values)

	for i := 0; i < inputValues.Len(); i++ {
		result := inputValues.Index(i).Interface()

		if result == search {
			return true, i, nil
		}
	}

	return false, -1, nil
}
