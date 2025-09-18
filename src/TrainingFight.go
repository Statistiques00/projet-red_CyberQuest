package main

import (
	"fmt"
)

func TrainingFight(player *Character, monster Monster, tour int) {
	ClearScreen()
	fmt.Println("+=========================================+")
	fmt.Println("|         === ENTRAINEMENT ===            |")
	fmt.Println("+=========================================+")
	fmt.Printf("| Joueur : %-30s |\n", player.Name)
	fmt.Printf("| Classe : %-30s |\n", player.Class)
	fmt.Printf("| HP     : %3d / %-24d |\n", player.HP, player.MaxHP)
	fmt.Println("+-----------------------------------------+")
	fmt.Printf("| Ennemi : %-30s |\n", monster.Name)
	fmt.Printf("| HP     : %3d / %-24d |\n", monster.HP, monster.MaxHP)
	fmt.Println("+=========================================+")

	if player.HP <= 0 {
		fmt.Println("|         Vous avez été vaincu !          |")
		fmt.Println("+=========================================+")
		fmt.Println()
		fmt.Print("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return
	}
	if monster.HP <= 0 {
		fmt.Println()
		fmt.Println("+=========================================+")
		fmt.Println("|      Vous avez vaincu l'entraînement    |")
		fmt.Println("+=========================================+")
		HandleDrop(player, &monster)
		fmt.Println("+=========================================+")
		LevelUp(player)
		fmt.Println("Appuie sur Entrée pour continuer...")
		// fmt.Println("+=========================================+")
		// fmt.Print("Appuie sur Entrée pour continuer...")
		player.XP += monster.Experience
		LevelUp(player)
		fmt.Scanln()
		return
	}
	if tour%2 == 0 {
		fmt.Println("|           Tour du joueur                |")
		fmt.Println("+=========================================+")
		CharacterTurn(*player, &monster)
	} else {
		fmt.Println("|           Tour du monstre               |")
		fmt.Println("+=========================================+")
		GoblinPattern(&monster, player, tour)
	}
	tour++

	PoisonStatut(&monster)
	TrainingFight(player, monster, tour)
}
