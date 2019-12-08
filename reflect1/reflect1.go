package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=== Go Reflection ===")

	item := struct {
		Name    string
		Age     int
		Job     string
		Address struct {
			Town    string
			Zipcode string
		}
	}{Name: "John Doe", Age: 30, Job: "Teacher", Address: struct{ Town, Zipcode string }{Town: "Paris", Zipcode: "75001"}}

	printValue(reflect.ValueOf(&item))
}

/*func printDetails(item interface{}) {
	val := reflect.ValueOf(item)
}*/

func printValue(val reflect.Value) {
	t := val.Type()

	switch t.Kind() {
	case reflect.Ptr:
		fmt.Println("This is a pointer")
		printValue(val.Elem())
	case reflect.Struct:
		fmt.Println("This is a struct")
		for i := 0; i < val.NumField(); i++ {
			printValue(val.Field(i))
		}
	default:
		fmt.Printf("Field %s equals %v\n", t.Name(), val.Interface())
	}

}
