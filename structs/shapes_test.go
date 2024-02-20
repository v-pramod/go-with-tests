package structs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{length: 12, breadth: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}

// 	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Perimeter()

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}

// 	t.Run("check area for rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 15.0}
// 		checkArea(t, rectangle, 150.0)
// 	})

// 	t.Run("check perimeter for rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 15.0}
// 		checkPerimeter(t, rectangle, 54.0)
// 	})

// 	t.Run("check area for circle", func(t *testing.T) {
// 		circle := Circle{8.0}
// 		checkArea(t, circle, 200.96)
// 	})

// 	t.Run("check perimeter for circle", func(t *testing.T) {
// 		circle := Circle{8.0}
// 		checkPerimeter(t, circle, 50.24)
// 	})
// }

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 15.0}
// 	got := rectangle.Perimeter()
// 	want := 50.0

// }

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{10.0, 15.0}
// 	got := rectangle.Area()
// 	want := 150.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestAreaCircle(t *testing.T) {
// 	circle := Circle{8.0}
// 	got := circle.Area()
// 	want := 200.96

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}

// }
