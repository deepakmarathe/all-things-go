package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m := make(map[string]Vertex)
	m["bangalore"] = Vertex{
		20.2, 40.4,
	}
	fmt.Println(m)

	var mm = map[string]Vertex{
		"a" : Vertex{20.1,100.1},
		"b" : {10.0,10},
	}
	fmt.Println(mm)

	a := make(map[string]string);
	a["deepak"] = "marathe"
	fmt.Println(a)
	a["deepak"] = "deepak"
	fmt.Println(a)
	delete(a,"deepak")
	fmt.Println(a)
	v, ok := a["deepak"]
	fmt.Println("value : ", v, "exists : ", ok )

strings.Fields(s)
}
