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
func BuyInventoryUpgrade(player Character) bool {
	if player.BTC >= 30 && InventoryUpgradeUsed < 3 {
		player.BTC -= 30
		return UpgradeInventorySlot()
	}
	return false
}
