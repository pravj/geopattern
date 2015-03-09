package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's SVG string with a specific base background color
func main() {
	args := map[string]string{"baseColor": "#e2b"}
	gp := geoPattern.Generate(args)
	fmt.Println(gp)
}
