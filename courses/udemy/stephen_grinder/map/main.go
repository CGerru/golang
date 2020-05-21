package main

import "fmt"

func main() {
	// This is how maps are declared map[key type]value type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "hexcode",
		"white": "#ffffff",
	}

	// make built-in function can be used too
	dogs := make(map[string]string)

	// Adding elements to map
	dogs["dalmatian"] = "rocky"

	// Removing elements from map
	delete(colors, "green")

	// Time to iterate over maps
	printMap(colors)
	fmt.Println(colors)
}

// printMap iterates over a [string]string map and print its elements
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
