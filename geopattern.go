// Package geopattern creates beautiful generative image patterns from a string.
package geopattern

import (
	"encoding/base64"
	"fmt"
	"time"
)

// Variable and Custom Function for Current Time so we can mock the current time and run unit tests
var now = func() time.Time { return time.Now().Local() }

func setTime(unix int) {
	now = func() time.Time { return time.Unix(int64(unix), 0) }
}

// Generate returns pattern's SVG string
func Generate(args Pattern) string {

	//If the phrase is empty, then generate a random phrase
	if args.Phrase == "" {
		args.Phrase = fmt.Sprintf("%s", now())
	}

	args.hash = Hash(args.Phrase)
	args.svg = new(SVG)

	return args.SvgStr()
}

// Base64String returns pattern's Base64 encoded string
func Base64String(args Pattern) string {
	svgStr := Generate(args)
	base64Str := base64.StdEncoding.EncodeToString([]byte(svgStr))

	return base64Str
}

// URIimage returns pattern's uri image string
func URIimage(args Pattern) string {
	base64Str := Base64String(args)

	return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64Str)
}
