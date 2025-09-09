package main

import "fmt"

func (c *Character) DisplayInfo() {
	fmt.Println("\n--- Infos du personnage ---")
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nHP: %d/%d\nMana: %d\nOr: %d\n", c.Name, c.Class, c.Level, c.HP, c.MaxHP, c.Mana, c.Gold)
	fmt.Println("Sorts:", c.Spells)
	fmt.Println("Equipement: Tête:", c.Equipment.Head, ", Torse:", c.Equipment.Torso, ", Pieds:", c.Equipment.Feet, ")")
}