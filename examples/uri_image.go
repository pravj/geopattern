package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's uri image string
func main() {
	args := map[string]string{}
	gp := geoPattern.URIimage(args)
	fmt.Println(gp)
}
