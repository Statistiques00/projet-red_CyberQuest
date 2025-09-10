package main

import (
	"bufio"
	"fmt"
	"os"
)

func marchand(c *Character) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- MARCHAND ---")
		fmt.Println("BTC :", c.BTC)
		fmt.Println("1 - Potion de vie (15 BTC)")
		fmt.Println("q - Quitter")
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			if c.BTC >= 15 {
				AddInventory(c, "Potion de vie")
				c.BTC -= 15
			} else {
				fmt.Println("Pas assez de BTC.")
			}
		case "q":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}