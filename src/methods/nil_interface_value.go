package main

import "fmt"

type I interface {
	M()
}

func main(){
	var i I
	i.M(); //causes runtime error.
}

func describe(i I){
	fmt.Println("(%v, %T)", i, i)
}