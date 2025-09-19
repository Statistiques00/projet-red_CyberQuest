package main

import (
	"bufio"
	"fmt"
	"os"
)

func AccessMarchand(c *Character) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearScreen()
		fmt.Println("+==========================================+")
		fmt.Println("|             === MARCHAND ===             |")
		fmt.Println("+==========================================+")
		fmt.Printf("| BTC : %-34d |\n", c.BTC)
		fmt.Println("+==========================================+")
		fmt.Println("| 1 - Acheter une potion de vie   (10 BTC) |")
		fmt.Println("| 2 - Acheter une potion de poison (8 BTC) |")
		fmt.Println("| 3 - Vendre un objet                      |")
		fmt.Println("| 4 - Quitter le marchand                  |")
		fmt.Println("+==========================================+")
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			if c.BTC >= 10 {
				c.BTC -= 10
				c.Inventory = append(c.Inventory, "Potion de vie")
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|          Potion de vie achetée !        |")
				fmt.Println("+=========================================+")
			} else {
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|           Pas assez de BTC !            |")
				fmt.Println("+=========================================+")
			}
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "2":
			if c.BTC >= 8 {
				c.BTC -= 8
				c.Inventory = append(c.Inventory, "Potion de poison")
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|      Potion de poison achetée !         |")
				fmt.Println("+=========================================+")

				} else {
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|          Pas assez de BTC !             |")
				fmt.Println("+=========================================+")
			}
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "3":
			if len(c.Inventory) == 0 {
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|          Votre inventaire est vide.     |")
				fmt.Println("+=========================================+")
				fmt.Println()
				fmt.Print("Appuie sur Entrée pour continuer...")
				scanner.Scan()
				continue
			}
			ClearScreen()
			fmt.Println()
			fmt.Println("+=========================================+")
			fmt.Println("| Sélectionnez l'objet à vendre :         |")
			for i, item := range c.Inventory {
				fmt.Printf("| %2d : %-35s|\n", i+1, item)
			}
			fmt.Println("+=========================================+")
			fmt.Println()
			fmt.Print("Numéro de l'objet : ")
			scanner.Scan()
			var idx int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &idx)
			if err == nil && idx > 0 && idx <= len(c.Inventory) {
				objet := c.Inventory[idx-1]
				c.Inventory = append(c.Inventory[:idx-1], c.Inventory[idx:]...)
				c.BTC += 15 
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Printf("| %s vendu pour 15 BTC !%s |\n", objet, spaces(34-len(objet)-15))
				fmt.Println("+=========================================+")
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|           Numéro invalide.              |")
				fmt.Println("+=========================================+")
				fmt.Println()
			}
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "4":
			return
		default:
			ClearScreen()
			fmt.Println()
				fmt.Println("+=========================================+")
				fmt.Println("|           Choix Invalide.               |")
				fmt.Println("+=========================================+")
				fmt.Println()
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		}
	}
}
