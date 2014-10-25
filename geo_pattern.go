package geo_pattern

import (
	"encoding/base64"
	"fmt"
	"github.com/pravj/geo_pattern/pattern"
)

func Generate(phrase string) string {
	p := pattern.New(phrase)

	return p.Svg_str()
}

func Base64_string(phrase string) string {
	svg_str := Generate(phrase)
	base64_str := base64.StdEncoding.EncodeToString([]byte(svg_str))

	return base64_str
}

func Uri_image(phrase string) string {
	base64_str := Base64_string(phrase)

	return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64_str)
}
