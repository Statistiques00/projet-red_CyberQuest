package main


// Fonction pour retirer un objet de l'inventaire
func RemoveInventory(c *Character, item string) {
	for i, v := range c.Inventory {
		if v == item {
			// On enlève l'élément i du slice
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return 
		}
	}

}





