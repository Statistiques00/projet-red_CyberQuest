package main

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayInfo(c *Character) {
	ClearScreen()
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

	// Calcul des stats totales (personnage + équipements)
	totalFirewall := c.firewall +
		c.Equipements.casque.firewall +
		c.Equipements.armure.firewall +
		c.Equipements.bottes.firewall +
		c.Equipements.armes.firewall

	totalPuissance := c.puissance_de_calcul +
		c.Equipements.casque.puissance_de_calcul +
		c.Equipements.armure.puissance_de_calcul +
		c.Equipements.bottes.puissance_de_calcul +
		c.Equipements.armes.puissance_de_calcul

	totalStability := c.stability +
		c.Equipements.casque.stability +
		c.Equipements.armure.stability +
		c.Equipements.bottes.stability +
		c.Equipements.armes.stability

	totalVitesse := c.vitesse_de_connection +
		c.Equipements.casque.vitesse_de_connection +
		c.Equipements.armure.vitesse_de_connection +
		c.Equipements.bottes.vitesse_de_connection +
		c.Equipements.armes.vitesse_de_connection

	totalValeur := c.Equipements.casque.valeur +
		c.Equipements.armure.valeur +
		c.Equipements.bottes.valeur +
		c.Equipements.armes.valeur

	// Nouvelle section stats
	fmt.Println("+-------------------------------+")
	fmt.Println("|           STATISTIQUES        |")
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Firewall             : %3d    |\n", totalFirewall)
	fmt.Printf("| Puissance de calcul  : %3d    |\n", totalPuissance)
	fmt.Printf("| Stabilité            : %3d    |\n", totalStability)
	fmt.Printf("| Vitesse de connexion : %3d    |\n", totalVitesse)
	fmt.Printf("| Valeur               : %3d    |\n", totalValeur)
	fmt.Println("+-------------------------------+")
	fmt.Println()

	fmt.Printf("Sorts : %v\n", c.Spells)

	fmt.Println()

	fmt.Println("Equipement :")
	fmt.Printf(" - Tête : %v\n", c.Equipements.casque)
	fmt.Printf(" - Torse: %v\n", c.Equipements.armure)
	fmt.Printf(" - Pieds: %v\n", c.Equipements.bottes)
	fmt.Printf(" - Arme : %v\n", c.Equipements.armes)
	fmt.Print("\nAppuie sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
