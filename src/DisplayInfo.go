package main

import "fmt"

func (c *Character) DisplayInfo() {
	// Sprite + Infos générales alignées à droite
	fmt.Println()
	fmt.Println("   ( O )       Name  :", c.Name)
	fmt.Println("    -_-        Class :", c.Class)
	fmt.Println("   /|_|\\       Level :", c.Level)
	fmt.Println("    / \\ ")
	fmt.Println("   /   \\ ")
	fmt.Println()

	// Encadré stats
	fmt.Println("+---------------------------+")
	fmt.Printf("| HP      : %3d / %-10d|\n", c.HP, c.MaxHP)
	fmt.Printf("| Energie : %3d / %-10d|\n", c.Energie, c.Max_Energie)
	fmt.Printf("| BTC     : %-15d |\n", c.BTC)
	fmt.Println("+---------------------------+")
	fmt.Println()

	fmt.Printf("Sorts : %v\n", c.Spells)

	fmt.Println()

	fmt.Println("Equipement :")
	fmt.Printf(" - Tête : %v\n", c.Equipements.casque)
	fmt.Printf(" - Torse: %v\n", c.Equipements.armure)
	fmt.Printf(" - Pieds: %v\n", c.Equipements.bottes)
	fmt.Printf(" - Arme : %v\n", c.Equipements.armes)
}