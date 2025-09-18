package main

import "fmt"

func PoisonStatut(Monster *Monster) {

	if Monster.Poisoned == 0 {
		return
	}
	if Monster.Poisoned > 0 {
		Monster.HP -= 5
		Monster.Poisoned--
		if Monster.HP < 0 {
			Monster.HP = 0
		}
		fmt.Printf("| %s subit 5 points de dégâts de poison. |\n", Monster.Name)
		fmt.Printf("| PV: %d/%d                             |\n", Monster.HP, Monster.MaxHP)
		fmt.Println("+=========================================+")
	}

}
