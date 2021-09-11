package main

import "fmt"

func main() {
	age := 26
	fmt.Println(age)

	// The & is used to get the address of the variable in memory
	myAge := &age

	fmt.Println(myAge)
	// The * will get the value that is stored at that address in  memory
	fmt.Println(*myAge)

	*myAge = 33

	fmt.Println(*myAge)
	// Here the age will change because we changed the value of the variable stored at the address.
	fmt.Println(age)

}
