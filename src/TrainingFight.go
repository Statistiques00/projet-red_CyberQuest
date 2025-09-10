package main

import "fmt"

// Fonction qui lance un combat d'entraînement tour par tour
func TrainingFight(player *Character, monster Monster) {
	tour := 1
	for player.HP > 0 && monster.HP > 0 {
		fmt.Printf("--- Tour %d ---\n", tour)
		// Tour du joueur
		CharacterTurn(player, &monster)
		if monster.HP <= 0 {
			fmt.Println("Le monstre est vaincu !")
			break
		}
		// Tour du monstre
		GoblinPattern(monster, player, 1)
		if player.HP <= 0 {
			fmt.Println("Vous avez été vaincu !")
			break
		}
		tour++
	}
}
