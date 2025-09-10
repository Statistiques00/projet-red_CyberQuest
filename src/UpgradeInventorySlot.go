package main

var MaxInventoryCapacity int = 20 // capacité initiale
var InventoryUpgradeUsed int = 0

func UpgradeInventorySlot() bool {
	if InventoryUpgradeUsed < 3 {
		MaxInventoryCapacity += 10
		InventoryUpgradeUsed++
		return true
	}
	return false
}

// Exemple d'ajout dans le marchand
func BuyInventoryUpgrade(playerGold *int) bool {
	if *playerGold >= 30 && InventoryUpgradeUsed < 3 {
		*playerGold -= 30
		return UpgradeInventorySlot()
	}
	return false
}
