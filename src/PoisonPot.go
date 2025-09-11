package main

import (
	"fmt"
	"time"
)

func PoisonPot(target *Monster) {
	for i := 0; i < 3; i++ {
		target.HP -= 10
		if target.HP < 0 {
			target.HP = 0
		}
		fmt.Printf("PV: %d/%d\n", target.HP, target.MaxHP)
		time.Sleep(1 * time.Second)
	}
}

// Ajout de la potion de poison dans le marchand
func BuyPoisonPotion(inventory *[]string, playerGold *int) bool {
	if *playerGold >= 20 { // prix exemple : 20 pièces d'or
		*playerGold -= 20
		*inventory = append(*inventory, "Potion de poison")
		return true
	}
	return false
}
