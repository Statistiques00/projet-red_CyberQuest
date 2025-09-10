package main

import "fmt"

func PrintCentered(text string) {
	width := 80
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Printf("%s%s\n", Spaces(padding), text)
}