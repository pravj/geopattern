package pattern

import (
    "fmt"
    "math"
    "github.com/pravj/geo_pattern/svg"
    "github.com/pravj/geo_pattern/shapes"
    "github.com/pravj/geo_pattern/utils"
)

var Svg = new(svg.SVG)
var hash = utils.Hash("Happy Diwali")

func Start() string {
    generate_background()
    geo_hexagons()

    return Svg.Str()
}

func generate_background() {
    args := make(map[string]interface{})
    args["fill"] = "rgb(120, 160, 200)"

    Svg.Rect(0, 0, "100%", "100%", args)
}

func geo_hexagons() {
    scale := utils.Hex_val(hash, 0, 1)
    side_length := utils.Map(scale, 0, 15, 8, 60)
    hex_height := side_length * math.Sqrt(3)
    hex_width := side_length * 2
    hex := shapes.Build_hexagon(side_length)

    Svg.Set_height(int(hex_height * 6))
    Svg.Set_width(int((hex_width * 3) + (side_length * 3)))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            var dy float64
            if x % 2 == 0 {
                dy = float64(y) * hex_height
            } else {
                dy = float64(y) * hex_height + hex_height / 2
            }
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            styles := make(map[string]interface{})
            styles["fill"] = fill
            styles["fill-opacity"] = opacity
            styles["stroke"] = utils.STROKE_COLOR
            styles["stroke-opacity"] = utils.STROKE_OPACITY

            style := make(map[string]interface{})

            style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * side_length * 1.5 - hex_width / 2, dy - hex_height / 2)
            Svg.Polyline(hex, utils.Merge(styles, style))

            if x == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v)", 6 * side_length * 1.5 - hex_width / 2, dy - hex_height / 2)
                Svg.Polyline(hex, utils.Merge(styles, style))
            }

            if y == 0 {
                if x % 2 == 0 {
                    dy = 6 * hex_height
                } else {
                    dy = 6 * hex_height + hex_height / 2
                }
                style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * side_length * 1.5 - hex_width / 2, dy - hex_height / 2)
                Svg.Polyline(hex, utils.Merge(styles, style))
            }

            if x == 0 && y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v)", 6 * side_length * 1.5 - hex_width / 2, 5 * hex_height + hex_height / 2)
                Svg.Polyline(hex, utils.Merge(styles, style))
            }

            i = i + 1
        }
    }
}

func geo_plaid() {
    height := 0
    width := 0

    i := 1
    j := 0
    for i <= 18 {

        space := utils.Hex_val(hash, j, 1)
        height = height + int(space) + 5

        val := utils.Hex_val(hash, j + 1, 1)
        opacity := utils.Opacity(val)
        fill := utils.Fill_color(val)
        stripe_height := val + 5

        styles := make(map[string]interface{})
        styles["opacity"] = opacity
        styles["fill"] = fill

        Svg.Rect(0, height, "100%", stripe_height, styles)

        height = height + int(stripe_height)
        j = j + 2

        i = i + 1
    }

    i = 1
    j = 0
    for i <= 18 {

        space := utils.Hex_val(hash, j, 1)
        width = width + int(space) + 5

        val := utils.Hex_val(hash, j + 1, 1)
        opacity := utils.Opacity(val)
        fill := utils.Fill_color(val)
        stripe_width := val + 5

        styles := make(map[string]interface{})
        styles["opacity"] = opacity
        styles["fill"] = fill

        Svg.Rect(width, 0, stripe_width, "100%", styles)

        width = width + int(stripe_width)
        j = j + 2

        i = i + 1
    }

    Svg.Set_height(int(height))
    Svg.Set_width(int(width))
}

func geo_nested_squares() {
    block_size := utils.Map(utils.Hex_val(hash, 0 ,1), 0, 15, 4, 12)
    square_size := block_size * 7

    Svg.Set_height(int((square_size + block_size) * 6 + block_size * 6))
    Svg.Set_width(int((square_size + block_size) * 6 + block_size * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            styles := make(map[string]interface{})
            styles["fill"] = "none"
            styles["stroke"] = fill
            styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", block_size)}

            Svg.Rect(float64(x) * square_size + float64(x) * block_size * 2 + block_size / 2, float64(y) * square_size + float64(y) * block_size * 2 + block_size / 2, square_size, square_size, styles)

            val = utils.Hex_val(hash, 39 - i, 1)
            opacity = utils.Opacity(val)
            fill = utils.Fill_color(val)

            styles = make(map[string]interface{})
            styles["fill"] = "none"
            styles["stroke"] = fill
            styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", block_size)}

            Svg.Rect(float64(x) * square_size + float64(x) * block_size * 2 + block_size / 2 + block_size * 2, float64(y) * square_size + float64(y) * block_size * 2 + block_size / 2 + block_size * 2, block_size * 3, block_size * 3, styles)

            i = i + 1
        }
    }
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
