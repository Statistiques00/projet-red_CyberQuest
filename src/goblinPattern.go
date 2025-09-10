package main

import "fmt"

// Fonction de pattern de combat du Gobelin d'entraînement
func GoblinPattern(goblin Monster, player *Character, nbTours int) {
	var degats int
	if nbTours%3 == 0 {
		player.HP -= goblin.Attack * 2
	} else {
		player.HP -= goblin.Attack
	}
	fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, degats)
	fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)
}
