package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's Base64 encoded string
func main() {
	args := geopattern.Pattern{}
	gp := geopattern.Base64String(args)
	fmt.Println(gp)
}
