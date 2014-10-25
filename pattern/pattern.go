package pattern

import (
    "fmt"
    "math"
    "github.com/pravj/geo_pattern/svg"
    "github.com/pravj/geo_pattern/shapes"
    "github.com/pravj/geo_pattern/utils"
)

var Svg = new(svg.SVG)
var hash = utils.Hash("gunnu")

func Start() string {
    generate_background()
    geo_tessellation()

    return Svg.Str()
}

func generate_background() {
    args := make(map[string]interface{})
    args["fill"] = "rgb(50, 74, 157)"

    Svg.Rect(0, 0, "100%", "100%", args)
}

func geo_tessellation() {
    side_length := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 5, 40)
    hex_height := side_length * math.Sqrt(3)
    hex_width := side_length * 2
    triangle_height := side_length/2 * math.Sqrt(3)
    triangle := shapes.Build_rotated_triangle(side_length, triangle_height)
    tile_width := side_length*3 + triangle_height*2
    tile_height := (hex_height * 2) + (side_length * 2)

    Svg.Set_height(int(tile_height))
    Svg.Set_width(int(tile_width))

    for i := 0; i <= 19; i++ {
        val := utils.Hex_val(hash, i, 1)
        opacity := utils.Opacity(val)
        fill := utils.Fill_color(val)

        styles := make(map[string]interface{})
        styles["fill"] = fill
        styles["fill-opacity"] = opacity
        styles["stroke"] = utils.STROKE_COLOR
        styles["stroke-opacity"] = utils.STROKE_OPACITY
        styles["stroke-width"] = 1

        style := make(map[string]interface{})

        switch i {
        case 0:
            Svg.Rect(-side_length/2, -side_length/2, side_length, side_length, styles)
            Svg.Rect(tile_width - side_length/2, -side_length/2, side_length, side_length, styles)
            Svg.Rect(-side_length/2, tile_height-side_length/2, side_length, side_length, styles)
            Svg.Rect(tile_width - side_length/2, tile_height-side_length/2, side_length, side_length, styles)
        case 1:
            Svg.Rect(hex_width/2 + triangle_height, hex_height/2, side_length, side_length, styles)
        case 2:
            Svg.Rect(-side_length/2, tile_height/2-side_length/2, side_length, side_length, styles)
            Svg.Rect(tile_width-side_length/2, tile_height/2-side_length/2, side_length, side_length, styles)
        case 3:
            Svg.Rect(hex_width/2 + triangle_height, hex_height * 1.5 + side_length, side_length, side_length, styles)
        case 4:
            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v)", side_length/2, -side_length/2, side_length/2, triangle_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v) scale(1, -1)", side_length/2, tile_height-(-side_length/2), side_length/2, triangle_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 5:
            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v) scale(-1, 1)", tile_width-side_length/2, -side_length/2, side_length/2, triangle_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(0, %v, %v) scale(-1, -1)", tile_width-side_length/2, tile_height+side_length/2, side_length/2, triangle_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 6:
            style["transform"] = fmt.Sprintf("translate(%v, %v)", tile_width/2+side_length/2, hex_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 7:
            style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", tile_width-tile_width/2-side_length/2, hex_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 8:
            style["transform"] = fmt.Sprintf("translate(%v, %v) scale(1, -1)", tile_width/2+side_length/2, tile_height-hex_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 9:
            style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, -1)", tile_width-tile_width/2-side_length/2, tile_height-hex_height/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 10:
            style["transform"] = fmt.Sprintf("translate(%v, %v)", side_length/2, tile_height/2 - side_length/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 11:
            style["transform"] = fmt.Sprintf("translate(%v, %v) scale(-1, 1)", tile_width-side_length/2, tile_height/2-side_length/2)
            Svg.Polyline(triangle, utils.Merge(styles, style))
        case 12:
            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(-30, 0, 0)", side_length/2, side_length/2)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 13:
            style["transform"] = fmt.Sprintf("scale(-1, 1) translate(%v, %v) rotate(-30, 0, 0)", -tile_width+side_length/2, side_length/2)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 14:
            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(30, 0, %v)", side_length/2, tile_height/2-side_length/2-side_length, side_length)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 15:
            style["transform"] = fmt.Sprintf("scale(-1, 1) translate(%v, %v) rotate(30, 0, %v)", -tile_width+side_length/2, tile_height/2-side_length/2-side_length, side_length)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 16:
            style["transform"] = fmt.Sprintf("scale(1, -1) translate(%v, %v) rotate(30, 0, %v)", side_length/2, -tile_height+tile_height/2-side_length/2-side_length, side_length)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 17:
            style["transform"] = fmt.Sprintf("scale(-1, -1) translate(%v, %v) rotate(30, 0, %v)", -tile_width+side_length/2, -tile_height+tile_height/2-side_length/2-side_length, side_length)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 18:
            style["transform"] = fmt.Sprintf("scale(1, -1) translate(%v, %v) rotate(-30, 0, 0)", side_length/2, -tile_height+side_length/2)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        case 19:
            style["transform"] = fmt.Sprintf("scale(-1, -1) translate(%v, %v) rotate(-30, 0, 0)", -tile_width+side_length/2, -tile_height+side_length/2)
            Svg.Rect(0, 0, side_length, side_length, utils.Merge(styles, style))
        }
    }
}

func geo_mosaic_squares() {
    triangle_size := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 15, 50)

    Svg.Set_height(int(triangle_size * 8))
    Svg.Set_width(int(triangle_size * 8))

    i := 0
    for y := 0; y <= 3; y++ {
        for x := 0; x <= 3; x++ {

            values := [2]float64{utils.Hex_val(hash, i, 1), utils.Hex_val(hash, i + 1, 1)}

            if x % 2 == 0 {
                if y % 2 == 0 {
                    shapes.Draw_outer_mosaic_tile(Svg, float64(x)*triangle_size*2, float64(y)*triangle_size*2, triangle_size, utils.Hex_val(hash, i, 1))
                } else {
                    shapes.Draw_inner_mosaic_tile(Svg, float64(x)*triangle_size*2, float64(y)*triangle_size*2, triangle_size, values)
                }
            } else {
                if y % 2 == 0 {
                    shapes.Draw_inner_mosaic_tile(Svg, float64(x)*triangle_size*2, float64(y)*triangle_size*2, triangle_size, values)
                } else {
                    shapes.Draw_outer_mosaic_tile(Svg, float64(x)*triangle_size*2, float64(y)*triangle_size*2, triangle_size, utils.Hex_val(hash, i, 1))
                }
            }

            i = i + 1
        }
    }

}

func geo_xes() {
    square_size := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 10, 25)
    x_shape := shapes.Build_plus(square_size)
    x_size := square_size * 3 * 0.943

    Svg.Set_height(int(x_size * 3))
    Svg.Set_width(int(x_size * 3))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            var dy float64
            if x % 2 == 0 {
                dy = float64(y)*x_size - x_size*0.5
            } else {
                dy = float64(y)*x_size - x_size*0.5 + x_size/4
            }

            styles := make(map[string]interface{})
            styles["fill"] = fill
            styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity)}

            style := make(map[string]interface{})

            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", float64(x)*x_size/2 - x_size/2, dy - float64(y)*x_size/2, x_size/2, x_size/2)
            Svg.Group(x_shape, utils.Merge(styles, style))

            if x == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", 6*x_size/2 - x_size/2, dy - float64(y)*x_size/2, x_size/2, x_size/2)
                Svg.Group(x_shape, utils.Merge(styles, style))
            }

            if y == 0 {
                if x % 2 == 0 {
                    dy = 6*x_size - x_size/2
                } else {
                    dy = 6*x_size - x_size/2 + x_size/4
                }

                style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", float64(x)*x_size/2 - x_size/2, dy - float64(y)*x_size/2, x_size/2, x_size/2)
                Svg.Group(x_shape, utils.Merge(styles, style))
            }

            if y == 5 {
                style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", float64(x)*x_size/2 - x_size/2, dy - 11*x_size/2, x_size/2, x_size/2)
                Svg.Group(x_shape, utils.Merge(styles, style))
            }

            if x == 0 && y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(45, %v, %v)", 6*x_size/2 - x_size/2, dy - 6*x_size/2, x_size/2, x_size/2)
                Svg.Group(x_shape, utils.Merge(styles, style))
            }

            i = i + 1
        }
    }
}

func geo_plus_signs() {
    square_size := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 10, 25)
    plus_size := square_size * 3
    plus_shape := shapes.Build_plus(square_size)

    Svg.Set_height(int(square_size * 12))
    Svg.Set_width(int(square_size * 12))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            var dx float64
            if y % 2 != 0 {
                dx = 1
            }

            styles := make(map[string]interface{})
            styles["fill"] = fill
            styles["stroke"] = utils.STROKE_COLOR
            styles["stroke-opacity"] = utils.STROKE_OPACITY
            styles["style"] = map[string]string{"fill-opacity": fmt.Sprintf("%v", opacity)}

            style := make(map[string]interface{})

            style["transform"] = fmt.Sprintf("translate(%v,%v)", float64(x)*(plus_size - square_size) + dx*square_size - square_size, float64(y) * (plus_size - square_size) - plus_size/2)
            Svg.Group(plus_shape, utils.Merge(styles, style))

            if x == 0 {
                style["transform"] = fmt.Sprintf("translate(%v,%v)", 4*plus_size - float64(x)*square_size + dx*square_size - square_size, float64(y) * (plus_size - square_size) - plus_size/2)
                Svg.Group(plus_shape, utils.Merge(styles, style))
            }

            if y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v,%v)", float64(x)*(plus_size - square_size) + dx*square_size - square_size, 4*(plus_size) - float64(y)*square_size - plus_size/2)
                Svg.Group(plus_shape, utils.Merge(styles, style))
            }

            if x == 0 && y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v,%v)", 4*plus_size - float64(x)*square_size + dx*square_size - square_size, 4*plus_size - float64(y)*square_size - plus_size/2)
                Svg.Group(plus_shape, utils.Merge(styles, style))
            }

            i = i + 1
        }
    }
}

func geo_chevrons() {
    chevron_width := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 30, 80)
    chevron_height := chevron_width
    chevron := shapes.Build_chevron(chevron_width, chevron_height)

    Svg.Set_height(int(chevron_height * 6 * 0.66))
    Svg.Set_width(int(chevron_width * 6))

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
            styles["stroke-width"] = 1

            style := make(map[string]interface{})

            style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * chevron_width, float64(y) * chevron_height * 0.66 - chevron_height/2)
            Svg.Group(chevron, utils.Merge(styles, style))

            if y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * chevron_width, 6 * chevron_height * 0.66 - chevron_height/2)
                Svg.Group(chevron, utils.Merge(styles, style))
            }

            i = i + 1
        }
    }
}

func geo_diamonds() {
    diamond_width := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 10, 50)
    diamond_height := utils.Map(utils.Hex_val(hash, 1, 1), 0, 15, 10, 50)
    diamond := shapes.Build_diamond(diamond_width, diamond_height)

    Svg.Set_height(int(diamond_height * 3))
    Svg.Set_width(int(diamond_width * 6))

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

            var dx float64
            if y % 2 != 0 {
                dx = diamond_width / 2
            }

            style := make(map[string]interface{})

            style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * diamond_width - diamond_width / 2 + dx, diamond_height/2 * float64(y) - diamond_height / 2)
            Svg.Polyline(diamond, utils.Merge(styles, style))

            if x == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v)", 6 * diamond_width - diamond_width / 2 + dx, diamond_height/2 * float64(y) - diamond_height / 2)
                Svg.Polyline(diamond, utils.Merge(styles, style))
            }

            if y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v)", float64(x) * diamond_width - diamond_width / 2 + dx, diamond_height/2 * 6 - diamond_height / 2)
                Svg.Polyline(diamond, utils.Merge(styles, style))
            }

            if x == 0 && y == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v)", 6 * diamond_width - diamond_width / 2 + dx, diamond_height/2 * 6 - diamond_height / 2)
                Svg.Polyline(diamond, utils.Merge(styles, style))
            }

            i = i + 1
        }
    }
}

func geo_triangles() {
    scale := utils.Hex_val(hash, 0, 1)
    side_length := utils.Map(scale, 0, 15, 15, 80)
    triangle_height := side_length/2 * math.Sqrt(3)
    triangle := shapes.Build_triangle(side_length, triangle_height)

    Svg.Set_height(int(triangle_height * 6))
    Svg.Set_width(int(side_length * 3))

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

            var rotation int
            if y % 2 == 0 {
                if x % 2 == 0 {
                    rotation = 180
                }
            } else {
                if x % 2 != 0 {
                    rotation = 180
                }
            }

            style := make(map[string]interface{})

            style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(%v, %v, %v)", float64(x) * side_length * 0.5 - side_length / 2, triangle_height * float64(y), rotation, side_length / 2, triangle_height / 2)
            Svg.Polyline(triangle, utils.Merge(styles, style))

            if x == 0 {
                style["transform"] = fmt.Sprintf("translate(%v, %v) rotate(%v, %v, %v)", 6 * side_length * 0.5 - side_length / 2, triangle_height * float64(y), rotation, side_length / 2, triangle_height / 2)
                Svg.Polyline(triangle, utils.Merge(styles, style))
            }

            i = i + 1
        }
    }
}

func geo_sine_waves() {
    period := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 100, 400)
    amplitude := utils.Map(utils.Hex_val(hash, 1, 1), 0, 15, 30, 100)
    wave_width := utils.Map(utils.Hex_val(hash, 2, 1), 0, 15, 3, 30)

    Svg.Set_height(int(wave_width * 36))
    Svg.Set_width(int(period))

    for i := 0; i <= 35; i++ {
        val := utils.Hex_val(hash, i, 1)
        opacity := utils.Opacity(val)
        fill := utils.Fill_color(val)
        x_offset := (period/4) * 0.7

        styles := make(map[string]interface{})
        styles["fill"] = "none"
        styles["stroke"] = fill
        styles["style"] = map[string]string{"opacity": fmt.Sprintf("%v", opacity), "stroke-width": fmt.Sprintf("%vpx", wave_width)}

        str := fmt.Sprintf("M0 %v C %v 0, %v 0, %v %v S %v %v, %v %v S %v 0, %v, %v", amplitude, x_offset, period / 2 - x_offset, period / 2, amplitude, period - x_offset, amplitude * 2, period, amplitude, period * 1.5 - x_offset, period * 1.5, amplitude)

        style := make(map[string]interface{})

        style["transform"] = fmt.Sprintf("translate(-%v, %v)", period / 4, (wave_width * float64(i)) - (amplitude * 1.5))
        Svg.Path(str, utils.Merge(styles, style))

        style["transform"] = fmt.Sprintf("translate(-%v, %v)", period / 4, (wave_width * float64(i)) - (amplitude * 1.5) + wave_width * 36)
        Svg.Path(str, utils.Merge(styles, style))
    }
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
