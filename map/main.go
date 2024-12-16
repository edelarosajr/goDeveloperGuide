package main

import "fmt"

func main() {
	// creation method 3: var colors map[string]string

	// creation method 2: colors := make(map[string]string)

	// no dot access syntax

	// delete(colors, "white")
	// access: colors["white"] = "#FFFFFF"

	// creation method 1
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "hex value is", hex)
	}
}
