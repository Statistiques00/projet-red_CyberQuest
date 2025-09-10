package main

import "fmt"

type Equipement struct {
	Chapeau string
	Tunique string
	Bottes  string
}

type Joueur struct {
	HPMax      int
	Equipement Equipement
	Inventory  []string
}

// Fonction pour équiper un objet
func Equiper(j *Joueur, objet string) {
	var section string
	var bonus int

	switch objet {
	case "Chapeau de l’aventurier":
		section = "Chapeau"
		bonus = 10
	case "Tunique de l’aventurier":
		section = "Tunique"
		bonus = 25
	case "Bottes de l’aventurier":
		section = "Bottes"
		bonus = 15
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}

	// Vérifie si l'objet est dans l'inventaire
	found := false
	for i, item := range j.Inventory {
		if item == objet {
			// Retire l'objet de l'inventaire
			j.Inventory = append(j.Inventory[:i], j.Inventory[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Vous ne possédez pas cet équipement dans votre inventaire.")
		return
	}

	// Si déjà équipé, retire le bonus et remet l'ancien équipement dans l'inventaire
	switch section {
	case "Chapeau":
		if j.Equipement.Chapeau != "" {
			j.Inventory = append(j.Inventory, j.Equipement.Chapeau)
			j.HPMax -= 10 // retire l'ancien bonus
		}
		j.Equipement.Chapeau = objet
	case "Tunique":
		if j.Equipement.Tunique != "" {
			j.Inventory = append(j.Inventory, j.Equipement.Tunique)
			j.HPMax -= 25
		}
		j.Equipement.Tunique = objet
	case "Bottes":
		if j.Equipement.Bottes != "" {
			j.Inventory = append(j.Inventory, j.Equipement.Bottes)
			j.HPMax -= 15
		}
		j.Equipement.Bottes = objet
	}

	// Ajoute le bonus de PV max
	j.HPMax += bonus
	fmt.Printf("%s équipé ! PV max : %d\n", objet, j.HPMax)
}
