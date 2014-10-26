package utils

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

const (
	BASE_COLOR       string  = "#933c3c"
	STROKE_COLOR     string  = "#000"
	STROKE_OPACITY   float64 = 0.02
	FILL_COLOR_DARK  string  = "#222"
	FILL_COLOR_LIGHT string  = "#ddd"
	OPACITY_MIN      float64 = 0.02
	OPACITY_MAX      float64 = 0.15
)

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func Map(value, a_min, a_max, b_min, b_max float64) float64 {
	a_range := a_max - a_min
	b_range := b_max - b_min

	return (b_max - (a_max-value)*(b_range/a_range))
}

func Hex_val(str string, index, length int) float64 {
	hex_str := str[index : index+length]

	hex_val, err := strconv.ParseInt(hex_str, 16, 0)
	if err != nil {
		panic(err)
	}

	return float64(hex_val)
}

func Merge(map_a map[string]interface{}, map_b map[string]interface{}) map[string]interface{} {
	for k, v := range map_a {
		map_b[k] = v
	}

	return map_b
}

func Opacity(value float64) float64 {
	return Map(value, 0, 15, OPACITY_MIN, OPACITY_MAX)
}

func Fill_color(value float64) string {
	if int(value)%2 == 0 {
		return FILL_COLOR_LIGHT
	} else {
		return FILL_COLOR_DARK
	}
}
