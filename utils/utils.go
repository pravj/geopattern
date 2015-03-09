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
	BaseColor       string  = "#9e2c7b"
	StrokeColor     string  = "#000"
	StrokeOpacity   float64 = 0.02
	FillColorDark  string  = "#222"
	FillColorLight string  = "#ddd"
	OpacityMin      float64 = 0.02
	OpacityMax      float64 = 0.15
)

// Hash returns SHA-1 encryption of a string
func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

// Map returns respective value of a number from a range to different range
func Map(value, aMin, aMax, bMin, bMax float64) float64 {
	aRange := aMax - aMin
	bRange := bMax - bMin

	return (bMax - (aMax-value)*(bRange/aRange))
}

// HexVal returns decimal representation of a substring of a hexa decimal string
func HexVal(str string, index, length int) float64 {
	hexStr := str[index : index+length]

	hexVal, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		panic(err)
	}

	return float64(hexVal)
}

// Merge merges two 'map' objects and returns the resultant object
func Merge(mapA map[string]interface{}, mapB map[string]interface{}) map[string]interface{} {
	for k, v := range mapA {
		mapB[k] = v
	}

	return mapB
}

// Opacity returns opacity value in a particular range
func Opacity(value float64) float64 {
	return Map(value, 0, 15, OpacityMin, OpacityMax)
}

// FillColor returns string to be used for fill color
func FillColor(value float64) string {
	if int(value)%2 == 0 {
		return FillColorLight
	}
	return FillColorDark
}
