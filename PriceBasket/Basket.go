package main

import (
	"fmt"
	"os"
	"strings"
)

type Basket struct {
	BasketName  string
	BasketItems []string
}

func (userBasket Basket) CreateBasketFromInput(input string) Basket {
	if input == "" {
		fmt.Println("user input not valid")
		os.Exit(1)
	}
	strArray := strings.Split(string(input), " ")
	userBasket.BasketName = strArray[0]
	for i := 1; i < len(strArray); i++ {
		userBasket.BasketItems = append(userBasket.BasketItems, strArray[i])
	}

	return userBasket
}
