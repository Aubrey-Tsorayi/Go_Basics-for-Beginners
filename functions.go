package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//result := sum(2, 3)
	result, err := sqrt(16) //-16

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefiened for negative number")
	}

	return math.Sqrt(x), nil
}
