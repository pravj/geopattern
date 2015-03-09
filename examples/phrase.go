package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's SVG string for a phrase argument
func main() {
	args := map[string]string{"phrase": "O"}
	gp := geoPattern.Generate(args)
	fmt.Println(gp)
}
