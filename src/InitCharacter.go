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
			Spells:                []Spell{Overclock, CodeInjection, Pingflood, DataSurge, KernelPanic},
			Statut:                0,
			Max_Energie:           100,
			puissance_de_calcul:   10,
			firewall:              5,
			stability:             5,
			vitesse_de_connection: 5,
			Equipements:           equipements,
			Inventory:             []string{"antivirus", "malware"},
			MaxHP:                 100,
			HP:                    100,
		}
	case '2':
		player = Character{
			Class:                 "SysAdmin",
			Level:                 1,
			BTC:                   0,
			Energie:               80,
			Spells:                []Spell{PatchUpdate, Defrag, Hotfix, ProxyBoost},
			Statut:                0,
			Max_Energie:           80,
			puissance_de_calcul:   10,
			firewall:              5,
			stability:             5,
			vitesse_de_connection: 5,
			Equipements:           equipements,
			Inventory:             []string{"antivirus", "malware"},
			MaxHP:                 130,
			HP:                    130,
		}
	case '3':
		player = Character{
			Class:                 "Netrunner",
			Level:                 1,
			BTC:                   0,
			Statut:                0,
			Energie:               130,
			Max_Energie:           130,
			puissance_de_calcul:   10,
			firewall:              5,
			stability:             5,
			vitesse_de_connection: 5,
			Equipements:           equipements,
			Inventory:             []string{"antivirus", "malware"},
			MaxHP:                 80,
			HP:                    80,
			Spells:                []Spell{DDoS, TrojanHorse, ManInTheMiddle, RootAccess},
		}
	}
	return nil
}
