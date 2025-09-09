package main

import "fmt"

func (c *Character) AddInventory(item string) {
	if len(c.Inventory) >= 10 {
		fmt.Println("L'inventaire est plein")
		return
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("%s a été ajouté à l'inventaire\n", item)
}