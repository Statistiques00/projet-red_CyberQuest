package main

import (
	"bufio"
	"fmt"
	"os"
)

func AccessInventory(c *Character) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		PrintCentered("\n--- INVENTAIRE ---")
		if len(c.Inventory) == 0 {
			PrintCentered("Inventaire vide.")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("%d : %s\n", i+1, item)
			}
		}
		fmt.Println("a - Ajouter un objet")
		fmt.Println("r - Retirer un objet")
		fmt.Println("u - Utiliser une potion de vie")
		fmt.Println("p - Utiliser une potion de poison")
		fmt.Println("q - Quitter l'inventaire")
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "a":
			fmt.Print("Nom de l'objet à ajouter : ")
			scanner.Scan()
			AddInventory(c, scanner.Text())
		case "u":
			TakePot(c)
		case "q":
			ClearScreen()
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}