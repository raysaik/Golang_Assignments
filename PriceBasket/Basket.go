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
		fmt.Println("User input not valid")
		os.Exit(1)
	}
	strArray := strings.Split(strings.ToLower(input), " ")
	basket.BasketName = strArray[0]
	for i := 1; i < len(strArray); i++ {
		if (GetItemDetails(strArray[i]) == Item{}) {
			fmt.Println("User input contains invalid items. Please try again")
			os.Exit(1)
		}
		//needs thread safety
		basket.BasketItems = append(basket.BasketItems, strArray[i])
	}
	UserBasket = basket
	return basket
}
