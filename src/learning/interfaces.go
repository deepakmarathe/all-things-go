package main

import
(
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64
//
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	lat, long float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.lat * v.lat + v.long * v.long)
}

func main() {
	var a Abser
	a = MyFloat(math.Sqrt2);
	fmt.Println("a : " , a , "a.Abs() : " , a.Abs())
	a = Vertex{1.0, 2.0}
	fmt.Println("a : " , a, "a.Abs() : ", a.Abs())
	a = &Vertex{1,2}
	fmt.Println(a.Abs())
}
