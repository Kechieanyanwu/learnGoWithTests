package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

// Introducing table driven tests. Useful for building a list of test cases that can be tested in the same manner

func TestArea(t *testing.T) {
	areaTests := []struct { // this just creates the type of the struct
		name    string
		shape   Shape
		hasArea float64
	}{ //this is the actual data in the struct
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 360.0},
	}

	for _, tt := range areaTests {
		// use tt.name as the name for t.Run
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, want %g", tt.shape, tt.shape.Area(), tt.hasArea)
			}
		})
	}

}

// - Tests should be assertions of truth, and not just a sequence of operations with an output
// Previous version of test

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) { //helper function that allows me check the area of the shape no matter the type of shape
// 		t.Helper()
// 		got := shape.Area() // will only work if the type passed satisfies the shape interface
// 		if got != want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		// got := rectangle.Area()
// 		want := 72.0
// 		checkArea(t, rectangle, want)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793

// 		checkArea(t, circle, want)

// 	})
// }
