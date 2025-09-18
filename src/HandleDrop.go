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
	fmt.Printf("| Vous obtenez : %-26s|\n", loot.Name)
	AddInventory(*player, loot.Name)

	if monster.BTC > 0 {
		player.BTC += monster.BTC
		fmt.Printf("| Vous ramassez %d BTC sur le monstre.    |\n", monster.BTC)
	} else {
		fmt.Println("| Aucun BTC trouvé sur l'ennemi.          |")
	}
}
