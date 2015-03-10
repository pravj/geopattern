geopattern
===========
> Create beautiful generative image patterns from a string in golang.
> > Go port of [Jason Long](https://github.com/jasonlong)'s awesome [GeoPattern](https://github.com/jasonlong/geo_pattern) library.

[![GoDoc](https://godoc.org/github.com/pravj/geopattern?status.svg)](http://godoc.org/github.com/pravj/geopattern)

> Read geopattern's development story [**geo_pattern: going on the Go path**](http://pravj.github.io/blog/2014/11/03/going-on-the-go-path/)

![Nested Squares Pattern](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/pattern.png)

Generate beautiful tiling SVG patterns from a string. The string is converted into a SHA and a color and pattern are determined based on the values in the hash. The color is determined by shifting the hue from a default (or passed in) base color. One of 16 patterns is used (or you can specify one) and the sizing of the pattern elements is also determined by the hash values.

You can use the generated pattern as the background-image for a container. Using the base64 representation of the pattern still results in SVG rendering, so it looks great on retina displays.

See the GitHub Guides [site](https://guides.github.com) as an example of this library in action. GitHub Guides use [Original](https://github.com/jasonlong/geo_pattern) Ruby implementation.

####Installation
`go get github.com/pravj/geopattern`

####Usage
[Example](https://github.com/pravj/geopattern/tree/master/examples) directory contains sample go programs that explains use of `geopattern`

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

######`baseColor` : custom base color that decides background color
```
args := map[string]string{"baseColor": "#ffcc00"}
```
---
#####Functions provided by package for pattern's string representation

######Get the SVG string :
```
Generate(args)
```
######Get the Base64 encoded string :
```
Base64String(args)
```
######Get uri image string :
```
URIimage(args)
```

####Available Pattern

###### chevrons
![Chevrons](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/chevrons.png)

###### concentric-circles
![Concentric-Circles](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/concentric-circles.png)

###### diamonds
![Diamonds](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/diamonds.png)

###### hexagons
![Hexagons](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/hexagons.png)

###### mosaic-squares
![Mosaic-Squares](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/mosaic-squares.png)

###### nested-squares
![Nested-Squares](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/nested-squares.png)

###### octagons
![Octagons](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/octagons.png)

###### overlapping-circles
![Overlapping-Circles](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/overlapping-circles.png)

###### overlapping-rings
![Overlapping-rings](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/overlapping-rings.png)

###### plaid
![Plaid](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/plaid.png)

###### plus-signs
![Plus-Signs](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/plus-signs.png)

###### sine-waves
![Sine-Waves](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/sine-waves.png)

###### squares
![Squares](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/squares.png)

###### tessellation
![Tessellation](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/tessellation.png)

###### triangles
![Triangles](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/triangles.png)

###### xes
![Xes](https://raw.githubusercontent.com/pravj/geopattern/master/examples/patterns/xes.png)

####Dependencies
[go-colorful](https://github.com/lucasb-eyer/go-colorful) : for color space conversion

---

Made with *Muzi* and *Appy* by [Pravendra Singh](https://pravj.github.io)
