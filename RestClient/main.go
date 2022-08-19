package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	REGION      = "eu-west-1"
	BUCKET_NAME = "lambdatestyoutube"
)

func main() {
	response, err := http.Get("https://catfact.ninja/fact")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseData))
}
