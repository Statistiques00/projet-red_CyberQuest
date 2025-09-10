package main

import "fmt"

// Fonction de pattern de combat du Gobelin d'entraînement
func GoblinPattern(goblin Monster, player *Player, nbTours int) {
	for tour := 1; tour <= nbTours; tour++ {
		var degats int
		if tour%3 == 0 {
			degats = goblin.Attack * 2
		} else {
			degats = goblin.Attack
		}
		player.HP -= degats
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, degats)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)
	}
}
