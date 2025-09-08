package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

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

// Declare player as a package-level variable
var player Character

func initCharacter() {
	// initialiser les personnages ici
	player = Character{
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		printIndented("=== MENU ===", 25)
		printIndented("Bienvenue dans le jeu !", 20)
		printIndented("1 - Option 1", 23)
		printIndented("2 - Option 2", 23)
		printIndented("3 - Quitter", 23)
		fmt.Print("Choix : ")
		scanner.Scan()
		choix := scanner.Text()
		switch choix {
		case "1":
			fmt.Println("Tu as choisi l'option 1")
		case "2":
			fmt.Println("Tu as choisi l'option 2")
		case "3":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
		scanner.Scan()
	}
}
// Affiche le texte centré dans le terminal
func printCentered(text string) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= len(text) {
		fmt.Println(text)
		return
	}
	padding := (width - len(text)) / 2
	fmt.Printf("%s%s\n", spaces(padding), text)
}

func spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

func characterCreation() {
	// creer un personnage ici
	// ne pas oublier les classes
	// utiliser des switch cases
	println("Quel est votre nom ?")
	player.Name = scan(&player.Name)
	println("Choisissez une classe parmi Overclockeur, Sysadmin et Netrunner")
	switch scan(&player.Class) {
	case "Overclockeur":
		// set les stats
	case "Sysadmin":
		// set les stats
	case "Netrunner":
		// set les stats
	}
}

func displayInfo() {
	// afficher les infos des personnages ici
	println(player.Name)
	println(player.Class)
	println(player.Level)
	println(player.Gold)
	println(player.Equipment)
	println(player.Mana)
	println(player.Spells)
	println(player.MaxHP)
	println(player.HP)
	println(player.Defense)

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

func upgradeInventorySlot() {
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

func clearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func printCenteredLines(lines []string) {
    for _, line := range lines {
        printCentered(line)
    }
}

func printIndented(text string, indent int) {
    fmt.Printf("%s%s\n", spaces(indent), text)
}