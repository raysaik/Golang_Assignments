package main

import (
	"fmt"
	"strings"
)

type Bill struct {
	UserBasket Basket
	TotalPrice float32
	Offers     []string
	OfferPrice float32
}

const Currency string = "GBP"

func (bill Bill) GenerateBillForBasket(basket Basket) Bill {
	var price float32

	for _, itemName := range basket.BasketItems {
		itemDetails := GetItemDetails(itemName)
		unitPrice := ConvertCurrencyDenomination(itemDetails)
		price += unitPrice
	}
	fmt.Printf("SubTotal: £%.2f", price)
	bill.TotalPrice = price
	bill.UserBasket = basket
	return bill
}

func ConvertCurrencyDenomination(item Item) float32 {
	currencyDenomination := item.CurrencyDenominator
	finalPriceInUpperDenomination := item.UnitPrice
	if strings.ToUpper(Currency) == "GBP" {
		if currencyDenomination != "£" {
			finalPriceInUpperDenomination = item.UnitPrice / 100
		}
	}
	return finalPriceInUpperDenomination
}

func (bill Bill) ApplyOffersToTotalPrice(offers Offers) {
	var discountValue float32

	for _, offer := range offers {
		itemDetails := GetItemDetails(offer.ItemName)
		originalPrice := ConvertCurrencyDenomination(itemDetails)
		discount := originalPrice * offer.DiscountPercent / 100
		discountValue += discount
	}
	finalPrice := bill.TotalPrice - discountValue
	fmt.Printf("Total:£%.2f", finalPrice)
}
