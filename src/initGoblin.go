package main

// Fonction pour initialiser un Gobelin d'entraînement
func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entrainement",
		HP:         40,
		MaxHP:      40,
		Attack:     5,
		Experience: 10,
		BTC:        5,
	}
}
