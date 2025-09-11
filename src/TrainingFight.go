package main

import "fmt"

// Fonction qui lance un combat d'entraînement tour par tour
func TrainingFight(player Character, monster Monster, tour int) {
		if player.HP <= 0{
			fmt.Println("Vous avez été vaincu !")
		}
		if monster.HP <= 0{
			fmt.Println("Vous avez vaincu l'ennemi d'entraînement !")
			return
		}
		if tour%2 == 0{
			fmt.Println("Tour du joueur")
			CharacterTurn(&player, &monster)
		} else {
			fmt.Println("Tour du monstre")
			GoblinPattern(&monster, &player, tour)
		}
		tour++
		TrainingFight(player, monster, tour)
}
