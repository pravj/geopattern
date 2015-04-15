package geopattern

import (
	"io/ioutil"
	"testing"
)

func TestPatternGenerators(t *testing.T) {

	p := Pattern{
		Color:     "#123456",
		BaseColor: "#123456",
		Phrase:    "test",
	}

	for _, v := range PATTERNS {
		p.Generator = v

		svg := Generate(p)

		f, err := ioutil.ReadFile("tests/" + v + ".svg")

		if err != nil {
			t.Errorf("Could not open file " + v + ".svg for testing.")
		} else {
			if svg != string(f) {
				t.Errorf("For Pattern %v, Expected: \n %v \n\n\n Got: \n %v", v, string(f), svg)
			}
		}
	}
}

func TestNoPhraseGenerator(t *testing.T) {
	//Set fake time
	setTime(12345)

	p := Pattern{
		Generator: PATTERNS[9],
		Color:     "#123456",
		BaseColor: "#123456",
	}

	svg := Generate(p)

	expectedSvg := "<svg xmlns='http://www.w3.org/2000/svg' width='432' height='432'><rect x='0' y='0' width='100%' height='100%' fill='rgb(7, 21, 50)'  /><rect x='0' y='16' width='100%' height='9' fill='#ddd' opacity='0.05466666666666667'  /><rect x='0' y='42' width='100%' height='10' fill='#222' opacity='0.06333333333333332'  /><rect x='0' y='72' width='100%' height='11' fill='#ddd' opacity='0.072'  /><rect x='0' y='90' width='100%' height='6' fill='#222' opacity='0.028666666666666674'  /><rect x='0' y='114' width='100%' height='14' fill='#222' opacity='0.098'  /><rect x='0' y='143' width='100%' height='12' fill='#222' opacity='0.08066666666666666'  /><rect x='0' y='163' width='100%' height='16' fill='#222' opacity='0.11533333333333333'  /><rect x='0' y='196' width='100%' height='7' fill='#ddd' opacity='0.03733333333333333'  /><rect x='0' y='211' width='100%' height='10' fill='#222' opacity='0.06333333333333332'  /><rect x='0' y='226' width='100%' height='8' fill='#222' opacity='0.046'  /><rect x='0' y='249' width='100%' height='6' fill='#222' opacity='0.028666666666666674'  /><rect x='0' y='265' width='100%' height='15' fill='#ddd' opacity='0.10666666666666666'  /><rect x='0' y='295' width='100%' height='14' fill='#222' opacity='0.098'  /><rect x='0' y='323' width='100%' height='15' fill='#ddd' opacity='0.10666666666666666'  /><rect x='0' y='351' width='100%' height='18' fill='#222' opacity='0.13266666666666665'  /><rect x='0' y='375' width='100%' height='11' fill='#ddd' opacity='0.072'  /><rect x='0' y='399' width='100%' height='13' fill='#ddd' opacity='0.08933333333333333'  /><rect x='0' y='417' width='100%' height='15' fill='#ddd' opacity='0.10666666666666666'  /><rect x='16' y='0' width='9' height='100%' fill='#ddd' opacity='0.05466666666666667'  /><rect x='42' y='0' width='10' height='100%' fill='#222' opacity='0.06333333333333332'  /><rect x='72' y='0' width='11' height='100%' fill='#ddd' opacity='0.072'  /><rect x='90' y='0' width='6' height='100%' fill='#222' opacity='0.028666666666666674'  /><rect x='114' y='0' width='14' height='100%' fill='#222' opacity='0.098'  /><rect x='143' y='0' width='12' height='100%' fill='#222' opacity='0.08066666666666666'  /><rect x='163' y='0' width='16' height='100%' fill='#222' opacity='0.11533333333333333'  /><rect x='196' y='0' width='7' height='100%' fill='#ddd' opacity='0.03733333333333333'  /><rect x='211' y='0' width='10' height='100%' fill='#222' opacity='0.06333333333333332'  /><rect x='226' y='0' width='8' height='100%' fill='#222' opacity='0.046'  /><rect x='249' y='0' width='6' height='100%' fill='#222' opacity='0.028666666666666674'  /><rect x='265' y='0' width='15' height='100%' fill='#ddd' opacity='0.10666666666666666'  /><rect x='295' y='0' width='14' height='100%' fill='#222' opacity='0.098'  /><rect x='323' y='0' width='15' height='100%' fill='#ddd' opacity='0.10666666666666666'  /><rect x='351' y='0' width='18' height='100%' fill='#222' opacity='0.13266666666666665'  /><rect x='375' y='0' width='11' height='100%' fill='#ddd' opacity='0.072'  /><rect x='399' y='0' width='13' height='100%' fill='#ddd' opacity='0.08933333333333333'  /><rect x='417' y='0' width='15' height='100%' fill='#ddd' opacity='0.10666666666666666'  /></svg>"

	if svg != expectedSvg {
		t.Errorf("Expected: \n %v \n\n\n Got: \n %v", expectedSvg, svg)
	}

}

func TestNoPhraseBase64(t *testing.T) {
	//Set fake time
	setTime(12345)

	p := Pattern{
		Generator: PATTERNS[9],
		Color:     "#123456",
		BaseColor: "#123456",
	}

	svg := Base64String(p)

	expectedSvg := "PHN2ZyB4bWxucz0naHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmcnIHdpZHRoPSc0MzInIGhlaWdodD0nNDMyJz48cmVjdCB4PScwJyB5PScwJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMDAlJyBmaWxsPSdyZ2IoNywgMjEsIDUwKScgIC8+PHJlY3QgeD0nMCcgeT0nMTYnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzknIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMDU0NjY2NjY2NjY2NjY2NjcnICAvPjxyZWN0IHg9JzAnIHk9JzQyJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMCcgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wNjMzMzMzMzMzMzMzMzMzMicgIC8+PHJlY3QgeD0nMCcgeT0nNzInIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzExJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjA3MicgIC8+PHJlY3QgeD0nMCcgeT0nOTAnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzYnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDI4NjY2NjY2NjY2NjY2Njc0JyAgLz48cmVjdCB4PScwJyB5PScxMTQnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzE0JyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA5OCcgIC8+PHJlY3QgeD0nMCcgeT0nMTQzJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMicgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wODA2NjY2NjY2NjY2NjY2NicgIC8+PHJlY3QgeD0nMCcgeT0nMTYzJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxNicgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4xMTUzMzMzMzMzMzMzMzMzMycgIC8+PHJlY3QgeD0nMCcgeT0nMTk2JyB3aWR0aD0nMTAwJScgaGVpZ2h0PSc3JyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjAzNzMzMzMzMzMzMzMzMzMzJyAgLz48cmVjdCB4PScwJyB5PScyMTEnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzEwJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA2MzMzMzMzMzMzMzMzMzMyJyAgLz48cmVjdCB4PScwJyB5PScyMjYnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzgnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDQ2JyAgLz48cmVjdCB4PScwJyB5PScyNDknIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzYnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDI4NjY2NjY2NjY2NjY2Njc0JyAgLz48cmVjdCB4PScwJyB5PScyNjUnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzE1JyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjEwNjY2NjY2NjY2NjY2NjY2JyAgLz48cmVjdCB4PScwJyB5PScyOTUnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzE0JyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA5OCcgIC8+PHJlY3QgeD0nMCcgeT0nMzIzJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxNScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4xMDY2NjY2NjY2NjY2NjY2NicgIC8+PHJlY3QgeD0nMCcgeT0nMzUxJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxOCcgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4xMzI2NjY2NjY2NjY2NjY2NScgIC8+PHJlY3QgeD0nMCcgeT0nMzc1JyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4wNzInICAvPjxyZWN0IHg9JzAnIHk9JzM5OScgd2lkdGg9JzEwMCUnIGhlaWdodD0nMTMnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMDg5MzMzMzMzMzMzMzMzMzMnICAvPjxyZWN0IHg9JzAnIHk9JzQxNycgd2lkdGg9JzEwMCUnIGhlaWdodD0nMTUnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMTA2NjY2NjY2NjY2NjY2NjYnICAvPjxyZWN0IHg9JzE2JyB5PScwJyB3aWR0aD0nOScgaGVpZ2h0PScxMDAlJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjA1NDY2NjY2NjY2NjY2NjY3JyAgLz48cmVjdCB4PSc0MicgeT0nMCcgd2lkdGg9JzEwJyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDYzMzMzMzMzMzMzMzMzMzInICAvPjxyZWN0IHg9JzcyJyB5PScwJyB3aWR0aD0nMTEnIGhlaWdodD0nMTAwJScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4wNzInICAvPjxyZWN0IHg9JzkwJyB5PScwJyB3aWR0aD0nNicgaGVpZ2h0PScxMDAlJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjAyODY2NjY2NjY2NjY2NjY3NCcgIC8+PHJlY3QgeD0nMTE0JyB5PScwJyB3aWR0aD0nMTQnIGhlaWdodD0nMTAwJScgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wOTgnICAvPjxyZWN0IHg9JzE0MycgeT0nMCcgd2lkdGg9JzEyJyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDgwNjY2NjY2NjY2NjY2NjYnICAvPjxyZWN0IHg9JzE2MycgeT0nMCcgd2lkdGg9JzE2JyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMTE1MzMzMzMzMzMzMzMzMzMnICAvPjxyZWN0IHg9JzE5NicgeT0nMCcgd2lkdGg9JzcnIGhlaWdodD0nMTAwJScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4wMzczMzMzMzMzMzMzMzMzMycgIC8+PHJlY3QgeD0nMjExJyB5PScwJyB3aWR0aD0nMTAnIGhlaWdodD0nMTAwJScgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wNjMzMzMzMzMzMzMzMzMzMicgIC8+PHJlY3QgeD0nMjI2JyB5PScwJyB3aWR0aD0nOCcgaGVpZ2h0PScxMDAlJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA0NicgIC8+PHJlY3QgeD0nMjQ5JyB5PScwJyB3aWR0aD0nNicgaGVpZ2h0PScxMDAlJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjAyODY2NjY2NjY2NjY2NjY3NCcgIC8+PHJlY3QgeD0nMjY1JyB5PScwJyB3aWR0aD0nMTUnIGhlaWdodD0nMTAwJScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4xMDY2NjY2NjY2NjY2NjY2NicgIC8+PHJlY3QgeD0nMjk1JyB5PScwJyB3aWR0aD0nMTQnIGhlaWdodD0nMTAwJScgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wOTgnICAvPjxyZWN0IHg9JzMyMycgeT0nMCcgd2lkdGg9JzE1JyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMTA2NjY2NjY2NjY2NjY2NjYnICAvPjxyZWN0IHg9JzM1MScgeT0nMCcgd2lkdGg9JzE4JyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMTMyNjY2NjY2NjY2NjY2NjUnICAvPjxyZWN0IHg9JzM3NScgeT0nMCcgd2lkdGg9JzExJyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMDcyJyAgLz48cmVjdCB4PSczOTknIHk9JzAnIHdpZHRoPScxMycgaGVpZ2h0PScxMDAlJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjA4OTMzMzMzMzMzMzMzMzMzJyAgLz48cmVjdCB4PSc0MTcnIHk9JzAnIHdpZHRoPScxNScgaGVpZ2h0PScxMDAlJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjEwNjY2NjY2NjY2NjY2NjY2JyAgLz48L3N2Zz4="

	if svg != expectedSvg {
		t.Errorf("Expected: \n %v \n\n\n Got: \n %v", expectedSvg, svg)
	}

}

func TestNoPhraseURI(t *testing.T) {
	//Set fake time
	setTime(12345)

	p := Pattern{
		Generator: PATTERNS[9],
		Color:     "#123456",
		BaseColor: "#123456",
	}

	svg := URIimage(p)

	expectedSvg := "url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0naHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmcnIHdpZHRoPSc0MzInIGhlaWdodD0nNDMyJz48cmVjdCB4PScwJyB5PScwJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMDAlJyBmaWxsPSdyZ2IoNywgMjEsIDUwKScgIC8+PHJlY3QgeD0nMCcgeT0nMTYnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzknIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMDU0NjY2NjY2NjY2NjY2NjcnICAvPjxyZWN0IHg9JzAnIHk9JzQyJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMCcgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wNjMzMzMzMzMzMzMzMzMzMicgIC8+PHJlY3QgeD0nMCcgeT0nNzInIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzExJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjA3MicgIC8+PHJlY3QgeD0nMCcgeT0nOTAnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzYnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDI4NjY2NjY2NjY2NjY2Njc0JyAgLz48cmVjdCB4PScwJyB5PScxMTQnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzE0JyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA5OCcgIC8+PHJlY3QgeD0nMCcgeT0nMTQzJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMicgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wODA2NjY2NjY2NjY2NjY2NicgIC8+PHJlY3QgeD0nMCcgeT0nMTYzJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxNicgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4xMTUzMzMzMzMzMzMzMzMzMycgIC8+PHJlY3QgeD0nMCcgeT0nMTk2JyB3aWR0aD0nMTAwJScgaGVpZ2h0PSc3JyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjAzNzMzMzMzMzMzMzMzMzMzJyAgLz48cmVjdCB4PScwJyB5PScyMTEnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzEwJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA2MzMzMzMzMzMzMzMzMzMyJyAgLz48cmVjdCB4PScwJyB5PScyMjYnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzgnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDQ2JyAgLz48cmVjdCB4PScwJyB5PScyNDknIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzYnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDI4NjY2NjY2NjY2NjY2Njc0JyAgLz48cmVjdCB4PScwJyB5PScyNjUnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzE1JyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjEwNjY2NjY2NjY2NjY2NjY2JyAgLz48cmVjdCB4PScwJyB5PScyOTUnIHdpZHRoPScxMDAlJyBoZWlnaHQ9JzE0JyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA5OCcgIC8+PHJlY3QgeD0nMCcgeT0nMzIzJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxNScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4xMDY2NjY2NjY2NjY2NjY2NicgIC8+PHJlY3QgeD0nMCcgeT0nMzUxJyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxOCcgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4xMzI2NjY2NjY2NjY2NjY2NScgIC8+PHJlY3QgeD0nMCcgeT0nMzc1JyB3aWR0aD0nMTAwJScgaGVpZ2h0PScxMScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4wNzInICAvPjxyZWN0IHg9JzAnIHk9JzM5OScgd2lkdGg9JzEwMCUnIGhlaWdodD0nMTMnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMDg5MzMzMzMzMzMzMzMzMzMnICAvPjxyZWN0IHg9JzAnIHk9JzQxNycgd2lkdGg9JzEwMCUnIGhlaWdodD0nMTUnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMTA2NjY2NjY2NjY2NjY2NjYnICAvPjxyZWN0IHg9JzE2JyB5PScwJyB3aWR0aD0nOScgaGVpZ2h0PScxMDAlJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjA1NDY2NjY2NjY2NjY2NjY3JyAgLz48cmVjdCB4PSc0MicgeT0nMCcgd2lkdGg9JzEwJyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDYzMzMzMzMzMzMzMzMzMzInICAvPjxyZWN0IHg9JzcyJyB5PScwJyB3aWR0aD0nMTEnIGhlaWdodD0nMTAwJScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4wNzInICAvPjxyZWN0IHg9JzkwJyB5PScwJyB3aWR0aD0nNicgaGVpZ2h0PScxMDAlJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjAyODY2NjY2NjY2NjY2NjY3NCcgIC8+PHJlY3QgeD0nMTE0JyB5PScwJyB3aWR0aD0nMTQnIGhlaWdodD0nMTAwJScgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wOTgnICAvPjxyZWN0IHg9JzE0MycgeT0nMCcgd2lkdGg9JzEyJyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMDgwNjY2NjY2NjY2NjY2NjYnICAvPjxyZWN0IHg9JzE2MycgeT0nMCcgd2lkdGg9JzE2JyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMTE1MzMzMzMzMzMzMzMzMzMnICAvPjxyZWN0IHg9JzE5NicgeT0nMCcgd2lkdGg9JzcnIGhlaWdodD0nMTAwJScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4wMzczMzMzMzMzMzMzMzMzMycgIC8+PHJlY3QgeD0nMjExJyB5PScwJyB3aWR0aD0nMTAnIGhlaWdodD0nMTAwJScgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wNjMzMzMzMzMzMzMzMzMzMicgIC8+PHJlY3QgeD0nMjI2JyB5PScwJyB3aWR0aD0nOCcgaGVpZ2h0PScxMDAlJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjA0NicgIC8+PHJlY3QgeD0nMjQ5JyB5PScwJyB3aWR0aD0nNicgaGVpZ2h0PScxMDAlJyBmaWxsPScjMjIyJyBvcGFjaXR5PScwLjAyODY2NjY2NjY2NjY2NjY3NCcgIC8+PHJlY3QgeD0nMjY1JyB5PScwJyB3aWR0aD0nMTUnIGhlaWdodD0nMTAwJScgZmlsbD0nI2RkZCcgb3BhY2l0eT0nMC4xMDY2NjY2NjY2NjY2NjY2NicgIC8+PHJlY3QgeD0nMjk1JyB5PScwJyB3aWR0aD0nMTQnIGhlaWdodD0nMTAwJScgZmlsbD0nIzIyMicgb3BhY2l0eT0nMC4wOTgnICAvPjxyZWN0IHg9JzMyMycgeT0nMCcgd2lkdGg9JzE1JyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMTA2NjY2NjY2NjY2NjY2NjYnICAvPjxyZWN0IHg9JzM1MScgeT0nMCcgd2lkdGg9JzE4JyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyMyMjInIG9wYWNpdHk9JzAuMTMyNjY2NjY2NjY2NjY2NjUnICAvPjxyZWN0IHg9JzM3NScgeT0nMCcgd2lkdGg9JzExJyBoZWlnaHQ9JzEwMCUnIGZpbGw9JyNkZGQnIG9wYWNpdHk9JzAuMDcyJyAgLz48cmVjdCB4PSczOTknIHk9JzAnIHdpZHRoPScxMycgaGVpZ2h0PScxMDAlJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjA4OTMzMzMzMzMzMzMzMzMzJyAgLz48cmVjdCB4PSc0MTcnIHk9JzAnIHdpZHRoPScxNScgaGVpZ2h0PScxMDAlJyBmaWxsPScjZGRkJyBvcGFjaXR5PScwLjEwNjY2NjY2NjY2NjY2NjY2JyAgLz48L3N2Zz4=);"

	if svg != expectedSvg {
		t.Errorf("Expected: \n %v \n\n\n Got: \n %v", expectedSvg, svg)
	}

}
