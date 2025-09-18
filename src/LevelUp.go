package main

import "fmt"

func LevelUp(c *Character) {
	if c.XP >= c.Max_XP {
		c.Level++
		c.Max_XP *= 2
		fmt.Printf("|  Vous êtes maintenant niveau %d !       |\n", c.Level)
		fmt.Println("+=========================================+")
		fmt.Print("Vous avez gagné 1 point de compétence à répartir.")
		fmt.Println("Choisissez une statistique à augmenter :")
		var choix int
		for {
			fmt.Println("1. Augmenter les points de vie (HP) de 10")
			fmt.Println("2. Augmenter l'énergie de 10")
			fmt.Println("3. Augmenter la puissance de calcul (attaque) de 5")
			fmt.Println("4. Augmenter le firewall (défense) de 5")
			fmt.Print("Votre choix (1-4) : ")
			_, err := fmt.Scan(&choix)
			if err != nil {
				fmt.Println("Entrée invalide. Veuillez réessayer.")
				continue
			}
			switch choix {
			case 1:
				c.MaxHP += 10
			case 2:
				c.Max_Energie += 10
			case 3:
				c.puissance_de_calcul += 5
			case 4:
				c.firewall += 5
			default:
				fmt.Println("Choix invalide. Veuillez réessayer.")
				continue
			}
			break
		}
	} else {
		fmt.Printf("|  Vous avez %d/%d XP. Continuez comme ça !   |\n", c.XP, c.Max_XP)
		fmt.Println("+=========================================+")
	}
}
