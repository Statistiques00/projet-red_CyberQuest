package main

import (
	"fmt"
)

func PoisonPot(target *Monster) {
	for i, c := range player.Inventory {
		if c == "malware" {
			fmt.Println("+=========================================+")
			fmt.Printf("| %s est empoisonné pendant 3 tours !      |\n", target.Name)
			fmt.Println("+=========================================+")
			target.Poisoned = 3
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire.")
}

// Ajout de la potion de poison dans le marchand
func BuyPoisonPotion(inventory *[]string, playerGold *int) bool {
	if *playerGold >= 20 { // prix exemple : 20 pièces d'or
		*playerGold -= 20
		*inventory = append(*inventory, "Potion de poison")
		return true
	}
	return false
}
