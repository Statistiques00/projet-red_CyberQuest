package main

import (
	"fmt"
	"math/rand"
)

func HandleDrop(player *Character, monster *Monster) {
	max := len(monster.LootTable) - 1
	min := 0
	n := rand.Intn(max-min+1) + min
	loot := monster.LootTable[n]

	// Affichage avec retour à la ligne si le nom est trop long
	maxLen := 26
	name := loot.Name
	for len(name) > maxLen {
		fmt.Printf("| %s |\n", name[:maxLen])
		name = name[maxLen:]
	}
	fmt.Printf("| Vous obtenez :%-*s|\n", maxLen, name)
	fmt.Println("+=========================================+")

	AddInventory(player, loot.Name)

	if monster.BTC > 0 {
		player.BTC += monster.BTC
		fmt.Println()
		fmt.Println("+=========================================+")
		fmt.Printf("| Vous ramassez %d BTC sur le monstre.    |\n", monster.BTC)
	} else {
		fmt.Println("| Aucun BTC trouvé sur l'ennemi.          |")
	}
}
