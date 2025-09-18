package main

func RemoveMaterials(inventory *[]string, required []string) {
	for _, mat := range required {
		for i, item := range *inventory {
			if item == mat {
				*inventory = append((*inventory)[:i], (*inventory)[i+1:]...)
			}
		}
	}
}