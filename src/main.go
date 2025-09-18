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
		| 5 - Combats                   |
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
			trainingMonster := InitGoblin()
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
