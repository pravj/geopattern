package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's SVG string with a specific base background color
func main() {
	args := map[string]string{"baseColor": "#e2b"}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
}
