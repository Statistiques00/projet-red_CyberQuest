package main

func InitCharacter() {
	// initialiser les personnages ici
	player = Character{
		Name:  "Hero",
		Class: "Warrior",
		Level: 1,
		Gold:  0,
		Equipment: Equipment{
			Head:   "Iron Helmet",
			Torso:  "Iron Armor",
			Legs:   "Iron Leggings",
			Feet:   "Iron Boots",
			Hands:  "Iron Gloves",
			Weapon: "Iron Sword",
			Shield: "Wooden Shield",
		},
		Mana:    100,
		Spells:  []string{"Fireball", "Heal"},
		MaxHP:   100,
		HP:      100,
		Defense: 10,
	}

}