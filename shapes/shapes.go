package shapes

import (
    "fmt"
    "math"
)

func Build_octagon(square_size float64) string {
    s := square_size
    c := 0.33 * s

    return fmt.Sprintf("%v,0,%v,0,%v,%v,%v,%v,%v,%v,%v,%v,0,%v,0,%v,%v,0", c, s-c, s, c, s, s-c, s-c, s, c, s, s-c, c, c)
}

func Build_triangle(side_length, height float64) string {
    half_width := side_length / 2

    return fmt.Sprintf("%v,0,%v,%v,0,%v,%v,0", half_width, side_length, height, height, half_width)
}

func Build_diamond(width, height float64) string {
    return fmt.Sprintf("%v,0,%v,%v,%v,%v,0,%v", width/2, width, height/2, width/2, height, height/2)
}

func Build_right_triangle(side_length float64) string {
    return fmt.Sprintf("0,0,%v,%v,0,%v,0,0", side_length, side_length, side_length)
}

func Rotated_triangle(side_length, width float64) string {
    half_height := side_length / 2

    return fmt.Sprintf("0,0,%v,%v,0,%v,0,0", width, half_height, side_length)
}

func Build_hexagon(side_length float64) string {
    c := side_length
    a := c / 2
    b := math.Sin(60 * math.Pi / 180) * c

    return fmt.Sprintf("0,%v,%v,0,%v,0,%v,%v,%v,%v,%v,%v,0,%v", b, a, a + c, 2 * c, b, a + c, 2 * b, a, 2 * b, b)
}

func Build_chevron(width, height float64) [2]string {
    e := height * 0.66
    var elements [2]string

    elements[0] = fmt.Sprintf("<polyline points='0,0,%v,%v,%v,%v,0,%v,0,0' />", width/2, height-e, width/2, height, e)
    elements[1] = fmt.Sprintf("<polyline points='%v,%v,%v,0,%v,%v,%v,%v,%v,%v' />", width/2, height-e, width, width, e, width/2, height, width/2, height-e)

    return elements
}
