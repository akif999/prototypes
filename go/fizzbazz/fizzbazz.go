package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	if isFizz(num) && isBuzz(num) {
		return "Fizzbuzz"
	} else if isFizz(num) {
		return "Fizz"
	} else if isBuzz(num) {
		return "Buzz"
	} else {
		return strconv.Itoa(num)
	}
}

func isFizz(num int) bool {
	return num%3 == 0
}

func isBuzz(num int) bool {
	return num%5 == 0
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
