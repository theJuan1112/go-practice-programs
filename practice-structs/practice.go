package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	id          string
	title       string
	description string
	price       float64
}

// func newBook(id string, title string, description string, price float64) *Book {
// 	return &Book{id, title, description, price}

// }

func (book Book) outputData() {
	fmt.Printf("ID: %v | Title: %v | Description: %v | Price: USD  $%v \n", book.id, book.title, book.description, book.price)
}

func (book Book) writeToFile() {
	file, _ := os.Create(book.id + ".txt")

	content := fmt.Sprintf("ID: %v | Title: %v | Description: %v | Price: USD  $%v \n", book.id, book.title, book.description, book.price)

	file.WriteString(content)
	file.Close()
}

func main() {
	createdBook := getBookInfo()

	createdBook.outputData()
	createdBook.writeToFile()

}

func getBookInfo() *Book {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Book Info")
	fmt.Println("---------------------")

	idInput := cleanInput(reader, "Book ID: ")
	titleInput := cleanInput(reader, "Book Title: ")
	descriptionInput := cleanInput(reader, "Book Description: ")
	priceInput := cleanInput(reader, "Book Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	book := Book{idInput, titleInput, descriptionInput, priceValue}

	return &book
}

func cleanInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)

	return input
}

// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
