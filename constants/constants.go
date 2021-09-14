package main

const userName = "Juan"
const number = 64 / 2 // allowed

// const random = rand.Int() // This is not allowed because constants cannot be derived from a function call

const (
	InputAttack = iota // special keyword that assigns value to the variable starting from 0 and adding 1
	InputSpecialAttack
	InputHeal
)

func main() {

}
