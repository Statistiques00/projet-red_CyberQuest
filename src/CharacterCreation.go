package main

import (
	"bufio"
	"fmt"
	"os"
)

func CharacterCreation() {
	scanner := bufio.NewScanner(os.Stdin)
	ClearScreen()
	fmt.Print(`
+===================================+ 
|    === CREATION PERSONNAGE ===    | 
+===================================+ 
| Entrez votre nom : 		    |
+===================================+
`)

	fmt.Print("\033[2A")  // Remonte d'une ligne
	fmt.Print("\033[21C") // Avance de 21 colonnes (après "Entrez votre nom : ")

	scanner.Scan()
	name := scanner.Text()

	// Affiche la suite avec le nom saisi
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

	class := ""

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
		fmt.Print("Choix invalide. Réessayez : ")

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
