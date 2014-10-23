package pattern

import (
    "fmt"
    "github.com/pravj/geo_pattern/svg"
    "github.com/pravj/geo_pattern/shapes"
    "github.com/pravj/geo_pattern/utils"
)

var Svg = new(svg.SVG)
var hash = utils.Hash("Happy Diwali")

func Start() string {
    generate_background()
    geo_overlapping_rings()

    return Svg.Str()
}

func generate_background() {
    args := make(map[string]interface{})
    args["fill"] = "rgb(120, 160, 200)"

    Svg.Rect(0, 0, "100%", "100%", args)
}

func geo_overlapping_rings() {
    scale := utils.Hex_val(hash, 0 ,1)
    ring_size := utils.Map(scale, 0, 15, 10, 60)
    stroke_width := ring_size / 4

    Svg.Set_height(int(ring_size * 6))
    Svg.Set_width(int(ring_size * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            styles := make(map[string]interface{})
            styles["fill"] = "none"
            styles["stroke"] = fill
            styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", stroke_width)}

            Svg.Circle(float64(x) * ring_size, float64(y) * ring_size, ring_size - (stroke_width/2), styles)

            if x == 0 {
                Svg.Circle(6 * ring_size, float64(y) * ring_size, ring_size - (stroke_width/2), styles) 
            }

            if y == 0 {
                Svg.Circle(float64(x) * ring_size, 6 * ring_size, ring_size - (stroke_width/2), styles)
            }

            if x == 0 && y == 0 {
                Svg.Circle(6 * ring_size, 6 * ring_size, ring_size - (stroke_width/2), styles)
            }

            i = i + 1
        }
    }
}

func geo_concentric_circles() {
    scale := utils.Hex_val(hash, 0, 1)
    ring_size := utils.Map(scale, 0, 15, 10, 60)
    stroke_width := ring_size / 5

    fmt.Println((ring_size + stroke_width) * 6)

    Svg.Set_height(int((ring_size + stroke_width) * 6))
    Svg.Set_width(int((ring_size + stroke_width) * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            cx := float64(x) * ring_size + float64(x) * stroke_width + (ring_size + stroke_width) / 2
            cy := float64(y) * ring_size + float64(y) * stroke_width + (ring_size + stroke_width) / 2

            styles := make(map[string]interface{})
            styles["fill"] = "none"
            styles["stroke"] = fill
            styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", stroke_width)}

            Svg.Circle(cx, cy, ring_size / 2, styles)

            val = utils.Hex_val(hash, 39 - i, 1)
            opacity = utils.Opacity(val)
            fill = utils.Fill_color(val)

            styles = make(map[string]interface{})
            styles["fill"] = fill
            styles["fill-opacity"] = opacity

            Svg.Circle(cx, cy, ring_size / 4, styles)

            i = i + 1
        }
    }
}

func geo_octagons() {
    square_size := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 10, 60)
    tile := shapes.Build_octagon(square_size)

    Svg.Set_height(int(square_size * 6))
    Svg.Set_width(int(square_size * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {
            val := utils.Hex_val(hash, i ,1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            styles := make(map[string]interface{})
            styles["fill"] = fill
            styles["fill-opacity"] = opacity
            styles["stroke"] = utils.STROKE_COLOR
            styles["stroke-opacity"] = utils.STROKE_OPACITY
            styles["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * square_size, float64(y) * square_size)

            Svg.Polyline(tile, styles)

            i = i + 1
        }
    }
}

func geo_overlapping_circles() {
    scale := utils.Hex_val(hash, 0, 1)
    diameter := utils.Map(scale, 0, 15, 25, 200)
    radius := diameter/2

    Svg.Set_height(int(radius * 6))
    Svg.Set_width(int(radius * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            styles := make(map[string]interface{})
            styles["fill"] = fill
            styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity)}

            Svg.Circle(float64(x) * radius, float64(y) * radius, radius, styles)

            if x == 0 {
                Svg.Circle(6 * radius, float64(y) * radius, radius, styles)
            }

            if y == 0 {
                Svg.Circle(float64(x) * radius, 6 * radius, radius, styles)
            }

            if x == 0 && y == 0 {
                Svg.Circle(6 * radius, 6 * radius, radius, styles)
            }

            i = i + 1
        }
    }
}

func geo_squares() {
    square_size := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 10, 60)

    Svg.Set_height(int(square_size * 6))
    Svg.Set_width(int(square_size * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            styles := make(map[string]interface{})
            styles["fill"] = fill
            styles["fill-opacity"] = opacity
            styles["stroke"] = utils.STROKE_COLOR
            styles["stroke-opacity"] = utils.STROKE_OPACITY

            Svg.Rect(float64(x) * square_size, float64(y) * square_size, square_size, square_size, styles)

            i = i + 1
        }
    }
}
