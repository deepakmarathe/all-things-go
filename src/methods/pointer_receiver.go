package main

import ( "fmt" )

type Vertex struct {
	x, y float64
}

func (v *Vertex) scale(f float64){
	v.x = v.x * f;
	v.y = v.y * f;
}

func scale(v *Vertex, f float64){
	v.x = v.x * f;
	v.y = v.y * f;
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v)

	scale(&v, 2)
	fmt.Println(v)

	v.scale(10)
	fmt.Println(v)

	p := &Vertex{3,4}
	p.scale(10)
	fmt.Println(v)
}

