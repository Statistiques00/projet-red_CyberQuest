package main

import (
	"bufio"
	"fmt"
	"os"
)

func CharacterCreation() Character {
	scanner := bufio.NewScanner(os.Stdin)
	ClearScreen()
	fmt.Print(`
+===================================+ 
|    === CREATION PERSONNAGE ===    | 
+===================================+ 
| Entrez votre nom :`)
scanner.Scan()
name := scanner.Text()
fmt.Print(`
+===================================+
`)
	ClearScreen()
	fmt.Printf(`
+===================================+
|    === CREATION PERSONNAGE ===    |
+===================================+
| Entrez votre nom : %s%s|
+===================================+
`, name, spaces(27-len(name))) // pour aligner le cadre

	fmt.Print(`
+=================================+
| Choisissez votre classe :       |
|   1 - Humain  (HP:100 / Mana:30)|
|   2 - Elfe    (HP: 80 / Mana:50)|
|   3 - Nain    (HP:120 / Mana:20)|
+=================================+
Classe : `)

	class := ""
	for {
		scanner.Scan()
		c := scanner.Text()
		switch c {
		case "1":
			class = "Humain"
		case "2":
			class = "Elfe"
		case "3":
			class = "Nain"
		default:
			fmt.Print("Choix invalide. Réessayez : ")
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
	fmt.Printf("| Nom    : %-25s |\n", name)
	fmt.Printf("| Classe : %-25s |\n", class)
	fmt.Println("+===================================+")
	fmt.Print("\nAppuie sur Entrée pour continuer...")
	scanner.Scan() // Attend que l'utilisateur appuie sur Entrée

	return Character{
		Name:  name,
		Class: class,
	}
}

// Ajoute cette fonction utilitaire si tu ne l'as pas déjà
func spaces(n int) string {
	if n <= 0 {
		return ""
	}
	return fmt.Sprintf("%*s", n, "")
}
