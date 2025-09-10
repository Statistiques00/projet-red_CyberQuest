package main

import (
	"bufio"
	"fmt"
	"os"
)

func CharacterCreation() Character {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre nom : ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Choisissez votre classe : 1 - Humain / 2 - Elfe / 3 - Nain")
	class := ""
	// maxHP := 0
	// mana := 0
	for {
		fmt.Print("Classe : ")
		scanner.Scan()
		c := scanner.Text()
		switch c {
		case "1":
			class = "Humain"
			// maxHP = 100
			// mana = 30
		case "2":
			class = "Elfe"
			// maxHP = 80
			// mana = 50
		case "3":
			class = "Nain"
			// maxHP = 120
			// mana = 20
		default:
			fmt.Println("Choix invalide.")
			continue
		}
		break
	}
	// fmt.Println("Bienvenue, ", name, "le", class, "!")
	return Character{
		Name:  name,
		Class: class,
		// Uncomment and set these fields if needed:
		// MaxHP: maxHP,
		// Mana:  mana,
	}
}