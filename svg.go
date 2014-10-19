// what about changing package to svg??
package main

import "fmt"

type SVG struct {
    svg_string string
    width, height int
}

func (s *SVG) set_width(w int) {
    s.width = w
}

func (s *SVG) set_height(h int) {
    s.height = h
}

func (s *SVG) header() string {
    return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' width='%v' height='%v'>", s.width, s.height)
}

func (s *SVG) footer() string {
    return "</svg>"
}

func (s *SVG) to_s() string {
    return s.header() + s.svg_string + s.footer()
}

func (s *SVG) rect(x, y, w, h int) {
    rect_str := fmt.Sprintf("<rect x='%v' y='%v' width='%v' height='%v' />", x, y, w, h)
    s.svg_string += rect_str
}

func (s *SVG) circle(cx, cy, r int) {
    circle_str := fmt.Sprintf("<circle cx='%v' cy='%v' r='%v' />", cx, cy, r)
    s.svg_string += circle_str
}

func (s *SVG) path(str string) {
    path_str := fmt.Sprintf("<path d='%s' />", str)
    s.svg_string += path_str
}

func (s *SVG) polyline(str string) {
    polyline_str := fmt.Sprintf("<polyline points='%s' />", str)
    s.svg_string += polyline_str
}

func main() {
    svg := new(SVG)

    svg.set_height(100)
    svg.set_width(100)

    svg.rect(1,2,4,5)
    svg.circle(1,2,3)
    svg.path("path_string")
    svg.polyline("polyline_string")

    fmt.Println(svg.to_s())
}
