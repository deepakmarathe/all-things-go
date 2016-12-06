package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Printf("OSX")
	case "linux" :
		fmt.Printf("Linux")
	default :
		fmt.Printf("%s.",os)

	}
}
