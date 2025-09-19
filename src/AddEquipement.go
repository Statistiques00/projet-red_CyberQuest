package main

func AddEquipmentToPlayer(player *Character, itemName string) {
	// Implementation for adding equipment to player goes here
}

// Fonction pour équiper un casque
func AddCasque(p *Character, c Casque) {
	p.Equipements.casque = c
}

// Fonction pour équiper une armure
func AddArmure(p *Character, a Armure) {
	p.Equipements.armure = a
}

// Fonction pour équiper une épée
func AddEpee(p *Character, e Armes) {
	p.Equipements.armes = e
}