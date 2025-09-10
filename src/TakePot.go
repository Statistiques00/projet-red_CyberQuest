package main

import (
	"fmt"
)

// Fonction pour utiliser une potion
func TakePot(p *Player) {
	potionIndex := -1
	for i, item := range p.Inventory {
		if item == "potion" {
			potionIndex = i
			break
		}
	}

	if potionIndex == -1 {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}

	// Consommer la potion
	p.Inventory = append(p.Inventory[:potionIndex], p.Inventory[potionIndex+1:]...)

	// Soigner le joueur
	heal := 50
	p.HP += heal
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}

	fmt.Printf("Vous avez utilisé une potion ! Points de vie : %d/%d\n", p.HP, p.MaxHP)
}
