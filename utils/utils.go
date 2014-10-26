// Package utils provide utility functions that helps
// development across the project.
package utils

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

// constants representing stroke, fill's opacity and color,
// base color, minimum and maximum opacity value
const (
	BASE_COLOR       string  = "#9e2c7b"
	STROKE_COLOR     string  = "#000"
	STROKE_OPACITY   float64 = 0.02
	FILL_COLOR_DARK  string  = "#222"
	FILL_COLOR_LIGHT string  = "#ddd"
	OPACITY_MIN      float64 = 0.02
	OPACITY_MAX      float64 = 0.15
)

// Hash returns SHA-1 encryption of a string
func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

// Map returns respective value of a number from a range to different range
func Map(value, a_min, a_max, b_min, b_max float64) float64 {
	a_range := a_max - a_min
	b_range := b_max - b_min

	return (b_max - (a_max-value)*(b_range/a_range))
}

// Hex_val returns decimal representation of a substring of a hexa decimal string
func Hex_val(str string, index, length int) float64 {
	hex_str := str[index : index+length]

	hex_val, err := strconv.ParseInt(hex_str, 16, 0)
	if err != nil {
		panic(err)
	}

	return float64(hex_val)
}

// Merge merges two 'map' objects and returns the resultant object
func Merge(map_a map[string]interface{}, map_b map[string]interface{}) map[string]interface{} {
	for k, v := range map_a {
		map_b[k] = v
	}

	return map_b
}

// Opacity returns opacity value in a particular range
func Opacity(value float64) float64 {
	return Map(value, 0, 15, OPACITY_MIN, OPACITY_MAX)
}

// Fill_color returns string to be used for fill color
func Fill_color(value float64) string {
	if int(value)%2 == 0 {
		return FILL_COLOR_LIGHT
	} else {
		return FILL_COLOR_DARK
	}
}
