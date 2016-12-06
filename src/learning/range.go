package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i:= range pow {
		pow[i] = 1 << uint(i)
	}
	for _,v := range pow {
		fmt.Printf("%d \n", v)
	}
}
