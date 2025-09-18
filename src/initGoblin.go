package main

import (
	"math/rand"
	"time"
)

// Fonction pour initialiser un Gobelin d'entraînement
func InitGoblin() Monster {
	rand.Seed(time.Now().UnixNano()) // Initialiser la graine

	min, max := 1, 3
	randomNumber := rand.Intn(max-min+1) + min

	switch randomNumber {
	case 1:
		return Monster{
			Name:       "Firewall Golem",
			MaxHP:      200,
			HP:         200,
			Poisoned:   0,
			Attack:     5,
			Defense:    20,
			Speed:      10,
			Experience: 50,
			LootTable: []LootItem{
				{Name: "micro-batterie", DropChance: 1.0},
				{Name: "plaque métallique", DropChance: 1.0},
				{Name: "isolant plastique", DropChance: 1.0},
				{Name: "alimentation haute capacité", DropChance: 1.0},
				{Name: "capteurs de mouvement", DropChance: 1.0},
				{Name: "puce Bluetooth", DropChance: 1.0},
				{Name: "monture légère", DropChance: 1.0},
				{Name: "verre spécial", DropChance: 1.0},
				{Name: "capteur de mouvement", DropChance: 1.0},
			},
			BTC: 5,
		}
	case 2:
		return Monster{
			Name:       "Data Imp",
			MaxHP:      150,
			HP:         150,
			Poisoned:   0,
			Attack:     10,
			Defense:    10,
			Speed:      20,
			Experience: 50,
			LootTable: []LootItem{
				{Name: "micro-batterie", DropChance: 1.0},
				{Name: "plaque métallique", DropChance: 1.0},
				{Name: "isolant plastique", DropChance: 1.0},
				{Name: "alimentation haute capacité", DropChance: 1.0},
				{Name: "capteurs de mouvement", DropChance: 1.0},
				{Name: "puce Bluetooth", DropChance: 1.0},
				{Name: "monture légère", DropChance: 1.0},
				{Name: "verre spécial", DropChance: 1.0},
				{Name: "capteur de mouvement", DropChance: 1.0},
			},
			BTC: 5,
		}
	default:
		return Monster{
			Name:       "Script Kiddie",
			MaxHP:      100,
			HP:         100,
			Poisoned:   0,
			Attack:     15,
			Defense:    5,
			Speed:      30,
			Experience: 50,
			LootTable: []LootItem{
				{Name: "micro-batterie", DropChance: 1.0},
				{Name: "plaque métallique", DropChance: 1.0},
				{Name: "isolant plastique", DropChance: 1.0},
				{Name: "alimentation haute capacité", DropChance: 1.0},
				{Name: "capteurs de mouvement", DropChance: 1.0},
				{Name: "puce Bluetooth", DropChance: 1.0},
				{Name: "monture légère", DropChance: 1.0},
				{Name: "verre spécial", DropChance: 1.0},
				{Name: "capteur de mouvement", DropChance: 1.0},
			},
			BTC: 5,
		}
	}
}
