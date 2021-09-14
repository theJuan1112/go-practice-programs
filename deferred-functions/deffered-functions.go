package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userInput := getUserInput()

	storeData(userInput)
}

func getUserInput() string {
	fmt.Println("Please enter the text that should be stored.")
	fmt.Print("Your input: ")

	reader := bufio.NewReader(os.Stdin)

	enteredText, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Failed to read user input.")
		return ""
	}

	return enteredText
}

func storeData(data string) {
	file, err := os.Create("data.txt")

	if err != nil {
		fmt.Println("Creating the file failed!")
		panic(err)
	}
	defer func() {
		err := file.Close() //built-in keyword that will execute whatever is in front of it once the whole function executes first.

		if err != nil {
			fmt.Println("FAILED TO CLOSE FILE")
			return
		}
	}()

	file.WriteString(data)

	fmt.Println("Successfully stored data in file!")
}
