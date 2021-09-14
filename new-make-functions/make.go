package main

import "fmt"

func main() {
	hobbies := make([]string, 2, 10) // Go will behind the scenes, create a slice with 2 slots occupied but with a capacity of 10 already preassigned
	// This pre allocation could be a huge time saver because you do not have to create a new slice every time you call append

	moreHobbies := new([]string) // used to allocate space in memory. It returns the pointer.

	fmt.Println(moreHobbies)

	hobbies[0] = "Reading"
	hobbies[1] = "Coding"

}
