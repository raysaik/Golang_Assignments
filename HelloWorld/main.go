package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"time"
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

	fmt.Println("Time exercise")
	var now time.Time = time.Now()
	fmt.Println(now.Year())

	fmt.Println("Get user input exercise")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Println(input)

}
