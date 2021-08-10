package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}

	actual := Perimeter(r)
	exp := 40.0

	if actual != exp {
		t.Errorf("Expected %.2f but got %.2f", exp, actual)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, exp float64) {
// 		t.Helper()
// 		actual := shape.Area()
// 		if actual != exp {
// 			t.Errorf("Expected %g but got %g", exp, actual)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		r := Rectangle{12.0, 6.0}

// 		exp := 72.0

// 		checkArea(t, r, exp)
// 	})
// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}

// 		exp := 314.1592653589793

// 		checkArea(t, circle, exp)
// 	})
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.want)
			}
		})
	}
}
