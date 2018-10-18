package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fibo(i))
	}
}

func fibo(x int) int {
	sqrt5 := math.Sqrt(5)
	return int(1 / sqrt5 * (math.Pow((1+sqrt5)/2, float64(x)) - (math.Pow((1-sqrt5)/2, float64(x)))))
}
