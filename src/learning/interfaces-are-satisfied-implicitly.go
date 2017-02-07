package main
import (
	"fmt"
	"math"
)

type T struct {
	S string
}
type I interface {
	M()
}

func (t *T) M(){
	fmt.Println(t.S)
}

type MyFloat float64

func (f MyFloat)M(){
	fmt.Println(f)
}

func describe(i I){
	fmt.Printf("(%v, %T)\n", i, i)
}
func main(){
	var i I = &T{"hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()
}

