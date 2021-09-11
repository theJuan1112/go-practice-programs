package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	age, err := getUserAge()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	if (age >= 30 && age < 50) || age > 60 {
		fmt.Println("You are now a senior, LOL")
	} else if age >= 50 {
		fmt.Println("You are now eligible for palliative care, RIP")
	} else if age >= 18 {
		fmt.Print("Welcome to the Club")
	} else {
		fmt.Println("Sorry, you are not old enough :(")
	}
}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter your Age:")

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	age, err := strconv.ParseInt(input, 0, 64)

	return int(age), err

}
