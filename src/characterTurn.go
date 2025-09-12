package main

import (
	"fmt"
)

// Fonction qui simule le tour de jeu du joueur
func CharacterTurn(player *Character, monster *Monster) {
	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("+=========================================+")
		fmt.Println("|             MENU DE COMBAT              |")
		fmt.Println("+=========================================+")
		fmt.Println("| 1 - Attaquer                            |")
		fmt.Println("| 2 - Inventaire                          |")
		fmt.Println("+=========================================+")
		fmt.Print("Choisissez une action : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			ClearScreen()
			fmt.Println("+=========================================+")
			fmt.Println("|           CHOISISSEZ UN SORT            |")
			fmt.Println("+=========================================+")
			for i, spell := range player.Spells {
				fmt.Printf("| %d - %-35s|\n", i+1, spell.nom)
			}
			fmt.Println("+=========================================+")
			var selec int
			fmt.Printf("Choisissez un sort à utiliser (1-%d) : ", len(player.Spells))
			fmt.Scan(&selec)
			if selec < 1 || selec > len(player.Spells) {
				fmt.Println("Choix invalide. Veuillez réessayer.")
				fmt.Print("Appuie sur Entrée pour continuer...")
				fmt.Scanln()
				continue
			}
			if player.Spells[selec-1].cost > player.Energie {
				fmt.Println("Pas assez d'énergie pour lancer ce sort.")
				fmt.Print("Appuie sur Entrée pour continuer...")
				fmt.Scanln()
				continue
			}
			monster.HP -= player.Spells[selec-1].degats + player.puissance_de_calcul
			player.Energie -= player.Spells[selec-1].cost

			fmt.Printf("| %s utilise %s et inflige %d dégâts !\n", player.Name, player.Spells[selec-1].nom, player.Spells[selec-1].degats)
			fmt.Printf("| PV restants du monstre : %d/%d\n", monster.HP, monster.MaxHP)
			fmt.Println("+=========================================+")
			fmt.Print("Appuie sur Entrée pour continuer...")
			fmt.Scanln()
			return
		case 2:
			AccessInventory(*player)
			return
		default:
			fmt.Println("| Choix invalide. Veuillez réessayer.     |")
			fmt.Print("Appuie sur Entrée pour continuer...")
			fmt.Scanln()
		}
	}
}