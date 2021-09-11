package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	createdDate time.Time
}

func (user User) outputData() {

	fmt.Printf("My name is %v %v and was born on %v", user.firstName, user.lastName, user.birthdate)

}

func NewUser(fname string, lname string, bdate string) *User {
	user := User{
		fname,
		lname,
		bdate,
		time.Now(),
	}

	return &user
}

func main() {
	var newUser User

	firstname := getUserData("Please enter your first name: ")
	lastname := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter you birthdate (MM/DD/YYYY): ")

	newUser = *NewUser(firstname, lastname, birthdate)

	newUser.outputData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\r\n", "", -1)
	return cleanedInput
}
