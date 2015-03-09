// Package svg provide methods to work effortlessly with SVG.
package svg

import (
	"fmt"
	"reflect"
)

// SVG struct, SVG contains elementry attributes like
// svg string, width, height
type SVG struct {
	svgString    string
	width, height int
}

// SetWidth sets SVG object's width
func (s *SVG) SetWidth(w int) {
	s.width = w
}

// SetHeight sets SVG object's height
func (s *SVG) SetHeight(h int) {
	s.height = h
}

// header returns string representing SVG object's header(staring) part
func (s *SVG) header() string {
	return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' width='%v' height='%v'>", s.width, s.height)
}

// footer returns string representing SVG object's footer(ending) part
func (s *SVG) footer() string {
	return "</svg>"
}

// Str returns string representing whole SVG object
func (s *SVG) Str() string {
	return s.header() + s.svgString + s.footer()
}

// Rect adds a rectangle element to SVG object
func (s *SVG) Rect(x, y, w, h interface{}, args map[string]interface{}) {
	rectStr := fmt.Sprintf("<rect x='%v' y='%v' width='%v' height='%v' %s />", x, y, w, h, s.WriteArgs(args))
	s.svgString += rectStr
}

// Circle adds a circle element to SVG object
func (s *SVG) Circle(cx, cy, r interface{}, args map[string]interface{}) {
	circleStr := fmt.Sprintf("<circle cx='%v' cy='%v' r='%v' %s />", cx, cy, r, s.WriteArgs(args))
	s.svgString += circleStr
}

// Path adds a path element to SVG object
func (s *SVG) Path(str string, args map[string]interface{}) {
	pathStr := fmt.Sprintf("<path d='%s' %s />", str, s.WriteArgs(args))
	s.svgString += pathStr
}

// Polyline adds a polyline element to SVG object
func (s *SVG) Polyline(str string, args map[string]interface{}) {
	polylineStr := fmt.Sprintf("<polyline points='%s' %s />", str, s.WriteArgs(args))
	s.svgString += polylineStr
}

// Group adds a group element to SVG object.
//
// It groups optionally provided elements together.
func (s *SVG) Group(elements [2]string, args map[string]interface{}) {
	s.svgString += fmt.Sprintf("<g %s>", s.WriteArgs(args))
	s.svgString += elements[0] + elements[1]
	s.svgString += "</g>"
}

// WriteArgs adds additional attributes to a SVG elements.
//
// It parses provides 'map' arguments to add attributes to SVG element.
func (s *SVG) WriteArgs(args map[string]interface{}) string {
	str := ""

	for k, v := range args {
		objType := fmt.Sprintf("%s", reflect.TypeOf(v))

		switch objType {
		case "string":
			str += fmt.Sprintf("%s='%s' ", k, v)
		case "int":
			str += fmt.Sprintf("%s='%v' ", k, v)
		case "float64":
			str += fmt.Sprintf("%s='%v' ", k, v)
		default:
			{
				str += fmt.Sprintf("%s='", k)
				for K, V := range v.(map[string]string) {
					str += fmt.Sprintf("%s:%s;", K, V)
				}
				str += "' "
			}
		}
	}

	return str
}
