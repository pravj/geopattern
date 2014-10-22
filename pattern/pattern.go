package pattern

import (
    //"fmt"
    "github.com/pravj/geo_pattern/svg"
    "github.com/pravj/geo_pattern/utils"
)

var Svg = new(svg.SVG)
var hash = utils.Hash("nkman")

func Start() string {
    generate_background()
    geo_squares()

    return Svg.Str()
}

func generate_background() {
    args := make(map[string]interface{})
    args["fill"] = "rgb(255, 0, 150)"

    Svg.Rect(0, 0, "100%", "100%", args)
}

func geo_squares() {
    square_size := utils.Map(utils.Hex_val(hash, 0, 1), 0, 15, 10, 60)

    Svg.Set_height(int(square_size * 6))
    Svg.Set_width(int(square_size * 6))

    i := 0
    for y := 0; y <= 5; y++ {
        for x := 0; x <= 5 ; x++ {

            val := utils.Hex_val(hash, i, 1)
            opacity := utils.Opacity(val)
            fill := utils.Fill_color(val)

            args := make(map[string]interface{})
            args["fill"] = fill
            args["fill-opacity"] = opacity
            args["stroke"] = utils.STROKE_COLOR
            args["stroke-opacity"] = utils.STROKE_OPACITY

            Svg.Rect(float64(x) * square_size, float64(y) * square_size, square_size, square_size, args)

            i = i + 1
        }
    }
}
