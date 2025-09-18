package main

// Fonction pour retirer un objet de l'inventaire
func RemoveInventory(c *Character, item int) {
	for i, v := range c.Inventory {
		if v == c.Inventory[item-1] {
			// On enlève l'élément i du slice
			c.Inventory = append(c.Inventory[:i-1], c.Inventory[i-1:]...)
			return
		}
	}

}
