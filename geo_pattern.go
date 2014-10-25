package geo_pattern

import (
	"github.com/pravj/geo_pattern/pattern"
)

func Generate(phrase string) string {
	p := pattern.New(phrase)
	return p.Start()
}
