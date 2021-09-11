package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserChoice() (string, error) {
	fmt.Println("Please Make Your Choice!")
	fmt.Println("1) Add up all the numbers ou to number x")
	fmt.Println("2) Build Factorial up to number x")
	fmt.Println("3) Sum up manually entered numbers ")
	fmt.Println("4) Sum up a list of entered numbers")

	fmt.Print("Please Make Your Choice: ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\r\n", "", -1)

	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}
}

func getInputNumber() (int, error) {

	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\r\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(chosenNumber), nil

}

func calculateSumUptoNumber() {
	fmt.Print("Please input your number: ")
	input, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	sum := 0

	for i := 1; i <= input; i++ {
		sum = sum + i
	}

	fmt.Printf("Result: %v", sum)

}

func calculateFactorial() {
	fmt.Print("Please Input your number: ")
	input, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	fact := 1

	for i := 1; i <= input; i++ {
		fact *= i
	}

	fmt.Printf("Result: %v", fact)

}

func calculateSumManually() {
	isEntering := true
	sum := 0

	fmt.Println("Keep entering numbers, the program will stop once you enter not a number")

	for isEntering {
		input, err := getInputNumber()
		sum += input
		isEntering = err == nil
	}

	fmt.Printf("Result: %v\n", sum)

}

func calculateSumList() {

	fmt.Println("Please enter comma-separated list of numbers.")
	inputNumberList, err := reader.ReadString('\n')

	if err != nil {

		fmt.Println("Invalid Input")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\r\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, ",")

	sum := 0
	for _, value := range inputNumbers {
		number, _ := strconv.ParseInt(value, 0, 64)

		sum += int(number)
	}

	fmt.Printf("Result: %v\n", sum)
}

func main() {

	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid Choice, exiting!")
		return
	}

	if selectedChoice == "1" {
		calculateSumUptoNumber()

	} else if selectedChoice == "2" {
		calculateFactorial()

	} else if selectedChoice == "3" {
		calculateSumManually()

	} else {
		calculateSumList()

	}

}
