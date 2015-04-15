package main

import (
	"fmt"
	"github.com/Thesandlord/geopattern"
)

// Prints pattern's SVG string with a specific background color
func main() {
	geopattern.SetTime(1234)
	args := geopattern.Pattern{
		Color: "#f9b",
	}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
}
