package main

import (
    "fmt"
    "github.com/pravj/pattern/svg"
)

func main() {
    svg := new(svg.SVG)

    svg.Set_height(100)
    svg.Set_width(100)

    svg.Rect(1,2,4,5)
    svg.Circle(1,2,3)
    svg.Path("path_string")
    svg.Polyline("polyline_string")

    args  := make(map[string]interface{})

    args["first"] = "pravendra"
    args["last"] = "singh"
    args["age"] = 13

    args["things"] = map[string]string{"alpha": "beta"}

    fmt.Println(svg.Write_args(args))
    fmt.Println(svg.Str())
}
