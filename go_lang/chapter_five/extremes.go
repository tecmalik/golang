package main

import "fmt"

func main() {
	fmt.Print("Enter the number of values you would like to enter : ")
	var numbersOfValues = 0
	_, err := fmt.Scan(&numbersOfValues)
	if err != nil {
		fmt.Println("Error reading input")
	}
	var firstNumber int
	var lastNumber int
	var maximumNumber = 0
	var minimumNumber = 0
	for count := 0; count < numbersOfValues; count++ {
		fmt.Printf("Enter value %d : ", count+1)
		var userInput = 0
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println("Error reading input")
		}
		if count == 0 {
			firstNumber = userInput
			maximumNumber = userInput
			minimumNumber = userInput
		}
		if count == numbersOfValues-1 {
			lastNumber = userInput
		}
		if userInput > maximumNumber {
			maximumNumber = userInput
		}
		if userInput < minimumNumber {
			minimumNumber = userInput
		}

	}
	fmt.Printf("The maximum number is %d\n", maximumNumber)
	fmt.Printf("The minimum number is %d\n", minimumNumber)
	fmt.Printf("The sum of both number is %d\n", maximumNumber+minimumNumber)
	fmt.Print("the first number is :\n", firstNumber)
	fmt.Print("the last number is :\n", lastNumber)
	fmt.Printf("the sum pf first and last number is %d\n", firstNumber+lastNumber)
}
