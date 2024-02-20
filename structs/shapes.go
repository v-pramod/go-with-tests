package structs

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.breadth + r.length)
}

func (r Rectangle) Area() float64 {
	return r.breadth * r.length
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.base + t.height)
}

func (t Triangle) Perimeter() float64 {
	// considering equilaterial triangle
	return 3 * t.base
}
