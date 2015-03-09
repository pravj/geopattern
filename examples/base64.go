package main

import (
	"fmt"
	"github.com/pravj/geo_pattern"
)

// Prints pattern's Base64 encoded string
func main() {
	args := map[string]string{}
	gp := geo_pattern.Base64String(args)
	fmt.Println(gp)
}
