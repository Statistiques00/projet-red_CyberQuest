package main

import (
	"bufio"
	"fmt"
	"os"
)

// importer les packages necessaires ici
type Casque struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int

	// definir les equipements ici
}

type Armure struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Bottes struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Armes struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Equipements struct {
	// stuff Equipé
	casque Casque
	armure Armure
	bottes Bottes
	armes  Armes
}

type Spell struct {
	nom         string
	cost        int
	description string
	degats      int
}
type Character struct {
	// definir les attributs des personnages ici
	Name                  string
	Class                 string
	Level                 int
	BTC                   int
	puissance_de_calcul   int //attaque
	firewall              int //defense
	stability             int //tx critique
	vitesse_de_connection int //initiative
	Energie               int //mana
	Max_Energie           int //max mana
	MaxHP                 int
	HP                    int
	Inventory             []string
	Equipements           Equipements
	Spells                []Spell
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
	BTC        int
}

// Declare player as a package-level variable
var player Character

var Spells1 = Spell{
	nom:         "Attaque basique",
	cost:        0,
	description: "Une attaque simple mais efficace.",
	degats:      10,
}

var equipements = Equipements{
	casque: Casque{
		nom:                   "Casque de base",
		firewall:              5,
		puissance_de_calcul:   0,
		stability:             0,
		vitesse_de_connection: 0,
		valeur:                10,
	},
	armure: Armure{
		nom:                   "Armure de base",
		firewall:              15,
		puissance_de_calcul:   0,
		stability:             0,
		vitesse_de_connection: -1,
		valeur:                30,
	},
	bottes: Bottes{
		nom:                   "Bottes de base",
		firewall:              5,
		puissance_de_calcul:   0,
		stability:             0,
		vitesse_de_connection: 2,
		valeur:                20,
	},
	armes: Armes{
		nom:                   "Epée de base",
		firewall:              0,
		puissance_de_calcul:   10,
		stability:             5,
		vitesse_de_connection: 0,
		valeur:                25,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	player = CharacterCreation()
	for {
		ClearScreen()
		PrintIndented("=== MENU ===", 25)
		PrintIndented("Bienvenue dans le jeu !", 20)
		PrintIndented("1 - Info", 23)
		PrintIndented("2 - Option 2", 23)
		PrintIndented("3 - Quitter", 23)
		PrintIndented("4 - Marchand", 23)
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			player.DisplayInfo()
		case "2":
			AccessInventory(&player)
		case "3":
			fmt.Println("Au revoir !")
			return
		case "4":
			fmt.Println("Marchand")
			marchand(&player)
		default:
			fmt.Println("Choix invalide.")
		}
		scanner.Scan()
	}
}
