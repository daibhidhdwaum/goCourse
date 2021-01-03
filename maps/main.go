package main

import "fmt"

// ! Each column in a map must be of the same type. So, all keys must be of the same type and all values must be of the same type
func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// * maps only use bracket notation. dot notation is not available
	// colors["white"] = "#ffffff"
	// * this is how to delete
	// delete(colors, "white")

	printMap(colors)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
