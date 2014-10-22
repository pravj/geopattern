package pattern

import (
    //"fmt"
    "github.com/pravj/geo_pattern/svg"
)

var Svg = new(svg.SVG)

func Start() string {
    Svg.Set_height(100)
    Svg.Set_width(100)

    generate_background()

    return Svg.Str()
}

func generate_background() {
    args := make(map[string]interface{})
    args["fill"] = "rgb(50, 60, 70)"

    Svg.Rect("0", "0", "100%", "100%", args)
}

func geo_squares() {
    //
}
