package main

import "fmt"

func main() {
	// 1. var colors map[string]string
	/// 2. colors := make(map[string]string) // 2.
	colors := map[string]string{ // 3.
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors["white"] = "#fffff"

	// delete elements from map
	//delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
