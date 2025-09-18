package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	CharacterCreation()
	for {
		ClearScreen()
		fmt.Print(`
		+===============================+
		|    ===  MENU PRINCIPAL ===    |
		+===============================+
		| 1 - Infos personnage          |
		| 2 - Inventaire                |
		| 3 - Marchand                  |
		| 4 - Forgeron                  |
		| 5 - Entrainement              |
		| 6 - Quitter                   |
		+===============================+
		Choix : `)
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			DisplayInfo(&player)
		case "2":
			AccessInventory(&player)
		case "3":
			fmt.Println("Marchand")
			AccessMarchand(&player)
		case "4":
			fmt.Println("Forgeron")
			ForgeronMenu(&player)
		case "5":
			trainingMonster := Monster{
				Name:       "Ennemi d'entraînement",
				MaxHP:      30,
				HP:         30,
				Attack:     5,
				Defense:    2,
				Speed:      3,
				Experience: 10,
				LootTable: []LootItem{
					{Name: "micro-batterie", DropChance: 1.0},
					{Name: "plaque métallique", DropChance: 1.0},
					{Name: "isolant plastique", DropChance: 1.0},
					{Name: "alimentation haute capacité", DropChance: 1.0},
					{Name: "capteurs de mouvement", DropChance: 1.0},
					{Name: "puce Bluetooth", DropChance: 1.0},
					{Name: "monture légère", DropChance: 1.0},
					{Name: "verre spécial", DropChance: 1.0},
					{Name: "capteur de mouvement", DropChance: 1.0},
				},
				BTC: 10,
			}
			TrainingFight(&player, trainingMonster, 1)
		case "6":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
		scanner.Scan()
	}
}
