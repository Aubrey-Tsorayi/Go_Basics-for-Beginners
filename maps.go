package main

import (
	"fmt"
)

func main() {
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecage"] = 4

	//fmt.Println(vertices)
	//ft.Println(vertices["triangle"])

	delete(vertices, "sqaure")
	fmt.Println(vertices)
}
