package main

import(
	"fmt"
)

func main()  {
	i, j := 21, 2902

	p := &i;
	fmt.Println(*p)
	*p=100
	fmt.Println(*p)

	p = &j;
	*p = *p / 20
	fmt.Println(*p)
}