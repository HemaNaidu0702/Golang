package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	printmap(colors)
}

func printmap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)
	}
}
