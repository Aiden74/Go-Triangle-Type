package main

// Triangle is a stuct representing a triangle.
type Triangle struct {
	A int
	B int
	C int
}

// GetTriangleType returns the type of the triangle
func (t *Triangle) GetTriangleType() string {
	if t.A >= t.B+t.C || t.B >= t.A+t.C || t.C >= t.B+t.A {
		return "NotATriangle"
	} else if t.A == t.B && t.A == t.C {
		return "Equilateral"
	} else if t.A == t.B || t.A == t.C || t.B == t.C {
		return "Isosceles"
	} else {
		return "Scalene"
	}
}
