package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

// Receiver method
func (v Vertex) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

//function
func abs(v Vertex) float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}


func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex)Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Scale(v *Vertex, f float64 ) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	x := Vertex{3, 4}
	fmt.Println(x.abs())
	t := abs(x);
	fmt.Println(t)

	f := MyFloat(math.SqrtE)
	fmt.Println(f.Abs())

	(&x).Scale(10)
	fmt.Println(abs(x))
	x.Scale(0.1)
	fmt.Println(x.abs())
	Scale(&x, 10);
	fmt.Println(abs(x))

	fmt.Println("--- interface ---")
	var a Abser = &x
	fmt.Println(a.Abs())
	//a = f;
	//fmt.Println(a.Abs())
}
