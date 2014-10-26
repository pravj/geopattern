package geo_pattern

import (
	"encoding/base64"
	"fmt"
	"github.com/pravj/geo_pattern/pattern"
)

func Generate(args map[string]string) string {
	p := pattern.New(args)

	return p.Svg_str()
}

func Base64_string(args map[string]string) string {
	svg_str := Generate(args)
	base64_str := base64.StdEncoding.EncodeToString([]byte(svg_str))

	return base64_str
}

func Uri_image(args map[string]string) string {
	base64_str := Base64_string(args)

	return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64_str)
}
