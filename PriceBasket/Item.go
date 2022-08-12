package main

import "strings"

type Item struct {
	ItemName            string
	UnitPrice           float32
	Unit                string
	CurrencyDenominator string
}

type Items []Item

var AllItemsAdded Items

func (I Items) AddInitialItems() Items {
	newItem := Item{"Soup", 65, "tin", "p"}
	I = append(I, newItem)
	newItem = Item{"Bread", 80, "loaf", "p"}
	I = append(I, newItem)
	newItem = Item{"Milk", 1.30, "bottle", "£"}
	I = append(I, newItem)
	newItem = Item{"Apples", 1.00, "bag", "£"}
	I = append(I, newItem)
	AllItemsAdded = I
	return I
}

func GetItemDetails(ItemName string) Item {
	var currentItem Item
	for _, item := range AllItemsAdded {
		if strings.ToLower(item.ItemName) == strings.ToLower(ItemName) {
			currentItem = item
			return currentItem
		}
	}
	return currentItem
}
