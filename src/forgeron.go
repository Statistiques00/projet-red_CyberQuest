package main

import (
	"fmt"
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
	for {
		fmt.Println("Menu Forgeron :")
		fmt.Println("1. Lunettes anti-lumière bleue")
		fmt.Println("2. Blindage électromagnétique")
		fmt.Println("3. Chaussures connectées IoT")
		fmt.Println("4. Quitter")
		var choix int
		fmt.Scan(&choix)

		if choix == 4 {
			break
		}

		if len(*inventory) >= MaxInventoryCapacity1 {
			fmt.Println("Inventaire plein !")
			continue
		}

		if *playerGold < 5 {
			fmt.Println("Pas assez d'or !")
			continue
		}

		var item string
		var required []string
		if choix == 1 {
			item = "Lunettes anti-lumière bleue"
			required = []string{"verre spécial", "isolant plastique", "monture légère"}
		} else if choix == 2 {
			item = "Blindage électromagnétique"
			required = []string{"plaque métallique", "isolant plastique", "alimentation haute capacité"}
		} else if choix == 3 {
			item = "Chaussures connectées IoT"
			required = []string{"capteurs de mouvement", "micro-batterie", "puce Bluetooth"}
		} else {
			fmt.Println("Choix invalide.")
			continue
		}

		if !hasMaterials(inventory, required) {
			fmt.Println("Ressources insuffisantes pour fabriquer", item)
			continue
		}

		*playerGold -= 5
		removeMaterials(inventory, required)
		*inventory = append(*inventory, item)
		fmt.Println(item, "fabriqué et ajouté à l'inventaire !")
	}
}
