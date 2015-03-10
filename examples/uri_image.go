package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's uri image string
func main() {
	args := map[string]string{}
	gp := geopattern.URIimage(args)
	fmt.Println(gp)
}
