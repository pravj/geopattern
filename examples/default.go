package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's SVG string without any argument
func main() {
	args := map[string]string{}
	gp := geoPattern.Generate(args)
	fmt.Println(gp)
}
