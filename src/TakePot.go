package main

import (
	"fmt"
)

// Fonction pour utiliser une potion
func TakePot(p *Character) {
	potionIndex := -1
	for i, item := range p.Inventory {
		if item == "antivirus" {
			potionIndex = i
			heal := 50
			p.HP += heal
			if p.HP > p.MaxHP {
				p.HP = p.MaxHP
			}

		}
	}

	if potionIndex == -1 {
		fmt.Println("Vous n'avez pas d'antivirus dans votre inventaire.")
		return 		
	}

	// Consommer la potion
	p.Inventory = append(p.Inventory[:potionIndex], p.Inventory[potionIndex+1:]...)

	fmt.Printf("Vous avez utilisé une potion ! Points de vie : %d/%d\n", p.HP, p.MaxHP)


	// Soigner le joueur
	p.HP += 50
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}

	fmt.Printf("Vous avez utilisé une potion ! PV : %d/%d\n", p.HP, p.MaxHP)
}

