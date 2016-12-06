package main
import "fmt"

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{2,1}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func main(){
	fmt.Println(v1, p, v2, v3)

	v := Vertex{10,20}
	fmt.Println(v)
	v.x = 30
	fmt.Println(v)

	p := &v
	p.x = 100
	fmt.Println(v)
}
