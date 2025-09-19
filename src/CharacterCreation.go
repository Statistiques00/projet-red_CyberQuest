package main

import (
	"bufio"
	"fmt"
	"os"
)

func CharacterCreation() {
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	var class string
	for {
		ClearScreen()
		fmt.Print(`
+===================================+ 
|    === CREATION PERSONNAGE ===    | 
+===================================+ 
| Entrez votre nom : 		    |
+===================================+
`)
		fmt.Print("\033[2A")
		fmt.Print("\033[21C")
		scanner.Scan()
		name = scanner.Text()
		if len(name) == 0 {
			ClearScreen()
			fmt.Print(`
+===================================+
|    === CREATION PERSONNAGE ===    |
+===================================+
|      Nom du personnage manquant   |
+===================================+

+===================================+
| Appuie sur Entrée pour réessayer  |
+===================================+
`)
			scanner.Scan()
			continue
		}
		break
	}

	for {
		ClearScreen()
		const cadreLargeur = 35
		fmt.Printf(`
+===================================+
|    === CREATION PERSONNAGE ===    |
+===================================+
| Entrez votre nom : %-*s|
+===================================+
`, cadreLargeur-len(" Entrez votre nom : "), name)

		fmt.Print(`
+==========================================+
|  Choisissez votre classe :               |
|    1 - Overclocker  (HP:100 / Mana:100)  |
|    2 - Sysadmin   (HP: 130 / Mana:80)    |
|    3 - Netrunner   (HP:80 / Mana:130)    |
+==========================================+
Classe : `)

		scanner.Scan()
		c := scanner.Text()
		switch c {
		case "1":
			class = "Overclocker"
			InitCharacter(rune('1'))
		case "2":
			class = "SysAdmin"
			InitCharacter(rune('2'))
		case "3":
			class = "Netrunner"
			InitCharacter(rune('3'))
		default:
			class = ""
		}
		if len(class) == 0 {
			ClearScreen()
			fmt.Print(`
+===================================+
|    === CREATION PERSONNAGE ===    |
+===================================+
|        Classe du personnage       |
|             manquante             |
+===================================+

+===================================+
| Appuie sur Entrée pour réessayer  |
+===================================+
`)
			scanner.Scan()
			continue
		}
		break
	}

	// Affichage du récapitulatif
	ClearScreen()
	fmt.Print(`
+===================================+
|      RECAPITULATIF PERSONNAGE     |
+===================================+
`)
	fmt.Printf("| Nom    : %-25s|\n", name)
	fmt.Printf("| Classe : %-25s|\n", class)
	fmt.Println("+===================================+")
	fmt.Print("\nAppuie sur Entrée pour continuer...")
	scanner.Scan() // Attend que l'utilisateur appuie sur Entrée
}

// Fonction utilitaire pour l'alignement
func spaces(n int) string {
	if n <= 0 {
		return ""
	}
	return fmt.Sprintf("%*s", n, "")
}

