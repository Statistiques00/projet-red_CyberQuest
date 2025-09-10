package main

func IsDead(p *Character) {
	if p.HP <= 0 {
		// Le joueur meurt et ressuscite avec 50% de ses PV max
		p.HP = p.MaxHP / 2
	}
}
