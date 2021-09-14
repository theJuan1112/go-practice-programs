package main

import "fmt"

type person struct { //structs are a special type but there are more
	name string
	age  int
}

type customNumber int

// You cant add custom functions to built in types like int, instead you can build a custom type and add that functionality
func (number customNumber) pow(power int) customNumber {
	result := number
	for i := 1; i < power; i++ {
		result = result * number
	}

	return result
}

type personData map[string]person

func main() {
	var people personData = personData{
		"p1": {"Juan", 26},
	}
	fmt.Println(people)

	var dummy customNumber = 5
	powerNumber := dummy.pow(3)

	fmt.Println(powerNumber)
}
