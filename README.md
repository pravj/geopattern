geo_pattern
===========
> now create beautiful generative background images from a strinig in golang.
> > Go port of [Jason Long](https://github.com/jasonlong)'s awesome [GeoPattern](https://github.com/jasonlong/geo_pattern) library.

[![GoDoc](https://godoc.org/github.com/pravj/geo_pattern?status.svg)](http://godoc.org/github.com/pravj/geo_pattern)

![Nested Squares Pattern](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/pattern.png)

Generate beautiful tiling SVG patterns from a string. The string is converted into a SHA and a color and pattern are determined based on the values in the hash. The color is determined by shifting the hue from a default (or passed in) base color. One of 16 patterns is used (or you can specify one) and the sizing of the pattern elements is also determined by the hash values.

You can use the generated pattern as the background-image for a container. Using the base64 representation of the pattern still results in SVG rendering, so it looks great on retina displays.

See the GitHub Guides [site](https://guides.github.com) as an example of this library in action. GitHub Guides use [Original](https://github.com/jasonlong/geo_pattern) Ruby implementation.

####Installation
`go get github.com/pravj/geo_pattern`

####Usage
[Example](https://github.com/pravj/geo_pattern/tree/master/examples) directory contains sample go programs that explains use of `geo_pattern`

####API

#####Arguments for functions returning pattern's string

######`phrase` : custom pattern phrase
```
args := map[string]string{"phrase": "My Custom Phrase"}
```

######`generator` : custom pattern type
```
args := map[string]string{"generator": "plaid"}
```

######`color` : custom background color
```
args := map[string]string{"color": "#3b5998"}
```

######`base_color` : custom base color that decides background color
```
args := map[string]string{"base_color": "#ffcc00"}
```
---
#####Functions provided by package for pattern's string representation

######Get the SVG string :
```
Generate(args)
```
######Get the Base64 encoded string :
```
Base64_string(args)
```
######Get uri image string :
```
Uri_image(args)
```

####Available Pattern

###### chevrons
![Chevrons](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/chevrons.png)

###### concentric_circles
![Concentric_Circles](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/concentric_circles.png)

###### hexagons
![Hexagons](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/hexagons.png)

###### mosaic_squares
![Mosaic_Squares](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/mosaic_squares.png)

###### nested_squares
![Nested_Squares](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/nested_squares.png)

###### octagons
![Octagons](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/octagons.png)

###### overlapping_circles
![Overlapping_Circles](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/overlapping_circles.png)

###### overlapping_rings
![Overlapping_rings](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/overlapping_rings.png)

###### plaid
![Plaid](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/plaid.png)

###### plus_signs
![Plus_Signs](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/plus_signs.png)

###### sine_waves
![Sine_Waves](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/sine_waves.png)

###### squares
![Squares](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/squares.png)

###### tessellation
![Tessellation](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/tessellation.png)

###### triangles
![Triangles](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/triangles.png)

###### xes
![Xes](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/patterns/xes.png)

####Dependencies
[go-colorful](https://github.com/lucasb-eyer/go-colorful) : for color space conversion
