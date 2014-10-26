package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's SVG string without any argument
func main() {
	args := map[string]string{}
	gp := geo_pattern.Generate(args)
	fmt.Println(gp)
}
