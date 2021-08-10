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
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
