package main

import "fmt"

var userName = "Juan" // Global variable that can be used by all functions/structs/etc in this package

func main() {
	shouldContinue := true // Local variable only available for use inside of this function

	//userName := "Juanito" // This will override the variable of the higher scope, in this case the global variable. Called shadowing

	if shouldContinue {
		userAge := 26 // local variable only available for use inside of this IF statement
		fmt.Printf("Name: %v, Age:%v\n", userName, userAge)
	}
}
