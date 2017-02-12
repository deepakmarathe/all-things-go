package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age  int
}

type IpAddress [4]byte

func (ip IpAddress) String() string {
	var s = make([]string, len(ip));
	for i, v := range ip[0:len(ip)] {
		s[i] = strconv.Itoa(int(v))
	}
	return strings.Join(s, ".");
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	var p Person = Person{"a", 10}
	q := Person{"b", 20}
	fmt.Println(p, q)

	ip := map[string]IpAddress{
		"loopback": {127, 0, 0, 1},
		"googledns": {8, 8, 8, 8},
	}

	i := ip["loopback"];
	fmt.Println(i.String())
}
