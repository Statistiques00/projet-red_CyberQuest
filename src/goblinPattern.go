package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GoblinPattern(goblin *Monster, player *Character, nbTours int) {
	var degats int
	if nbTours%3 == 0 {
		player.HP -= goblin.Attack * 2
		degats = goblin.Attack * 2
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, degats)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)
	} else {
		player.HP -= goblin.Attack
		degats = goblin.Attack
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, degats)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)
	}
}

func HandleDrop(player *Character, monster *Monster) {
    rand.Seed(time.Now().UnixNano())

    lootFound := false
    for _, loot := range monster.LootTable {
        if rand.Float64() < loot.DropChance {
            player.Inventory = append(player.Inventory, loot.Name)
            fmt.Printf("| Vous obtenez : %-26s|\n", loot.Name)
            lootFound = true
        }
    }
    if !lootFound {
        fmt.Println("| Aucun objet trouvé sur l'ennemi.        |")
    }
    if monster.BTC > 0 {
        player.BTC += monster.BTC
        fmt.Printf("| Vous ramassez %d BTC sur le monstre.    |\n", monster.BTC)
    } else {
        fmt.Println("| Aucun BTC trouvé sur l'ennemi.          |")
    }
}