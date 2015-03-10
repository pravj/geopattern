// Package pattern implements methods to generate a new pattern
// background for the pattern, type of pattern and particular SVG structure
// for that pattern
package pattern

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/pravj/geopattern/shapes"
	"github.com/pravj/geopattern/svg"
	"github.com/pravj/geopattern/utils"
	"math"
	"time"
)

// All available geo patterns
var PATTERNS = [16]string{
	"chevrons",
	"concentric-circles",
	"diamonds",
	"hexagons",
	"mosaic-squares",
	"nested-squares",
	"octagons",
	"overlapping-circles",
	"overlapping-rings",
	"plaid",
	"plus-signs",
	"sine-waves",
	"squares",
	"tessellation",
	"triangles",
	"xes",
}

// Pattern struct that contains attributes like base color, background color
// pattern type, phrase for the pattern and SHA-1 hash of phrase
type Pattern struct {
	BaseColor string
	Color      string
	Generator  string
	Hash       string
	Svg        *svg.SVG
}

// New parses the arguments and returns an instance of Pattern struct that has linked
// methods to work further things
func New(args map[string]string) *Pattern {
	var phrase, generator, color, baseColor string

	phrase = fmt.Sprintf("%s", time.Now().Local())

	if args["phrase"] != "" {
		phrase = args["phrase"]
	}

	if args["generator"] != "" {
		generator = args["generator"]
	}

	if args["color"] != "" {
		color = args["color"]
	}

	if args["baseColor"] != "" {
		baseColor = args["baseColor"]
	}

	return &Pattern{BaseColor: baseColor, Color: color, Generator: generator, Hash: utils.Hash(phrase), Svg: new(svg.SVG)}
}

// SvgStr returns string representing pattern's SVG string
func (p *Pattern) SvgStr() string {
	p.generateBackground()
	p.genaratePattern()

	return p.Svg.Str()
}

// generateBackground decides on background color for the pattern.
//
// It uses 'color' or 'baseColor' arguments for this task.
func (p *Pattern) generateBackground() {
	var rgb, color colorful.Color

	if p.Color != "" {
		rgb, _ = colorful.Hex(p.Color)
	} else {
		hueOffset := utils.Map(utils.HexVal(p.Hash, 14, 3), 0, 4095, 0, 359)
		satOffset := utils.Map(utils.HexVal(p.Hash, 17, 1), 0, 15, -1, 1)

		if p.BaseColor == "" {
			color, _ = colorful.Hex(utils.BaseColor)
		} else {
			color, _ = colorful.Hex(p.BaseColor)
		}

		h, c, l := color.Hcl()

		h -= hueOffset
		if satOffset >= 0 {
			c += float64(satOffset)
		} else {
			c -= float64(satOffset)
		}

		rgb = colorful.Color{h, c, l}
	}

	r, g, b := int(rgb.R*105), int(rgb.G*105), int(rgb.B*150)

	args := make(map[string]interface{})
	args["fill"] = fmt.Sprintf("rgb(%v, %v, %v)", r, g, b)

	p.Svg.Rect(0, 0, "100%", "100%", args)
}

// isPattern decides whether a pattern is a valid one or not
func isPattern(generator string) bool {
	for _, ptn := range PATTERNS {
		if ptn == generator {
			return true
		}
	}
	return false
}

// genaratePattern decides on type of pattern and build respective SVG object
func (p *Pattern) genaratePattern() {
	if p.Generator == "" {
		p.Generator = PATTERNS[int(utils.HexVal(p.Hash, 20, 1))]
	} else {
		if !isPattern(p.Generator) {
			panic("Error: the requested generator is invalid.")
		}
	}

	switch p.Generator {
	case "chevrons":
		p.geoChevrons()
	case "concentric-circles":
		p.geoConcentricCircles()
	case "diamonds":
		p.geoDiamonds()
	case "hexagons":
		p.geoHexagons()
	case "mosaic-squares":
		p.geoMosaicSquares()
	case "nested-squares":
		p.geoNestedSquares()
	case "octagons":
		p.geoOctagons()
	case "overlapping-circles":
		p.geoOverlappingCircles()
	case "overlapping-rings":
		p.geoOverlappingRings()
	case "plaid":
		p.geoPlaid()
	case "plus-signs":
		p.geoPlusSigns()
	case "sine-waves":
		p.geoSineWaves()
	case "squares":
		p.geoSquares()
	case "tessellation":
		p.geoTessellation()
	case "triangles":
		p.geoTriangles()
	case "xes":
		p.geoXes()
	}
}

// geoChevrons build the chevrons SVG pattern
func (p *Pattern) geoChevrons() {
	chevronWidth := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 30, 80)
	chevronHeight := chevronWidth
	chevron := shapes.BuildChevron(chevronWidth, chevronHeight)

	p.Svg.SetHeight(int(chevronHeight * 6 * 0.66))
	p.Svg.SetWidth(int(chevronWidth * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity
			styles["stroke-width"] = 1

			style := make(map[string]interface{})

			style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x)*chevronWidth, float64(y)*chevronHeight*0.66-chevronHeight/2)
			p.Svg.Group(chevron, utils.Merge(styles, style))

			if y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x)*chevronWidth, 6*chevronHeight*0.66-chevronHeight/2)
				p.Svg.Group(chevron, utils.Merge(styles, style))
			}

			i = i + 1
		}
	}
}

// geoConcentricCircles build the concentric_circles SVG pattern
func (p *Pattern) geoConcentricCircles() {
	scale := utils.HexVal(p.Hash, 0, 1)
	ringSize := utils.Map(scale, 0, 15, 10, 60)
	strokeWidth := ringSize / 5

	p.Svg.SetHeight(int((ringSize + strokeWidth) * 6))
	p.Svg.SetWidth(int((ringSize + strokeWidth) * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			cx := float64(x)*ringSize + float64(x)*strokeWidth + (ringSize+strokeWidth)/2
			cy := float64(y)*ringSize + float64(y)*strokeWidth + (ringSize+strokeWidth)/2

			styles := make(map[string]interface{})
			styles["fill"] = "none"
			styles["stroke"] = fill
			styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", strokeWidth)}

			p.Svg.Circle(cx, cy, ringSize/2, styles)

			val = utils.HexVal(p.Hash, 39-i, 1)
			opacity = utils.Opacity(val)
			fill = utils.FillColor(val)

			styles = make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity

			p.Svg.Circle(cx, cy, ringSize/4, styles)

			i = i + 1
		}
	}
}

// geoDiamonds build the diamonds SVG pattern
func (p *Pattern) geoDiamonds() {
	diamondWidth := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 10, 50)
	diamondHeight := utils.Map(utils.HexVal(p.Hash, 1, 1), 0, 15, 10, 50)
	diamond := shapes.BuildDiamond(diamondWidth, diamondHeight)

	p.Svg.SetHeight(int(diamondHeight * 3))
	p.Svg.SetWidth(int(diamondWidth * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity

			var dx float64
			if y%2 != 0 {
				dx = diamondWidth / 2
			}

			style := make(map[string]interface{})

			style["transform"] = fmt.Sprintf("translate(%v, %v)", dx + float64(x)*diamondWidth - diamondWidth/2, diamondHeight/2*float64(y) - diamondHeight/2)
			p.Svg.Polyline(diamond, utils.Merge(styles, style))

			if x == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v)", dx + 6*diamondWidth - diamondWidth/2, diamondHeight/2*float64(y) - diamondHeight/2)
				p.Svg.Polyline(diamond, utils.Merge(styles, style))
			}

			if y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v)", dx + float64(x)*diamondWidth - diamondWidth/2, diamondHeight/2*6 - diamondHeight/2)
				p.Svg.Polyline(diamond, utils.Merge(styles, style))
			}

			if x == 0 && y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v)", dx + 6*diamondWidth - diamondWidth/2, diamondHeight/2*6 - diamondHeight/2)
				p.Svg.Polyline(diamond, utils.Merge(styles, style))
			}

			i = i + 1
		}
	}
}

// geoHexagons build the hexagons SVG pattern
func (p *Pattern) geoHexagons() {
	scale := utils.HexVal(p.Hash, 0, 1)
	sideLength := utils.Map(scale, 0, 15, 8, 60)
	hexHeight := sideLength * math.Sqrt(3)
	hexWidth := sideLength * 2
	hex := shapes.BuildHexagon(sideLength)

	p.Svg.SetHeight(int(hexHeight * 6))
	p.Svg.SetWidth(int((hexWidth * 3) + (sideLength * 3)))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			var dy float64
			if x%2 == 0 {
				dy = float64(y) * hexHeight
			} else {
				dy = float64(y)*hexHeight + hexHeight/2
			}
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity

			style := make(map[string]interface{})

			style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x)*sideLength*1.5-hexWidth/2, dy-hexHeight/2)
			p.Svg.Polyline(hex, utils.Merge(styles, style))

			if x == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v)", 6*sideLength*1.5-hexWidth/2, dy-hexHeight/2)
				p.Svg.Polyline(hex, utils.Merge(styles, style))
			}

			if y == 0 {
				if x%2 == 0 {
					dy = 6 * hexHeight
				} else {
					dy = 6*hexHeight + hexHeight/2
				}
				style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x)*sideLength*1.5-hexWidth/2, dy-hexHeight/2)
				p.Svg.Polyline(hex, utils.Merge(styles, style))
			}

			if x == 0 && y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v)", 6*sideLength*1.5-hexWidth/2, 5*hexHeight+hexHeight/2)
				p.Svg.Polyline(hex, utils.Merge(styles, style))
			}

			i = i + 1
		}
	}
}

// geoMosaicSquares build the mosaic_squares SVG pattern
func (p *Pattern) geoMosaicSquares() {
	triangleSize := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 15, 50)

	p.Svg.SetHeight(int(triangleSize * 8))
	p.Svg.SetWidth(int(triangleSize * 8))

	i := 0
	for y := 0; y <= 3; y++ {
		for x := 0; x <= 3; x++ {

			values := [2]float64{utils.HexVal(p.Hash, i, 1), utils.HexVal(p.Hash, i+1, 1)}

			if x%2 == 0 {
				if y%2 == 0 {
					shapes.DrawOuterMosaicTile(p.Svg, float64(x)*triangleSize*2, float64(y)*triangleSize*2, triangleSize, utils.HexVal(p.Hash, i, 1))
				} else {
					shapes.DrawInnerMosaicTile(p.Svg, float64(x)*triangleSize*2, float64(y)*triangleSize*2, triangleSize, values)
				}
			} else {
				if y%2 == 0 {
					shapes.DrawInnerMosaicTile(p.Svg, float64(x)*triangleSize*2, float64(y)*triangleSize*2, triangleSize, values)
				} else {
					shapes.DrawOuterMosaicTile(p.Svg, float64(x)*triangleSize*2, float64(y)*triangleSize*2, triangleSize, utils.HexVal(p.Hash, i, 1))
				}
			}

			i = i + 1
		}
	}

}

// geoNestedSquares build the nested_squares SVG pattern
func (p *Pattern) geoNestedSquares() {
	blockSize := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 4, 12)
	squareSize := blockSize * 7

	p.Svg.SetHeight(int((squareSize+blockSize)*6 + blockSize*6))
	p.Svg.SetWidth(int((squareSize+blockSize)*6 + blockSize*6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = "none"
			styles["stroke"] = fill
			styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", blockSize)}

			p.Svg.Rect(float64(x)*squareSize+float64(x)*blockSize*2+blockSize/2, float64(y)*squareSize+float64(y)*blockSize*2+blockSize/2, squareSize, squareSize, styles)

			val = utils.HexVal(p.Hash, 39-i, 1)
			opacity = utils.Opacity(val)
			fill = utils.FillColor(val)

			styles = make(map[string]interface{})
			styles["fill"] = "none"
			styles["stroke"] = fill
			styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", blockSize)}

			p.Svg.Rect(float64(x)*squareSize+float64(x)*blockSize*2+blockSize/2+blockSize*2, float64(y)*squareSize+float64(y)*blockSize*2+blockSize/2+blockSize*2, blockSize*3, blockSize*3, styles)

			i = i + 1
		}
	}
}

// geoOctagons build the octagons SVG pattern
func (p *Pattern) geoOctagons() {
	squareSize := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 10, 60)
	tile := shapes.BuildOctagon(squareSize)

	p.Svg.SetHeight(int(squareSize * 6))
	p.Svg.SetWidth(int(squareSize * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {
			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity
			styles["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x)*squareSize, float64(y)*squareSize)

			p.Svg.Polyline(tile, styles)

			i = i + 1
		}
	}
}

// geoOverlappingCircles build the overlapping_circles SVG pattern
func (p *Pattern) geoOverlappingCircles() {
	scale := utils.HexVal(p.Hash, 0, 1)
	diameter := utils.Map(scale, 0, 15, 25, 200)
	radius := diameter / 2

	p.Svg.SetHeight(int(radius * 6))
	p.Svg.SetWidth(int(radius * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity)}

			p.Svg.Circle(float64(x)*radius, float64(y)*radius, radius, styles)

			if x == 0 {
				p.Svg.Circle(6*radius, float64(y)*radius, radius, styles)
			}

			if y == 0 {
				p.Svg.Circle(float64(x)*radius, 6*radius, radius, styles)
			}

			if x == 0 && y == 0 {
				p.Svg.Circle(6*radius, 6*radius, radius, styles)
			}

			i = i + 1
		}
	}
}

// geoOverlappingRings build the overlapping_rings SVG pattern
func (p *Pattern) geoOverlappingRings() {
	scale := utils.HexVal(p.Hash, 0, 1)
	ringSize := utils.Map(scale, 0, 15, 10, 60)
	strokeWidth := ringSize / 4

	p.Svg.SetHeight(int(ringSize * 6))
	p.Svg.SetWidth(int(ringSize * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = "none"
			styles["stroke"] = fill
			styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", strokeWidth)}

			p.Svg.Circle(float64(x)*ringSize, float64(y)*ringSize, ringSize-(strokeWidth/2), styles)

			if x == 0 {
				p.Svg.Circle(6*ringSize, float64(y)*ringSize, ringSize-(strokeWidth/2), styles)
			}

			if y == 0 {
				p.Svg.Circle(float64(x)*ringSize, 6*ringSize, ringSize-(strokeWidth/2), styles)
			}

			if x == 0 && y == 0 {
				p.Svg.Circle(6*ringSize, 6*ringSize, ringSize-(strokeWidth/2), styles)
			}

			i = i + 1
		}
	}
}

// geoPlaid build the plaid SVG pattern
func (p *Pattern) geoPlaid() {
	height := 0
	width := 0

	i := 1
	j := 0
	for i <= 18 {

		space := utils.HexVal(p.Hash, j, 1)
		height = height + int(space) + 5

		val := utils.HexVal(p.Hash, j+1, 1)
		opacity := utils.Opacity(val)
		fill := utils.FillColor(val)
		stripeHeight := val + 5

		styles := make(map[string]interface{})
		styles["opacity"] = opacity
		styles["fill"] = fill

		p.Svg.Rect(0, height, "100%", stripeHeight, styles)

		height = height + int(stripeHeight)
		j = j + 2

		i = i + 1
	}

	i = 1
	j = 0
	for i <= 18 {

		space := utils.HexVal(p.Hash, j, 1)
		width = width + int(space) + 5

		val := utils.HexVal(p.Hash, j+1, 1)
		opacity := utils.Opacity(val)
		fill := utils.FillColor(val)
		stripeWidth := val + 5

		styles := make(map[string]interface{})
		styles["opacity"] = opacity
		styles["fill"] = fill

		p.Svg.Rect(width, 0, stripeWidth, "100%", styles)

		width = width + int(stripeWidth)
		j = j + 2

		i = i + 1
	}

	p.Svg.SetHeight(int(height))
	p.Svg.SetWidth(int(width))
}

// geoPlusSigns build the plus_signs SVG pattern
func (p *Pattern) geoPlusSigns() {
	squareSize := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 10, 25)
	plusSize := squareSize * 3
	plusShape := shapes.BuildPlus(squareSize)

	p.Svg.SetHeight(int(squareSize * 12))
	p.Svg.SetWidth(int(squareSize * 12))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			var dx float64
			if y%2 != 0 {
				dx = 1
			}

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity
			styles["style"] = map[string]string{"fill-opacity": fmt.Sprintf("%v", opacity)}

			style := make(map[string]interface{})

			style["transform"] = fmt.Sprintf("translate(%v,%v)", float64(x)*(plusSize-squareSize)+dx*squareSize-squareSize, float64(y)*(plusSize-squareSize)-plusSize/2)
			p.Svg.Group(plusShape, utils.Merge(styles, style))

			if x == 0 {
				style["transform"] = fmt.Sprintf("translate(%v,%v)", 4*plusSize-float64(x)*squareSize+dx*squareSize-squareSize, float64(y)*(plusSize-squareSize)-plusSize/2)
				p.Svg.Group(plusShape, utils.Merge(styles, style))
			}

			if y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v,%v)", float64(x)*(plusSize-squareSize)+dx*squareSize-squareSize, 4*(plusSize)-float64(y)*squareSize-plusSize/2)
				p.Svg.Group(plusShape, utils.Merge(styles, style))
			}

			if x == 0 && y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v,%v)", 4*plusSize-float64(x)*squareSize+dx*squareSize-squareSize, 4*plusSize-float64(y)*squareSize-plusSize/2)
				p.Svg.Group(plusShape, utils.Merge(styles, style))
			}

			i = i + 1
		}
	}
}

// geoSineWaves build the sine_waves SVG pattern
func (p *Pattern) geoSineWaves() {
	period := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 100, 400)
	amplitude := utils.Map(utils.HexVal(p.Hash, 1, 1), 0, 15, 30, 100)
	waveWidth := utils.Map(utils.HexVal(p.Hash, 2, 1), 0, 15, 3, 30)

	p.Svg.SetHeight(int(waveWidth * 36))
	p.Svg.SetWidth(int(period))

	for i := 0; i <= 35; i++ {
		val := utils.HexVal(p.Hash, i, 1)
		opacity := utils.Opacity(val)
		fill := utils.FillColor(val)
		xOffset := (period / 4) * 0.7

		styles := make(map[string]interface{})
		styles["fill"] = "none"
		styles["stroke"] = fill
		styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", waveWidth)}

		str := fmt.Sprintf("M0 %v C %v 0, %v 0, %v %v S %v %v, %v %v S %v 0, %v, %v", amplitude, xOffset, period/2-xOffset, period/2, amplitude, period-xOffset, amplitude*2, period, amplitude, period*1.5-xOffset, period*1.5, amplitude)

		style := make(map[string]interface{})

		style["transform"] = fmt.Sprintf("translate(-%v, %v)", period/4, (waveWidth*float64(i))-(amplitude*1.5))
		p.Svg.Path(str, utils.Merge(styles, style))

		style["transform"] = fmt.Sprintf("translate(-%v, %v)", period/4, (waveWidth*float64(i))-(amplitude*1.5)+waveWidth*36)
		p.Svg.Path(str, utils.Merge(styles, style))
	}
}

// geoSquares build the squares SVG pattern
func (p *Pattern) geoSquares() {
	squareSize := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 10, 60)

	p.Svg.SetHeight(int(squareSize * 6))
	p.Svg.SetWidth(int(squareSize * 6))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity

			p.Svg.Rect(float64(x)*squareSize, float64(y)*squareSize, squareSize, squareSize, styles)

			i = i + 1
		}
	}
}

// geoTessellation build the tessellation SVG pattern
func (p *Pattern) geoTessellation() {
	sideLength := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 5, 40)
	hexHeight := sideLength * math.Sqrt(3)
	hexWidth := sideLength * 2
	triangleHeight := sideLength / 2 * math.Sqrt(3)
	triangle := shapes.BuildRotatedTriangle(sideLength, triangleHeight)
	tileWidth := sideLength*3 + triangleHeight*2
	tileHeight := (hexHeight * 2) + (sideLength * 2)

	p.Svg.SetHeight(int(tileHeight))
	p.Svg.SetWidth(int(tileWidth))

	for i := 0; i <= 19; i++ {
		val := utils.HexVal(p.Hash, i, 1)
		opacity := utils.Opacity(val)
		fill := utils.FillColor(val)

		styles := make(map[string]interface{})
		styles["fill"] = fill
		styles["fill-opacity"] = opacity
		styles["stroke"] = utils.StrokeColor
		styles["stroke-opacity"] = utils.StrokeOpacity
		styles["stroke-width"] = 1

		style := make(map[string]interface{})

		switch i {
		case 0:
			p.Svg.Rect(-sideLength/2, -sideLength/2, sideLength, sideLength, styles)
			p.Svg.Rect(tileWidth-sideLength/2, -sideLength/2, sideLength, sideLength, styles)
			p.Svg.Rect(-sideLength/2, tileHeight-sideLength/2, sideLength, sideLength, styles)
			p.Svg.Rect(tileWidth-sideLength/2, tileHeight-sideLength/2, sideLength, sideLength, styles)
		case 1:
			p.Svg.Rect(hexWidth/2+triangleHeight, hexHeight/2, sideLength, sideLength, styles)
		case 2:
			p.Svg.Rect(-sideLength/2, tileHeight/2-sideLength/2, sideLength, sideLength, styles)
			p.Svg.Rect(tileWidth-sideLength/2, tileHeight/2-sideLength/2, sideLength, sideLength, styles)
		case 3:
			p.Svg.Rect(hexWidth/2+triangleHeight, hexHeight*1.5+sideLength, sideLength, sideLength, styles)
		case 4:
			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v)", sideLength/2, -sideLength/2, sideLength/2, triangleHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v) scale(1, -1)", sideLength/2, tileHeight-(-sideLength/2), sideLength/2, triangleHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 5:
			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v) scale(-1, 1)", tileWidth-sideLength/2, -sideLength/2, sideLength/2, triangleHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v) scale(-1, -1)", tileWidth-sideLength/2, tileHeight+sideLength/2, sideLength/2, triangleHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 6:
			style["transform"] = fmt.Sprintf("translate(%v, %v)", tileWidth/2+sideLength/2, hexHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 7:
			style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", tileWidth-tileWidth/2-sideLength/2, hexHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 8:
			style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, -1)", tileWidth/2+sideLength/2, tileHeight-hexHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 9:
			style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, -1)", tileWidth-tileWidth/2-sideLength/2, tileHeight-hexHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 10:
			style["transform"] = fmt.Sprintf("translate(%v, %v)", sideLength/2, tileHeight/2-sideLength/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 11:
			style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", tileWidth-sideLength/2, tileHeight/2-sideLength/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))
		case 12:
			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(-30, 0, 0)", sideLength/2, sideLength/2)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 13:
			style["transform"] = fmt.Sprintf("scale(-1, 1) translate(%v, %v) rotate(-30, 0, 0)", -tileWidth+sideLength/2, sideLength/2)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 14:
			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(30, 0, %v)", sideLength/2, tileHeight/2-sideLength/2-sideLength, sideLength)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 15:
			style["transform"] = fmt.Sprintf("scale(-1, 1) translate(%v, %v) rotate(30, 0, %v)", -tileWidth+sideLength/2, tileHeight/2-sideLength/2-sideLength, sideLength)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 16:
			style["transform"] = fmt.Sprintf("scale(1, -1) translate(%v, %v) rotate(30, 0, %v)", sideLength/2, -tileHeight+tileHeight/2-sideLength/2-sideLength, sideLength)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 17:
			style["transform"] = fmt.Sprintf("scale(-1, -1) translate(%v, %v) rotate(30, 0, %v)", -tileWidth+sideLength/2, -tileHeight+tileHeight/2-sideLength/2-sideLength, sideLength)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 18:
			style["transform"] = fmt.Sprintf("scale(1, -1) translate(%v, %v) rotate(-30, 0, 0)", sideLength/2, -tileHeight+sideLength/2)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		case 19:
			style["transform"] = fmt.Sprintf("scale(-1, -1) translate(%v, %v) rotate(-30, 0, 0)", -tileWidth+sideLength/2, -tileHeight+sideLength/2)
			p.Svg.Rect(0, 0, sideLength, sideLength, utils.Merge(styles, style))
		}
	}
}

// geoTriangles build the triangles SVG pattern
func (p *Pattern) geoTriangles() {
	scale := utils.HexVal(p.Hash, 0, 1)
	sideLength := utils.Map(scale, 0, 15, 15, 80)
	triangleHeight := sideLength / 2 * math.Sqrt(3)
	triangle := shapes.BuildTriangle(sideLength, triangleHeight)

	p.Svg.SetHeight(int(triangleHeight * 6))
	p.Svg.SetWidth(int(sideLength * 3))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["fill-opacity"] = opacity
			styles["stroke"] = utils.StrokeColor
			styles["stroke-opacity"] = utils.StrokeOpacity

			var rotation int
			if y%2 == 0 {
				if x%2 == 0 {
					rotation = 180
				}
			} else {
				if x%2 != 0 {
					rotation = 180
				}
			}

			style := make(map[string]interface{})

			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(%v, %v, %v)", float64(x)*sideLength*0.5-sideLength/2, triangleHeight*float64(y), rotation, sideLength/2, triangleHeight/2)
			p.Svg.Polyline(triangle, utils.Merge(styles, style))

			if x == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(%v, %v, %v)", 6*sideLength*0.5-sideLength/2, triangleHeight*float64(y), rotation, sideLength/2, triangleHeight/2)
				p.Svg.Polyline(triangle, utils.Merge(styles, style))
			}

			i = i + 1
		}
	}
}

// geoXes build the xes SVG pattern
func (p *Pattern) geoXes() {
	squareSize := utils.Map(utils.HexVal(p.Hash, 0, 1), 0, 15, 10, 25)
	xShape := shapes.BuildPlus(squareSize)
	xSize := squareSize * 3 * 0.943

	p.Svg.SetHeight(int(xSize * 3))
	p.Svg.SetWidth(int(xSize * 3))

	i := 0
	for y := 0; y <= 5; y++ {
		for x := 0; x <= 5; x++ {

			val := utils.HexVal(p.Hash, i, 1)
			opacity := utils.Opacity(val)
			fill := utils.FillColor(val)

			var dy float64
			if x%2 == 0 {
				dy = float64(y)*xSize - xSize*0.5
			} else {
				dy = float64(y)*xSize - xSize*0.5 + xSize/4
			}

			styles := make(map[string]interface{})
			styles["fill"] = fill
			styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity)}

			style := make(map[string]interface{})

			style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", float64(x)*xSize/2-xSize/2, dy-float64(y)*xSize/2, xSize/2, xSize/2)
			p.Svg.Group(xShape, utils.Merge(styles, style))

			if x == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", 6*xSize/2-xSize/2, dy-float64(y)*xSize/2, xSize/2, xSize/2)
				p.Svg.Group(xShape, utils.Merge(styles, style))
			}

			if y == 0 {
				if x%2 == 0 {
					dy = 6*xSize - xSize/2
				} else {
					dy = 6*xSize - xSize/2 + xSize/4
				}

				style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", float64(x)*xSize/2-xSize/2, dy-float64(y)*xSize/2, xSize/2, xSize/2)
				p.Svg.Group(xShape, utils.Merge(styles, style))
			}

			if y == 5 {
				style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", float64(x)*xSize/2-xSize/2, dy-11*xSize/2, xSize/2, xSize/2)
				p.Svg.Group(xShape, utils.Merge(styles, style))
			}

			if x == 0 && y == 0 {
				style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", 6*xSize/2-xSize/2, dy-6*xSize/2, xSize/2, xSize/2)
				p.Svg.Group(xShape, utils.Merge(styles, style))
			}

			i = i + 1
		}
	}
}
