package geopattern

import (
	"fmt"
	"reflect"
	"sort"
)

// SVG struct, SVG contains elementry attributes like
// svg string, width, height
type SVG struct {
	svgString     string
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

	// Store the keys in slice in sorted order
	// This is necessary for testing (to ensure the same order of attributes), but it adds a performance hit
	// This needs to be a struct, NOT a map!
	var keys []string
	for k := range args {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		objType := fmt.Sprintf("%s", reflect.TypeOf(args[k]))

		switch objType {
		case "string":
			str += fmt.Sprintf("%s='%s' ", k, args[k])
		case "int":
			str += fmt.Sprintf("%s='%v' ", k, args[k])
		case "float64":
			str += fmt.Sprintf("%s='%v' ", k, args[k])
		default:
			{
				//Same kind of sorting as above
				str += fmt.Sprintf("%s='", k)
				var subkeys []string
				for e := range args[k].(map[string]string) {
					subkeys = append(subkeys, e)
				}
				sort.Strings(subkeys)
				for _, e := range subkeys {
					str += fmt.Sprintf("%s:%s;", e, args[k].(map[string]string)[e])
				}
				str += "' "
			}
		}
	}

	return str
}
