package main

import (
	"fmt"
)

func main() {
	//var a [5]int
	//a[2] = 7
	//b := [5]int{5, 3, 2, 6, 67}
	c := []int{1, 2, 3, 4, 5}
	c = append(c, 13)

	fmt.Println(c)
}
