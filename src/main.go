package main

import (
	"bufio"
	"fmt"
	"os"
)

// importer les packages necessaires ici
type Equipment struct {
	// definir les equipements ici
	Head   string
	Toros  string
	Legs   string
	Feet   string
	Hands  string
	Weapon string
	Shield string
}

type Character struct {
	// definir les attributs des personnages ici
	Name      string
	Class     string
	Level     int
	Gold      int
	Equipment Equipment
	Mana      int
	Spells    []string
	MaxHP     int
	HP        int
	Defense   int
}

type Monster struct {
	// definir les attributs des monstres ici
	Name       string
	MaxHP      int
	HP         int
	Attack     int
	Defense    int
	Speed      int
	Experience int
	Loot       []string
}

// Declare player as a package-level variable
var player Character

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearScreen()
		PrintIndented("=== MENU ===", 25)
		PrintIndented("Bienvenue dans le jeu !", 20)
		PrintIndented("1 - Option 1", 23)
		PrintIndented("2 - Option 2", 23)
		PrintIndented("3 - Quitter", 23)
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			DisplayInfo()
			AddInventory()
		case "2":
			fmt.Println("Tu as choisi l'option 2")
		case "3":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
		scanner.Scan()
	}
}
