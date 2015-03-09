package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's SVG string for a specific pattern
func main() {
	args := map[string]string{"generator": "squares"}
	gp := geoPattern.Generate(args)
	fmt.Println(gp)
}
