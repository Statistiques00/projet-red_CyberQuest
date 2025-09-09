package main

import (
	"fmt"
)

// Supposons que l'inventaire est un slice de string dans le personnage
func takePot() {
	potionIndex := -1
	for i, item := range player.Equipment.Loot {
		if item == "Potion" {
			potionIndex = i
			break
		}
	}
	if potionIndex == -1 {
		fmt.Println("Aucune potion dans l'inventaire.")
		return
	}

	// Supprimer la potion
	player.Equipment.Loot = append(player.Equipment.Loot[:potionIndex], player.Equipment.Loot[potionIndex+1:]...)

	// Soigner le joueur
	player.HP += 50
	if player.HP > player.MaxHP {
		player.HP = player.MaxHP
	}

	fmt.Printf("Vous avez utilisé une potion ! PV : %d/%d\n", player.HP, player.MaxHP)
}
