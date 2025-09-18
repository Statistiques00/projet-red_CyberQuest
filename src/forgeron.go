package main

import (
	"bufio"
	"fmt"
	"os"
)

var MaxInventoryCapacity1 int = 20 // à adapter selon votre gestion

func ForgeronMenu(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearScreen()
		fmt.Println("+=========================================+")
		fmt.Println("|            === FORGERON ===             |")
		fmt.Println("+=========================================+")
		fmt.Printf("| BTC : %-34d|\n", player.BTC)
		fmt.Println("+=========================================+")
		fmt.Println("| 1 - Forger Une Armure                   |")
		fmt.Println("| 2 - Améliorer le sac                    |")
		fmt.Println("| 3 - Quitter                             |")
		fmt.Println("+=========================================+")
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			ClearScreen()
			fmt.Println("+=========================================+")
			fmt.Println("|              === FORGERON 😊 ===        |")
			fmt.Println("+=========================================+")
			fmt.Println("| 1 - Lunettes anti-lumière bleue         |")
			fmt.Println("| 2 - Blindage électromagnétique          |")
			fmt.Println("| 3 - Chaussures connectées IoT           |")
			fmt.Println("| 4 - Quitter                             |")
			fmt.Println("+=========================================+")
			fmt.Print("Choix : ")
			var choix2 int
			fmt.Scan(&choix2)

			if choix2 == 4 {
				ClearScreen()
				return
			}
			var item string
			var required []string
			switch choix2 {
			case 1:
				item = "Lunettes anti-lumière bleue"
				required = []string{"verre spécial", "isolant plastique", "monture légère"}
				if HasMaterials(player, required) {
					player.Inventory = append(player.Inventory, item)
					RemoveMaterials(&player.Inventory, required)
					player.BTC -= 5
					fmt.Println("Vous avez forgé :", item)
				}
			case 2:
				item = "Blindage électromagnétique"
				required = []string{"plaque métallique", "isolant plastique", "alimentation haute capacité"}
				if HasMaterials(player, required) {
					player.Inventory = append(player.Inventory, item)
					RemoveMaterials(&player.Inventory, required)
					player.BTC -= 5
					fmt.Println("Vous avez forgé :", item)
				}
			case 3:
				item = "Chaussures connectées IoT"
				required = []string{"capteurs de mouvement", "micro-batterie", "puce Bluetooth"}
				if HasMaterials(player, required) {
					player.Inventory = append(player.Inventory, item)
					RemoveMaterials(&player.Inventory, required)
					player.BTC -= 5
					fmt.Println("Vous avez forgé :", item)
				}
			case 4:
				return
			}
		case "2":
			ClearScreen()
			BuyInventoryUpgrade(*player)
			return
		case "3":
			ClearScreen()
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}

}
