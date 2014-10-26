// Package shapes implements some geometric shapes to be used in project
package shapes

import (
	"fmt"
	"github.com/pravj/geo_pattern/svg"
	"github.com/pravj/geo_pattern/utils"
	"math"
)

// Build_octagon returns string representing an octagon shape
func Build_octagon(square_size float64) string {
	s := square_size
	c := 0.33 * s

	return fmt.Sprintf("%v,0,%v,0,%v,%v,%v,%v,%v,%v,%v,%v,0,%v,0,%v,%v,0", c, s-c, s, c, s, s-c, s-c, s, c, s, s-c, c, c)
}

// Build_triangle returns string representing a triangle shape
func Build_triangle(side_length, height float64) string {
	half_width := side_length / 2

	return fmt.Sprintf("%v,0,%v,%v,0,%v,%v,0", half_width, side_length, height, height, half_width)
}

// Build_diamond returns string representing a diamond shape
func Build_diamond(width, height float64) string {
	return fmt.Sprintf("%v,0,%v,%v,%v,%v,0,%v", width/2, width, height/2, width/2, height, height/2)
}

// Build_right_triangle returns string representing a right angle triangle shape
func Build_right_triangle(side_length float64) string {
	return fmt.Sprintf("0,0,%v,%v,0,%v,0,0", side_length, side_length, side_length)
}

// Build_rotated_triangle returns string representing a rotated triangle shape
func Build_rotated_triangle(side_length, width float64) string {
	half_height := side_length / 2

	return fmt.Sprintf("0,0,%v,%v,0,%v,0,0", width, half_height, side_length)
}

// Build_hexagon returns string representing a hexagon shape
func Build_hexagon(side_length float64) string {
	c := side_length
	a := c / 2
	b := math.Sin(60*math.Pi/180) * c

	return fmt.Sprintf("0,%v,%v,0,%v,0,%v,%v,%v,%v,%v,%v,0,%v", b, a, a+c, 2*c, b, a+c, 2*b, a, 2*b, b)
}

// Build_chevron returns string representing a chevron shape
func Build_chevron(width, height float64) [2]string {
	e := height * 0.66
	var elements [2]string

	elements[0] = fmt.Sprintf("<polyline points='0,0,%v,%v,%v,%v,0,%v,0,0' />", width/2, height-e, width/2, height, e)
	elements[1] = fmt.Sprintf("<polyline points='%v,%v,%v,0,%v,%v,%v,%v,%v,%v' />", width/2, height-e, width, width, e, width/2, height, width/2, height-e)

	return elements
}

// Build_plus returns string representing an plus shape
func Build_plus(square_size float64) [2]string {
	var elements [2]string

	elements[0] = fmt.Sprintf("<rect x='%v' y='0' width='%v' height='%v' />", square_size, square_size, square_size*3)
	elements[1] = fmt.Sprintf("<rect x='0' y='%v' width='%v' height='%v' />", square_size, square_size*3, square_size)

	return elements
}

// Draw_inner_mosaic_tile returns string representing an inner mosaic tile shape
func Draw_inner_mosaic_tile(s *svg.SVG, x, y, triangle_size float64, values [2]float64) {
	triangle := Build_right_triangle(triangle_size)
	opacity := utils.Opacity(values[0])
	fill := utils.Fill_color(values[0])

	styles := make(map[string]interface{})
	styles["fill"] = fill
	styles["fill-opacity"] = opacity
	styles["stroke"] = utils.STROKE_COLOR
	styles["stroke-opacity"] = utils.STROKE_OPACITY

	style := make(map[string]interface{})

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", x+triangle_size, y)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, -1)", x+triangle_size, y+triangle_size*2)
	s.Polyline(triangle, utils.Merge(styles, style))

	opacity = utils.Opacity(values[1])
	fill = utils.Fill_color(values[1])

	styles["fill"] = fill
	styles["fill-opacity"] = opacity

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, -1)", x+triangle_size, y+triangle_size*2)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, 1)", x+triangle_size, y)
	s.Polyline(triangle, utils.Merge(styles, style))
}

// Draw_outer_mosaic_tile returns string representing an outer mosaic tile shape
func Draw_outer_mosaic_tile(s *svg.SVG, x, y, triangle_size, value float64) {
	opacity := utils.Opacity(value)
	fill := utils.Fill_color(value)
	triangle := Build_right_triangle(triangle_size)

	styles := make(map[string]interface{})
	styles["fill"] = fill
	styles["fill-opacity"] = opacity
	styles["stroke"] = utils.STROKE_COLOR
	styles["stroke-opacity"] = utils.STROKE_OPACITY

	style := make(map[string]interface{})

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, -1)", x, y+triangle_size)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, -1)", x+triangle_size*2, y+triangle_size)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, 1)", x, y+triangle_size)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", x+triangle_size*2, y+triangle_size)
	s.Polyline(triangle, utils.Merge(styles, style))
}
