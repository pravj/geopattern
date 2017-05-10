package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's uri image string
func main() {
	args := geopattern.Pattern{}
	gp := geopattern.URIimage(args)
	fmt.Println(gp)
}
