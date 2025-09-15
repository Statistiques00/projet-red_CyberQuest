package main

func InitCharacter(c rune) *Character {
	// initialiser les personnages ici
	switch c {
	case '1':
		player = Character{
			Class:                 "Overclocker",
			Level:                 1,
			BTC:                   0,
			Energie:               100,
			Statut:                0,
			Max_Energie:           100,
			puissance_de_calcul:   10,
			firewall:              5,
			stability:             5,
			vitesse_de_connection: 5,
			Equipements:           equipements,
			MaxHP:                 100,
			HP:                    100,
		}
	case '2':
		player = Character{
			Class:                 "SysAdmin",
			Level:                 1,
			BTC:                   0,
			Energie:               100,
			Statut:                0,
			Max_Energie:           100,
			puissance_de_calcul:   10,
			firewall:              5,
			stability:             5,
			vitesse_de_connection: 5,
			Equipements:           equipements,
			MaxHP:                 100,
			HP:                    100,
		}
	case '3':
		player = Character{
			Class:                 "Netrunner",
			Level:                 1,
			BTC:                   0,
			Statut:                0,
			Energie:               100,
			Max_Energie:           100,
			puissance_de_calcul:   10,
			firewall:              5,
			stability:             5,
			vitesse_de_connection: 5,
			Equipements:           equipements,
			MaxHP:                 100,
			HP:                    100,
			Spells:                []Spell{Spells1, Spells2},
		}
	}
	return nil
}
