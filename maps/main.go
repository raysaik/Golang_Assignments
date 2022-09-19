package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 map[int]string = make(map[int]string)

func main() {
	m = make(map[string]Vertex)
	m["Burj-Al_Khalifa"] = Vertex{
		Lat:  55,
		Long: 25,
	}
	fmt.Println(m["Bell Labs"])

	//print a whole map
	fmt.Println(m)

	m2[0] = "Rohit Sharma"
	m2[1] = "KL Rahul"
	m2[2] = "Virat Kohli"

	fmt.Println(m2)

	//Retrieve an element:
	elem := m2[1]
	fmt.Println(elem)

	delete(m2, 1)

	//Test that a key is present with a two-value assignment:
	testelem, ok := m2[2]
	fmt.Printf("%s : %t", testelem, ok)
	fmt.Println()

	// var then variable name then variable type
	first := GetUserInput()
	// Taking input from user
	//mapCount := make(map[string]int)
	mapCount := getWordCounts(first)
	fmt.Println(mapCount)
}

func GetUserInput() string {
	fmt.Print("Enter Your Sentence ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(": ")

	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
