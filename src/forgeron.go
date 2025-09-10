package main

import (
	"bufio"
	"fmt"
	"os"
)

var MaxInventoryCapacity1 int = 20 // à adapter selon votre gestion

func hasMaterials(inventory *[]string, required []string) bool {
	inv := make(map[string]int)
	for _, item := range *inventory {
		inv[item]++
	}
	for _, mat := range required {
		if inv[mat] == 0 {
			return false
		}
		inv[mat]--
	}
	return true
}

func removeMaterials(inventory *[]string, required []string) {
	for _, mat := range required {
		for i, item := range *inventory {
			if item == mat {
				*inventory = append((*inventory)[:i], (*inventory)[i+1:]...)
				break
			}
		}
	}
}

func ForgeronMenu(playerGold *int, inventory *[]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearScreen()
		fmt.Println("+=========================================+")
		fmt.Println("|            === FORGERON ===             |")
		fmt.Println("+=========================================+")
		fmt.Printf("| Or : %-34d|\n", *playerGold)
		fmt.Println("+=========================================+")
		fmt.Println("| 1 - Améliorer une arme   (20 Or)        |")
		fmt.Println("| 2 - Améliorer une armure (15 Or)        |")
		fmt.Println("| 3 - Réparer un objet     (10 Or)        |")
		fmt.Println("| q - Quitter le forgeron                 |")
		fmt.Println("+=========================================+")
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			if *playerGold >= 20 {
				*playerGold -= 20
				fmt.Println("| Arme améliorée !                        |")
			} else {
				fmt.Println("| Pas assez d'or !                        |")
			}
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "2":
			if *playerGold >= 15 {
				*playerGold -= 15
				fmt.Println("| Armure améliorée !                      |")
			} else {
				fmt.Println("| Pas assez d'or !                        |")
			}
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "3":
			if *playerGold >= 10 {
				*playerGold -= 10
				fmt.Println("| Objet réparé !                          |")
			} else {
				fmt.Println("| Pas assez d'or !                        |")
			}
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		case "q":
			return
		default:
			fmt.Println("| Choix invalide.                         |")
			fmt.Print("Appuie sur Entrée pour continuer...")
			scanner.Scan()
		}
	}
}
