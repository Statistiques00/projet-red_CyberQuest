package main

import "fmt"

func PrintIndented(text string, indent int) {
	fmt.Printf("%s%s\n", Spaces(indent), text)
}