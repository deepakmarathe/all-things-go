package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	x, y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser = MyFloat(-math.Sqrt2)
	fmt.Println(a.Abs())

	a = &Vertex{3, 4}
	fmt.Println(a.Abs())

}

