// Package geo_pattern creates beautiful generative background image
// patterns from a string.
package geo_pattern

import (
	"encoding/base64"
	"fmt"
	"github.com/pravj/geo_pattern/pattern"
)

// Returns pattern's SVG string
func Generate(args map[string]string) string {
	p := pattern.New(args)

	return p.Svg_str()
}

// Returns pattern's Base64 encoded string
func Base64_string(args map[string]string) string {
	svg_str := Generate(args)
	base64_str := base64.StdEncoding.EncodeToString([]byte(svg_str))

	return base64_str
}

// Returns pattern's uri image string
func Uri_image(args map[string]string) string {
	base64_str := Base64_string(args)

	return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64_str)
}
