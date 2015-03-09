package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's uri image string
func main() {
	args := map[string]string{}
	gp := geo_pattern.URIimage(args)
	fmt.Println(gp)
}
