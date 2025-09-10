package main

import "fmt"

func (c *Character) DisplayInfo() {
	fmt.Println("\n--- Infos du personnage ---")
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nHP: %d/%d\nMana: %d\nOr: %d\n", c.Name, c.Class, c.Level, c.HP, c.MaxHP, c.Energie, c.Max_Energie)
	fmt.Printf("BTC: %d\n", c.BTC)
	fmt.Println("Sorts:", c.Spells)
	fmt.Println("Equipement: Tête:", equipements.casque, ", Torse:", equipements.armure, ", Pieds:", equipements.bottes, ", Arme:", equipements.armes)
}