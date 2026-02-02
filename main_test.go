package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestPerson struct {
	Name string
	Age  int
}

func TestReadReflect(t *testing.T) {
	var a interface{} = 1

	var p = TestPerson{
		Name: "John",
		Age:  20,
	}

	// aType := reflect.TypeOf(a)

	fmt.Println(reflect.ValueOf(a))

	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.ValueOf(p))

	fmt.Println(reflect.TypeOf(a).Kind())
}

func TestWriteReflect(t *testing.T) {
	// var a interface{} = map[string]string{
	// 	"name": "John",
	// 	"age":  "20",
	// }

	// aValue := reflect.ValueOf(&a).Elem()
	// aValue.SetString("name")
}
