package main

import "math"

// func Perimeter(width, height float64) float64 {
// 	return 2 * (width + height)
// }

// func Area(width, height float64) float64 {
// 	return width * height
// }

// AFTER refactor
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	raidus float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

// func Perimeter(r Rectangle) float64 {
// 	return 2 * (r.width + r.height)
// }

// func Area(r Rectangle) float64 {
// 	return r.width * r.height
// }

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.raidus * c.raidus
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
