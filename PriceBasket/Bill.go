package main

import (
	"fmt"
)

type Bill struct {
	UserBasket Basket
	TotalPrice float32
	Offers     []string
	OfferPrice float32
}

const CurrencyHigher string = "£"
const CurrencyLower string = "p"

var AllBills Bill

func (bill Bill) GenerateBillForBasket(basket Basket) Bill {
	var price float32

	for _, itemName := range basket.BasketItems {
		itemDetails := GetItemDetails(itemName)
		unitPrice := ConvertCurrencyDenomination(itemDetails)
		price += unitPrice
	}
	priceWithCurrency := fmt.Sprintf("%s%.2f", CurrencyHigher, price)
	fmt.Printf("Subtotal: %s\n", priceWithCurrency)
	bill.TotalPrice = price
	bill.UserBasket = basket
	AllBills = bill
	return bill
}

func ConvertCurrencyDenomination(item Item) float32 {
	currencyDenomination := item.CurrencyDenominator
	finalPriceInUpperDenomination := item.UnitPrice
	if CurrencyHigher == "£" {
		if currencyDenomination != CurrencyHigher {
			finalPriceInUpperDenomination = item.UnitPrice / 100
		}
	}
	return finalPriceInUpperDenomination
}

func ConvertCurrencyDenominationForDiscounts(currency string, discountValue float32) string {

	if currency == CurrencyHigher {
		discountValue *= 100
	}
	if discountValue >= 100 {
		return fmt.Sprintf("%v%s", discountValue/100, CurrencyHigher)
	} else {
		return fmt.Sprintf("%v%s", discountValue, CurrencyLower)
	}
}

func (bill Bill) ApplyOffersToTotalPrice(offers Offers) {
	var discountValue float32

	for _, offer := range offers {
		itemDetails := GetItemDetails(offer.ItemName)
		originalPrice := ConvertCurrencyDenomination(itemDetails)
		discount := originalPrice * offer.DiscountPercent / 100
		discountValue += discount
	}
	finalPrice := AllBills.TotalPrice - discountValue
	finalPriceWithCurrency := fmt.Sprintf("%s%.2f", CurrencyHigher, finalPrice)
	fmt.Printf("Total:%s", finalPriceWithCurrency)
}
