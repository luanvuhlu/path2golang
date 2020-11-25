package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)
	colors["white"] = "ffffff"
	colors["black"] = "000000"
	colors["yellow"] = "ffff00 "
	colors2 := map[string]string{
		"white": "ffffff",
		"black": "000000",
	}
	delete(colors2, "black")
	fmt.Println(colors2)
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
