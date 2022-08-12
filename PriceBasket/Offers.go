package main

import (
	"fmt"
	"sort"
	"strings"
)

type Offer struct {
	OfferId                  string
	ItemName                 string
	DiscountPercent          float32
	BasketConditionItemNames string
}

type Offers []Offer

var AllOffers Offers

func (o Offers) SetInitialOffers() {
	newOffer := Offer{"Offer1", "Apples", 10, ""}
	o = append(o, newOffer)
	newOffer = Offer{"Offer2", "Bread", 50, "Soup,Soup"}
	o = append(o, newOffer)
	AllOffers = o
}

func (o Offers) GetOffersForBasket(basket Basket) Offers {
	var basketOffers Offers
	sort.Strings(basket.BasketItems)
	sortedBasketItems := strings.Join(basket.BasketItems, ",")
	for _, offer := range AllOffers {

		splitArray := strings.Split(string(offer.BasketConditionItemNames), ",")
		sort.Strings(splitArray)
		offer.BasketConditionItemNames = strings.Join(splitArray, ",")
		if len(offer.BasketConditionItemNames) == 0 {
			count := CountOccurenceOfOfferItemInBasket(basket.BasketItems, offer.ItemName)
			for i := 0; i < count; i++ {
				basketOffers = append(basketOffers, offer)
			}

		} else {
			for true {
				if SplitAndCompareItemsMatch(offer.BasketConditionItemNames, sortedBasketItems) && strings.Contains(strings.ToLower(sortedBasketItems), strings.ToLower(offer.ItemName)) {
					basketOffers = append(basketOffers, offer)
					parentString := sortedBasketItems
					index := strings.LastIndex(parentString, offer.BasketConditionItemNames)
					replaceString := offer.BasketConditionItemNames
					if index != (len(parentString) - len(offer.BasketConditionItemNames)) {
						replaceString += ","
					}
					indexTwo := strings.LastIndex(parentString, offer.ItemName)
					replaceStringTwo := offer.ItemName
					if indexTwo != (len(parentString) - len(offer.ItemName)) {
						replaceStringTwo += ","
					}
					parentString = parentString[:index] + strings.Replace(parentString[index:], replaceString, "", 1)
					parentString = parentString[:indexTwo] + strings.Replace(parentString[indexTwo:], replaceStringTwo, "", 1)
					sortedBasketItems = parentString
				} else {
					break
				}
			}
		}
	}
	if len(basketOffers) > 0 {
		for _, offer := range basketOffers {
			percentString := fmt.Sprintf("%v%s", offer.DiscountPercent, "%")
			itemDetails := GetItemDetails(offer.ItemName)
			offValue := itemDetails.UnitPrice * offer.DiscountPercent / 100
			offValueWithCurrency := ConvertCurrencyDenominationForDiscounts(itemDetails.CurrencyDenominator, offValue)
			fmt.Printf("%s %v off:-%v\n", offer.ItemName, percentString, offValueWithCurrency)
		}
	} else {
		fmt.Printf("(no offers available)\n")
	}
	return basketOffers
}

func CountOccurenceOfOfferItemInBasket(items []string, item string) int {
	var count int
	for _, v := range items {
		if strings.ToLower(v) == strings.ToLower(item) {
			count++
		}
	}

	return count
}

func SplitAndCompareItemsMatch(offerItems string, basketItems string) bool {
	if len(offerItems) > len(basketItems) {
		return false
	}
	if !strings.Contains(strings.ToLower(basketItems), strings.ToLower(offerItems)) {
		return false
	}
	return true
}
