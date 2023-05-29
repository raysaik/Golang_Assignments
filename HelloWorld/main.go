package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("Hi There!")
	fmt.Println(math.Floor(2.141))
	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(2.5))
	fmt.Println(reflect.TypeOf(2.5))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))
	fmt.Println(reflect.TypeOf('A'))

	fmt.Printf("The sqrt number is %g for now. ", math.Sqrt(7))
}
