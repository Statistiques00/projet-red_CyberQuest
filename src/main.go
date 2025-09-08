package main

// importer les packages necessaires ici
type Equipment struct {
	// definir les equipements ici
	Head   string
	Toros  string
	Legs   string
	Feet   string
	Hands  string
	Weapon string
	Shield string
}

type Character struct {
	// definir les attributs des personnages ici
	Name      string
	Class     string
	Level     int
	Gold      int
	Equipment Equipment
	Mana      int
	Spells    []string
	MaxHP     int
	HP        int
	Defense   int
}

type Monster struct {
	// definir les attributs des monstres ici
	Name       string
	MaxHP      int
	HP         int
	Attack     int
	Defense    int
	Speed      int
	Experience int
	Loot       []string
}

func initCharacter() {
	// initialiser les personnages ici
	perso := Character{
		Name:  "Hero",
		Class: "Warrior",
		Level: 1,
		Gold:  0,
		Equipment: Equipment{
			Head:   "Iron Helmet",
			Toros:  "Iron Armor",
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

func main() {
	initCharacter()
}

func characterCreation() {
	// creer un personnage ici
	// ne pas oublier les classes
	// utiliser des switch case
}

func displayInfo() {
	// afficher les infos des personnages ici

}

func accesInventory() {
	// acceder a l'inventaire ici
}

func addInventory() {
	// ajouter un objet a l'inventaire ici
	// gerer la capacite maximale
}

func removeInventory() {
	// retirer un objet de l'inventaire ici
}

func upgradeIventorySlot() {
	// augmenter la capacite maximale de l'inventaire ici
}

func takePot() {
	// prendre une potion ici
}

// switch case menu

func poisonPot() {
	// appliquer l'effet de la potion empoisonnee ici
}

func isDead() {
	// verifier si le personnage est mort ici
}
