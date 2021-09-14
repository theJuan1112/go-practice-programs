package main

import "fmt"

func main() {
	numbers := []int{1, 5, 12, 3}
	sum := sumup(1, 5, 10, 64)
	anotherSum := sumup(numbers...) // This will return the values of the slice.
	fmt.Println(sum)
	fmt.Println(anotherSum)

}

func sumup(numbers ...int) int { // Instead of passing the array, this feature allows you to just enter manually a list
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
