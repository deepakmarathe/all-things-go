package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1], a)

	primes := [6]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	var b []int = primes[0:4]
	fmt.Println(b)

	b[0] = 10;
	fmt.Println(primes)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := [] struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(s)
	fmt.Println(s[0:2])
	fmt.Println(s[:1])
	fmt.Println(s[1:])

	printSlice(s[0:3])
	printSlice(s[0:2])

	var t []int
	fmt.Println(t)
	if t == nil {
		fmt.Println("nil")
	}

	x := make([]struct{ i int; b bool }, 10)
	printSlice(x)
	x = make([]struct{ i int; b bool }, 0, 10)
	printSlice(x)
	x = x[:3]
	printSlice(x)

	fmt.Println()
	fmt.Println("-- slices of slice ---")
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	fmt.Println(board)
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	x = append(x, struct{i int; b bool}{10, false});
	x = append(x, struct{i int; b bool}{20, false});
	fmt.Println(x)
}

func printSlice(a []struct{ i int; b bool }) {
	fmt.Printf("len = %d, cap = %d, %v\n", len(a), cap(a), a)
}
