package main

import "fmt"

// Important: All values in map must be the same, same keys and same values types
func main(){
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	colors["yellow"] = "#fa1230a" // add a value

	printMap(colors);
}

func printMap(c map[string]string){
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex);
	}
}