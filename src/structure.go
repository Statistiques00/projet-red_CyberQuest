package main

type Casque struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Armure struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Bottes struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Armes struct {
	nom                   string
	firewall              int
	puissance_de_calcul   int
	stability             int
	vitesse_de_connection int
	valeur                int
}

type Equipements struct {
	casque Casque
	armure Armure
	bottes Bottes
	armes  Armes
}

type Spell struct {
	nom         string
	cost        int
	description string
	degats      int
}
type Character struct {
	Name                  string
	Class                 string
	Level                 int
	BTC                   int
	Statut                int
	puissance_de_calcul   int //attaque
	firewall              int //defense
	stability             int //tx critique
	vitesse_de_connection int //initiative
	Energie               int //mana
	Max_Energie           int //max mana
	MaxHP                 int
	HP                    int
	Inventory             []string
	Equipements           Equipements
	Spells                []Spell
	MaxInventory          int
	valeur                int
}

type Monster struct {
	// definir les attributs des monstres ici
	Name       string
	MaxHP      int
	HP         int
	Poisoned   int
	Attack     int
	Defense    int
	Speed      int
	Experience int
	LootTable  []LootItem
	BTC        int
}

type LootItem struct {
	Name       string
	DropChance float64
}

var player Character

var Spells1 = Spell{
	nom:         "Attaque basique",
	cost:        0,
	description: "Une attaque simple mais efficace.",
	degats:      2,
}

var Spells2 = Spell{
	nom:         "Attaque puissante",
	cost:        10,
	description: "Une attaque qui inflige plus de dégâts.",
	degats:      5,
}

var equipements = Equipements{
	casque: Casque{
		nom:                   "Casque de base",
		firewall:              5,
		puissance_de_calcul:   0,
		stability:             0,
		vitesse_de_connection: 0,
		valeur:                10,
	},
	armure: Armure{
		nom:                   "Armure de base",
		firewall:              15,
		puissance_de_calcul:   0,
		stability:             0,
		vitesse_de_connection: -1,
		valeur:                30,
	},
	bottes: Bottes{
		nom:                   "Bottes de base",
		firewall:              5,
		puissance_de_calcul:   0,
		stability:             0,
		vitesse_de_connection: 2,
		valeur:                20,
	},
	armes: Armes{
		nom:                   "Epée de base",
		firewall:              0,
		puissance_de_calcul:   10,
		stability:             5,
		vitesse_de_connection: 0,
		valeur:                25,
	},
}

var Overclock = Spell{
	nom:         "Overclock",
	cost:        20,
	description: "Inflige des dégâts et augmente la vitesse de connexion.",
	degats:      8,
}

var CodeInjection = Spell{
	nom:         "Code Injection",
	cost:        15,
	description: "Inflige des dégâts et peut empoisonner l'ennemi.",
	degats:      7,
}

var Pingflood = Spell{
	nom:         "Pingflood",
	cost:        25,
	description: "Inflige des dégâts à tous les ennemis.",
	degats:      5,
}

var DataSurge = Spell{
	nom:         "Data Surge",
	cost:        30,
	description: "Inflige de lourds dégâts à un ennemi.",
	degats:      12,
}

var KernelPanic = Spell{
	nom:         "Kernel Panic",
	cost:        50,
	description: "Inflige des dégâts massifs à un ennemi.",
	degats:      20,
}

var PatchUpdate = Spell{
	nom:         "Patch Update",
	cost:        20,
	description: "Inflige des dégâts et peut réparer les erreurs.",
	degats:      5,
}

var Defrag = Spell{
	nom:         "Defrag",
	cost:        15,
	description: "Augmente temporairement la défense.",
	degats:      0,
}

var Hotfix = Spell{
	nom:         "Hotfix",
	cost:        3,
	description: "Soigne une grande partie des points de vie.",
	degats:      5,
}

var ProxyBoost = Spell{
	nom:         "Proxy Boost",
	cost:        10,
	description: "Augmente temporairement la vitesse de connexion.",
	degats:      7,
}

var DDoS = Spell{
	nom:         "DDoS",
	cost:        55,
	description: "Inflige des dégâts à tous les ennemis.",
	degats:      15,
}

var TrojanHorse = Spell{
	nom:         "Trojan Horse",
	cost:        40,
	description: "Inflige des dégâts et peut empoisonner l'ennemi.",
	degats:      10,
}

var ManInTheMiddle = Spell{
	nom:         "Man In The Middle",
	cost:        30,
	description: "Inflige des dégâts et vole des BTC à l'ennemi.",
	degats:      8,
}

var RootAccess = Spell{
	nom:         "Root Access",
	cost:        70,
	description: "Inflige des dégâts massifs à un ennemi.",
	degats:      25,
}
