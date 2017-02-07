package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

//receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func abs(v Vertex) float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) scale(f float64){
	v.x = v.x * f;
	v.y = v.y * f;
}

func main() {
	v := Vertex{4, 3};
	fmt.Println(v.Abs())
	fmt.Println(abs(v))

	f := MyFloat(-10.0);
	fmt.Println(f.abs())

	v.scale(5)
	fmt.Println(v)
	fmt.Println(abs(v))
}

