package main

import (
	"bufio"
	"fmt"
	"os"
)

var PriceBasketItems Items
var PriceBasket Basket
var PriceBasketOffers Offers
var PriceBasketBill Bill

func main() {
	fmt.Println("Hi There! Welcome to the PriceBasket App")
	//Set initial values
	PriceBasketItems.AddInitialItems()
	PriceBasketOffers.SetInitialOffers()

	//Create Data from User input
	input := GetUserInput()
	basket := PriceBasket.CreateBasketFromInput(input)

	//Generate Bill
	PriceBasketBill.GenerateBillForBasket(basket)
	offers := PriceBasketOffers.GetOffersForBasket(basket)
	PriceBasketBill.ApplyOffersToTotalPrice(offers)

}

func GetUserInput() string {
	fmt.Print("Enter Your Basket Items ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(": ")

	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
