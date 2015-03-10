package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's SVG string without any argument
func main() {
	args := map[string]string{}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
}
