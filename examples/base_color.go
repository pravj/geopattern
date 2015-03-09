package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's SVG string with a specific base background color
func main() {
	args := map[string]string{"baseColor": "#e2b"}
	gp := geo_pattern.Generate(args)
	fmt.Println(gp)
}
