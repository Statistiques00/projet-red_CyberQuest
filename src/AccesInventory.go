package main

import (
	"bufio"
	"fmt"
	"os"
)

func AccessInventory(c *Character) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearScreen()
		fmt.Println("+=========================================+")
		fmt.Println("|           === INVENTAIRE ===            |")
		fmt.Println("+=========================================+")
		if len(c.Inventory) == 0 {
			fmt.Println("|                                         |")
			fmt.Println("|         Votre inventaire est vide.      |")
			fmt.Println("|                                         |")
			fmt.Println("+=========================================+")
			fmt.Print("Appuie sur Entrée pour revenir au menu...")
			scanner.Scan()
			return
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("| %2d : %-35s|\n", i+1, item)
			}
		}
		fmt.Println("+=========================================+")
		fmt.Println("| 1 - Ajouter un objet                    |")
		fmt.Println("| 2 - Retirer un objet                    |")
		fmt.Println("| 3 - Utiliser une potion de vie          |")
		fmt.Println("| 4 - Quitter l'inventaire                |")
		fmt.Println("+=========================================+")
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			fmt.Print("Nom de l'objet à ajouter : ")
			scanner.Scan()
			AddInventory(*c, scanner.Text())
		case "2":
			fmt.Print("Numéro de l'objet à retirer : ")
			scanner.Scan()
			RemoveInventory(&c, scanner.Text())
		case "3":
			TakePot(c)
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "4":
			return
		default:
			fmt.Println("Choix invalide.")
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		}
	}
}
