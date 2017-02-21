package main

import (
	"fmt"
	"time"
)

type MyError struct {
	What string
	When time.Time
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{"didnt work", time.Now()}
}

func main() {
	if error := run(); error != nil {
		fmt.Println(error)
	}
}
