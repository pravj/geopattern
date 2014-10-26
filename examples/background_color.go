package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's SVG string with a specific background color
func main() {
	args := map[string]string{"color": "#f9b"}
	gp := geo_pattern.Generate(args)
	fmt.Println(gp)
}
