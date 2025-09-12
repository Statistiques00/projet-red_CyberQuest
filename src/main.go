package main

import (
	"bufio"
	"fmt"
	"os"
)

type Casque struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
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
	MaxInventory          int
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
	LootTable  []LootItem 
	BTC        int
}

type LootItem struct {
	Name       string
	DropChance float64 
}

var player Character

var Spells1 = Spell{
	nom:         "Attaque basique",
	cost:        0,
	description: "Une attaque simple mais efficace.",
	degats:      2,
}

var Spells2 = Spell{
	nom:         "Attaque puissante",
	cost:        10,
	description: "Une attaque qui inflige plus de dégâts.",
	degats:      5,
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
	CharacterCreation()
	for {
		ClearScreen()
		fmt.Print(`
		+===============================+
		|    ===  MENU PRINCIPAL ===    |
		+===============================+
		| 1 - Infos personnage          |
		| 2 - Inventaire                |
		| 3 - Marchand                  |
		| 4 - Forgeron                  |
		| 5 - Entrainement              |
		| 6 - Quitter                   |
		+===============================+
		Choix : `)
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			DisplayInfo(&player)
		case "2":
			AccessInventory(player)
		case "3":
			fmt.Println("Marchand")
			AccessMarchand(&player)
		case "4":
			fmt.Println("Forgeron")
			ForgeronMenu(&player.BTC, &player.Inventory)
		case "5":
			trainingMonster := Monster{
				Name:       "Ennemi d'entraînement",
				MaxHP:      30,
				HP:         30,
				Attack:     5,
				Defense:    2,
				Speed:      3,
				Experience: 10,
				LootTable:  []LootItem{},
				BTC:        0,
			}
			TrainingFight(player, trainingMonster, 1)
		case "6":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
		scanner.Scan()
	}
}
