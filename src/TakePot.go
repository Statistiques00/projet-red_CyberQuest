package main

import (
	"fmt"
)

// Fonction pour utiliser une potion
func TakePot(p *Character) {
	potionIndex := -1

	// Chercher une potion dans l'inventaire
	for i, item := range p.Inventory {
		if item == "antivirus" {
			potionIndex = i
			break // Arrêter dès qu'on trouve une potion
		}
	}

	// Si aucune potion n'est trouvée
	if potionIndex == -1 {
		fmt.Println("Vous n'avez pas d'antivirus dans votre inventaire.")
		return
	}

	// Utiliser la potion
	heal := 50
	oldHP := p.HP
	p.HP += heal
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}

	// Consommer la potion (la retirer de l'inventaire)
	p.Inventory = append(p.Inventory[:potionIndex], p.Inventory[potionIndex+1:]...)

	fmt.Printf("Vous avez utilisé un antivirus ! Points de vie : %d/%d (soigné de %d PV)\n",
		p.HP, p.MaxHP, p.HP-oldHP)
}
