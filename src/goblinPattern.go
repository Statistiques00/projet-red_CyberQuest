package main

import (
	"fmt"
)

func GoblinPattern(goblin *Monster, player *Character, nbTours int) {
	var degats int
    player.Energie += 10
    if player.Energie > player.Max_Energie {
        player.Energie = player.Max_Energie
    }
	if nbTours%3 == 0 {
		player.HP -= goblin.Attack * 2
		degats = goblin.Attack * 2
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, degats)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)
	} else {
		player.HP -= goblin.Attack - player.firewall
		degats = goblin.Attack - player.firewall
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, degats)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)
	}
}




