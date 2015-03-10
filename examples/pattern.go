package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's SVG string for a specific pattern
func main() {
	args := map[string]string{"generator": "squares"}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
}
