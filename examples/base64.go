package main

import (
	"fmt"
	"github.com/pravj/geoPattern"
)

// Prints pattern's Base64 encoded string
func main() {
	args := map[string]string{}
	gp := geoPattern.Base64String(args)
	fmt.Println(gp)
}
