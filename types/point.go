package types

import "fmt"

type Point struct {
	X int32
	Y int32
}

func NewPoint(x, y int32) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func MakePoint(x, y int32) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
