package main

import (
	"bufio"
	"fmt"
	"os"
)

func CharacterCreation() {
	// creer un personnage ici
	// ne pas oublier les classes
	// utiliser des switch cases
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Quel est votre nom ?")
	if scanner.Scan() {
		player.Name = scanner.Text()
	}
	fmt.Println("Choisissez une classe parmi Overclockeur, Sysadmin et Netrunner")
	var class string
	if scanner.Scan() {
		class = scanner.Text()
		player.Class = class
	}
	switch class {
	case "Overclockeur":
		// set les stats
	case "Sysadmin":
		// set les stats
	case "Netrunner":
		// set les stats
	}
}