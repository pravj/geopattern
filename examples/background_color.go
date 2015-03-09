package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's SVG string with a specific background color
func main() {
	args := map[string]string{"color": "#f9b"}
	gp := geoPattern.Generate(args)
	fmt.Println(gp)
}
