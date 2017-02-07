package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var i, t = 1.0, 1;
	for ;t<10;t++{
		i = i - ((i*i) - x)/(2*i)
		fmt.Println(i)
	}
	return i
}

func main() {
	fmt.Println(Sqrt(2))
}
