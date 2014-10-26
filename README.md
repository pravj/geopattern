geo_pattern
===========
> now create beautiful generative background images from a strinig in golang.
> > Go port of [Jason Long](https://github.com/jasonlong)'s awesome [GeoPattern](https://github.com/jasonlong/geo_pattern) library.

![Nested Squares Pattern](https://raw.githubusercontent.com/pravj/geo_pattern/master/examples/pattern.png)

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
* chevrons
* concentric_circles
* diamonds
*	hexagons
* mosaic_squares
* nested_squares
* octagons
* overlapping_circles
* overlapping_rings
* plaid
* plus_signs
* sine_waves
* squares
* tessellation
* triangles
* xes

####Dependencies
[go-colorful](https://github.com/lucasb-eyer/go-colorful) : for color space conversion
