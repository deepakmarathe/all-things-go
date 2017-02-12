package main

import "fmt"

func do(i interface{}) {
	switch s := i.(type) {
	case int : fmt.Println("int : ", s)
	case string : fmt.Println("string : ", s)
	default : fmt.Println("default : ", s)
	}
}
func main() {
	do(1)
	do("hello")
	do(34.3)
	do(false)
}