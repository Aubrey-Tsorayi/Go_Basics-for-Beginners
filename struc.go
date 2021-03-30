package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Tutsi", age: 19}
	fmt.Println(p)
	//fmt.Println(p.age)
}
