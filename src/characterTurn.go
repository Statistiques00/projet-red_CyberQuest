package main

import (
	"fmt"
)

// Fonction qui simule le tour de jeu du joueur
func CharacterTurn(player *Character, monster *Monster) {
	for {
		fmt.Println("Menu de combat :")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choisissez une action : ")
		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			// Attaque basique
			degats := 5
			monster.HP -= degats
			if monster.HP < 0 {
				monster.HP = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts !\n", player.Name, degats)
			fmt.Printf("PV restants du monstre : %d/%d\n", monster.HP, monster.MaxHP)
			break // Fin du tour du joueur, passage au tour du monstre
		} else if choix == 2 {
			fmt.Println("Ouverture de l'inventaire...")
			// Ajoutez ici la logique d'inventaire
		} else {
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}
