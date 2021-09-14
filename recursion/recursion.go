package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}
