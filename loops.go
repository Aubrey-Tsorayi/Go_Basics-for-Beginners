//Go only has for loop
package main

import (
	"fmt"
)

func main() {
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	//while loop format
	//i := 0

	//for i < 5 {
	//	fmt.Println(i)
	//	i++
	//}

	//arr := []string{"a", "b", "c"}

	//for index, value := range arr {
	//	fmt.Println("index", index, "value", value)
	//}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

}
