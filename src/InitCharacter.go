package main

func InitCharacter() {
	// initialiser les personnages ici
	player = Character{
		Name:                  "Hero",
		Class:                 "Warrior",
		Level:                 1,
		BTC:                   0,
		Energie:               100,
		Max_Energie:           100,
		puissance_de_calcul:   10,
		firewall:              5,
		stability:             5,
		vitesse_de_connection: 5,
		Equipements:           equipements,
		MaxHP:                 100,
		HP:                    100,
	}

}
