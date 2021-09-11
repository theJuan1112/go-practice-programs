package main

import (
	"fmt"
	"math/rand"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}

func generateRandomNumbers() (int, int) {
	number1 := rand.Intn(100)
	number2 := rand.Intn(100)
	return number1, number2
}

func main() {
	a, b := generateRandomNumbers()

	sum := add(a, b)
	fmt.Println(sum)

	printNumber(sum)
}
