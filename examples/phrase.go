package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's SVG string for a phrase argument
func main() {
	args := map[string]string{"phrase": "O"}
	gp := geo_pattern.Generate(args)
	fmt.Println(gp)
}
