package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/theJuan1112/go-practice-programs/bmi-calc/info"
)

func banner() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
}

func askInput() (clean float64) {
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	clean, _ = strconv.ParseFloat(input, 64)
	return
}

func bmi(height float64, weight float64) float64 {
	return weight / (height * height)
}

func main() {
	banner()

	fmt.Print(info.WeightPrompt)
	weight := askInput()

	fmt.Print(info.HeightPrompt)
	height := askInput()

	bmi := bmi(height, weight)

	fmt.Printf("Your BMI is: %.2f", bmi)

}
