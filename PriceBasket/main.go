package main

import (
	"bufio"
	"fmt"
	"os"
)

var AllItems Items
var UserBasket Basket
var UserOffers Offers
var UserBill Bill

func main() {
	fmt.Println("Hi There!")

	AllItems = AllItems.AddInitialItems()

	fmt.Print("Enter Your Basket Items ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(": ")

	var input string
	if scanner.Scan() {
		input = scanner.Text()
		//fmt.Printf("You wrote \"%s\"\n", input)
	}

	UserBasket = UserBasket.CreateBasketFromInput(input)
	UserOffers.SetInitialOffers()
	UserBill = UserBill.GenerateBillForBasket(UserBasket)
	UserOffers = UserOffers.GetOffersForBasket(UserBasket)
	UserBill.ApplyOffersToTotalPrice(UserOffers)
	// fmt.Println(AllItems)
	// fmt.Println(UserBasket)
	fmt.Println("The applciable offers are as ")
	fmt.Println(UserOffers)

}
