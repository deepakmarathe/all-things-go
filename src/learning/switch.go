package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Printf("OSX")
	case "linux" :
		fmt.Printf("Linux")
	default :
		fmt.Printf("%s.", os)

	}
	fmt.Println()
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday();
	switch time.Saturday {
	case today + 0 :
		fmt.Printf("Today.")
	case today + 1:
		fmt.Printf("Tomorrow")
	case today + 2:
		fmt.Printf("Day after tomorrow")
	default :
		fmt.Printf("Too far away")
	}
}
