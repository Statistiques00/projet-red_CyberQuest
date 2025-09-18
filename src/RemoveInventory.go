package main

// Fonction pour retirer un objet de l'inventaire
func RemoveInventory(c *Character, item int) {
	// Check bounds
	if item < 1 || item > len(c.Inventory) {
		return // Invalid item index
	}

	// Convert to 0-based index
	index := item - 1

	// Remove the item at the specified index
	if len(c.Inventory) == 1 {
		c.Inventory = []string{}
	} else {
		// Remove element at index using slice operations
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}
