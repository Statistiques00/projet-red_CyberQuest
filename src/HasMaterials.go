package main

func HasMaterials(player *Character, required []string) bool {
	for _, mat := range required {
		for _, item := range player.Inventory {
			if item == mat {
				return true
			}
		}
	}
	if player.BTC < 5 {
		return false
	}
	return false
}
