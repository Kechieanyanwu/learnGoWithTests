package shapes

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Rectangle struct { //I have a rectangle type with an Area method
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct { //I have a circle type with an Area method
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct { // Triangle type also with an Area method
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Shape interface { //both rectangles and circles have an area method returning a float64, so they satisfy the interface.
	Area() float64
}

// Essentially, the interface is just checking for whether they have an Area function that returns a float64. Does it? It satisfies the interface - HURRAY!
// Note that the implementation doesnt matter. How they do it doesnt matter. Do they have it? Yes? It satisfies the interface - HURRAY!
// This means you could have an empty function and it still satisfies, so this doesnt guarantee correct implementation. You still have to
// - make sure that implementation is correct
