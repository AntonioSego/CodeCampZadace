package main

import "fmt"

func main() {

	/* 1. Declare two integers, `firstNumber` and `secondNumber`, assign values 20 and 30 to them.
	   Swap values of `firstNumber` and `secondNumber` without using third variable
	   After all, print values of `firstNumber` and `secondNumber`.
	*/

	var firstNumber int = 20
	var secondNumber int = 30

	fmt.Println("firstNumber: ", firstNumber, " secondNumber: ", secondNumber)

	firstNumber = firstNumber + secondNumber
	secondNumber = firstNumber - secondNumber
	firstNumber = firstNumber - secondNumber

	fmt.Println("firstNumber: ", firstNumber, " secondNumber: ", secondNumber)

	/* 2. Declare two variables, `firstName` and `lastName` assign them with wanted values.
	Declare constant named `fullname`
	Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.
	*/

	/*
		var firstName string = "Antonio"
		var lastName string = "Sego"
		const fullName = firstName + " " + lastName // ne mozemo konstanti dodijeliti promjenjive vrijednosti
	*/
	const firstName string = "Antonio"
	const lastName string = "Sego"
	const fullName = firstName + " " + lastName

	fmt.Println(fullName)
}
