package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's SVG string for a specific pattern
func main() {
	args := map[string]string{"generator": "squares"}
	gp := geo_pattern.Generate(args)
	fmt.Println(gp)
}
