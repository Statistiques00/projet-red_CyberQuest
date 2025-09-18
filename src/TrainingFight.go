package main

import (
	"fmt"
	"math/rand"
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
		fmt.Print("Appuie sur Entrée pour continuer...")
		fmt.Scanln()
		return
	}
	if monster.HP <= 0 {
		fmt.Println("|  Vous avez vaincu l'ennemi d'entraînement ! |")
		fmt.Println("+=========================================+")

		// DEBUG : Affiche la LootTable du monstre
		fmt.Printf("DEBUG: LootTable du monstre (%s), nombre d'items: %d\n", monster.Name, len(monster.LootTable))
		for i, item := range monster.LootTable {
			fmt.Printf("DEBUG: LootTable[%d] = %s (chance %.2f)\n", i, item.Name, item.DropChance)
		}
		fmt.Printf("DEBUG: Inventaire joueur AVANT drop : %v\n", player.Inventory)
		fmt.Printf("DEBUG: BTC joueur AVANT drop : %d\n", player.BTC)

		droppedItems := []LootItem{}
		for _, item := range monster.LootTable {
			// Exemple : 50% de chance de loot par item
			if RandInt(0, 1) == 1 {
				droppedItems = append(droppedItems, item)
			}
		}
		if len(droppedItems) > 0 {
			fmt.Println("Vous obtenez :")
			for _, item := range droppedItems {
				fmt.Printf(" - %s\n", item.Name)
				player.Inventory = append(player.Inventory, item.Name)
			}
		} else {
			fmt.Println("Aucun objet obtenu.")
		}

		// DEBUG : Drop de BTC
		fmt.Printf("DEBUG: BTC du monstre = %d\n", monster.BTC)
		if monster.BTC > 0 {
			player.BTC += monster.BTC
			fmt.Printf("Vous ramassez %d BTC sur le monstre.\n", monster.BTC)
		} else {
			fmt.Println("Aucun BTC trouvé sur l'ennemi.")
		}
		fmt.Printf("DEBUG: BTC joueur APRES drop : %d\n", player.BTC)

		fmt.Printf("DEBUG: Inventaire joueur APRES drop : %v\n", player.Inventory)
		fmt.Print("Appuie sur Entrée pour continuer...")
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

func RandInt(i1, i2 int) int {
	return rand.Intn(i2-i1+1) + i1
}
