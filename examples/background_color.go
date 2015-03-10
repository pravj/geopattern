package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's SVG string with a specific background color
func main() {
	args := map[string]string{"color": "#f9b"}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
}
