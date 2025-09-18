package main

import "fmt"

func AddInventory(c Character, item string) {
	if len(c.Inventory) >= 10 {
		fmt.Println("L'inventaire est plein")
		return
	}
	c.Inventory = append(c.Inventory, item)
}