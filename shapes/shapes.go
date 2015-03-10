// Package shapes implements some geometric shapes to be used in project
package shapes

import (
	"fmt"
	"github.com/pravj/geopattern/svg"
	"github.com/pravj/geopattern/utils"
	"math"
)

// BuildOctagon returns string representing an octagon shape
func BuildOctagon(squareSize float64) string {
	s := squareSize
	c := 0.33 * s

	return fmt.Sprintf("%v,0,%v,0,%v,%v,%v,%v,%v,%v,%v,%v,0,%v,0,%v,%v,0", c, s-c, s, c, s, s-c, s-c, s, c, s, s-c, c, c)
}

// BuildTriangle returns string representing a triangle shape
func BuildTriangle(sideLength, height float64) string {
	halfWidth := sideLength / 2

	return fmt.Sprintf("%v,0,%v,%v,0,%v,%v,0", halfWidth, sideLength, height, height, halfWidth)
}

// BuildDiamond returns string representing a diamond shape
func BuildDiamond(width, height float64) string {
	return fmt.Sprintf("%v,0,%v,%v,%v,%v,0,%v", width/2, width, height/2, width/2, height, height/2)
}

// BuildRightTriangle returns string representing a right angle triangle shape
func BuildRightTriangle(sideLength float64) string {
	return fmt.Sprintf("0,0,%v,%v,0,%v,0,0", sideLength, sideLength, sideLength)
}

// BuildRotatedTriangle returns string representing a rotated triangle shape
func BuildRotatedTriangle(sideLength, width float64) string {
	halfHeight := sideLength / 2

	return fmt.Sprintf("0,0,%v,%v,0,%v,0,0", width, halfHeight, sideLength)
}

// BuildHexagon returns string representing a hexagon shape
func BuildHexagon(sideLength float64) string {
	c := sideLength
	a := c / 2
	b := math.Sin(60*math.Pi/180) * c

	return fmt.Sprintf("0,%v,%v,0,%v,0,%v,%v,%v,%v,%v,%v,0,%v", b, a, a+c, 2*c, b, a+c, 2*b, a, 2*b, b)
}

// BuildChevron returns string representing a chevron shape
func BuildChevron(width, height float64) [2]string {
	e := height * 0.66
	var elements [2]string

	elements[0] = fmt.Sprintf("<polyline points='0,0,%v,%v,%v,%v,0,%v,0,0' />", width/2, height-e, width/2, height, e)
	elements[1] = fmt.Sprintf("<polyline points='%v,%v,%v,0,%v,%v,%v,%v,%v,%v' />", width/2, height-e, width, width, e, width/2, height, width/2, height-e)

	return elements
}

// BuildPlus returns string representing an plus shape
func BuildPlus(squareSize float64) [2]string {
	var elements [2]string

	elements[0] = fmt.Sprintf("<rect x='%v' y='0' width='%v' height='%v' />", squareSize, squareSize, squareSize*3)
	elements[1] = fmt.Sprintf("<rect x='0' y='%v' width='%v' height='%v' />", squareSize, squareSize*3, squareSize)

	return elements
}

// DrawInnerMosaicTile returns string representing an inner mosaic tile shape
func DrawInnerMosaicTile(s *svg.SVG, x, y, triangleSize float64, values [2]float64) {
	triangle := BuildRightTriangle(triangleSize)
	opacity := utils.Opacity(values[0])
	fill := utils.FillColor(values[0])

	styles := make(map[string]interface{})
	styles["fill"] = fill
	styles["fill-opacity"] = opacity
	styles["stroke"] = utils.StrokeColor
	styles["stroke-opacity"] = utils.StrokeOpacity

	style := make(map[string]interface{})

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", x+triangleSize, y)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, -1)", x+triangleSize, y+triangleSize*2)
	s.Polyline(triangle, utils.Merge(styles, style))

	opacity = utils.Opacity(values[1])
	fill = utils.FillColor(values[1])

	styles["fill"] = fill
	styles["fill-opacity"] = opacity

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, -1)", x+triangleSize, y+triangleSize*2)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, 1)", x+triangleSize, y)
	s.Polyline(triangle, utils.Merge(styles, style))
}

// DrawOuterMosaicTile returns string representing an outer mosaic tile shape
func DrawOuterMosaicTile(s *svg.SVG, x, y, triangleSize, value float64) {
	opacity := utils.Opacity(value)
	fill := utils.FillColor(value)
	triangle := BuildRightTriangle(triangleSize)

	styles := make(map[string]interface{})
	styles["fill"] = fill
	styles["fill-opacity"] = opacity
	styles["stroke"] = utils.StrokeColor
	styles["stroke-opacity"] = utils.StrokeOpacity

	style := make(map[string]interface{})

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, -1)", x, y+triangleSize)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, -1)", x+triangleSize*2, y+triangleSize)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, 1)", x, y+triangleSize)
	s.Polyline(triangle, utils.Merge(styles, style))

	style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", x+triangleSize*2, y+triangleSize)
	s.Polyline(triangle, utils.Merge(styles, style))
}
