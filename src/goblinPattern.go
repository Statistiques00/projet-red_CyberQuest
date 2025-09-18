package main

import (
	"fmt"
	"math/rand"
	"time"
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

func HandleDrop(player *Character, monster *Monster) {
    rand.Seed(time.Now().UnixNano())

    fmt.Printf("DEBUG: LootTable du monstre (%s), nombre d'items: %d\n", monster.Name, len(monster.LootTable))
    for _, loot := range monster.LootTable {
        fmt.Printf(" - %s\n", loot.Name)
    }

    lootGiven := false
    if len(monster.LootTable) > 0 {
        loot := monster.LootTable[0] // drop garanti du premier item
        fmt.Printf("| Vous obtenez : %-26s|\n", loot.Name)
        AddEquipmentToPlayer(player, loot.Name)
        lootGiven = true
    }
    if !lootGiven {
        fmt.Println("| Aucun objet trouvé sur l'ennemi.        |")
    }

    fmt.Printf("DEBUG: BTC du monstre = %d\n", monster.BTC)
    if monster.BTC > 0 {
        player.BTC += monster.BTC
        fmt.Printf("| Vous ramassez %d BTC sur le monstre.    |\n", monster.BTC)
    } else {
        fmt.Println("| Aucun BTC trouvé sur l'ennemi.          |")
    }

    // Affiche l'inventaire du joueur pour debug
    fmt.Printf("DEBUG: Inventaire joueur : %v\n", player.Inventory)
}

// Ajoute l'équipement au joueur selon le nom
func AddEquipmentToPlayer(player *Character, itemName string) {
    // Casque
    if itemName == "Casque de base" {
        player.Equipements.casque = equipements.casque
    }
    // Armure
    if itemName == "Armure de base" {
        player.Equipements.armure = equipements.armure
    }
    // Bottes
    if itemName == "Bottes de base" {
        player.Equipements.bottes = equipements.bottes
    }
    // Armes
    if itemName == "Epée de base" {
        player.Equipements.armes = equipements.armes
    }
    // Ajoute aussi le nom dans l'inventaire texte si tu veux
    player.Inventory = append(player.Inventory, itemName)
}