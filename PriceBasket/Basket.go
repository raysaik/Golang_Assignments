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

var UserBasket Basket

func (basket Basket) CreateBasketFromInput(input string) Basket {
	if input == "" {
		fmt.Println("user input not valid")
		os.Exit(1)
	}
	strArray := strings.Split(string(input), " ")
	basket.BasketName = strArray[0]
	for i := 1; i < len(strArray); i++ {
		basket.BasketItems = append(basket.BasketItems, strArray[i])
	}
	UserBasket = basket
	return basket
}
