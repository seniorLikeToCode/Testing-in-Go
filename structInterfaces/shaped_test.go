package shape

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{
			Height: 12,
			Width:  6,
		}

		checkPerimeter(t, rectangle, 36)
	})

	t.Run("Circle Perimeter", func(t *testing.T) {
		circle := Circle{
			Radius: 10,
		}

		checkPerimeter(t, circle, 62.83185307179586)
	})
}
