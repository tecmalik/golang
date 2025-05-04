package main

import "fmt"

func main() {
	fmt.Print("Enter 10 numbers: ")
	var userInput int32
	var largestNumber int32
	for count := 0; count < 10; count++ {

		fmt.Printf("Enter number %d number out of 10 :\n", count+1)
		_, err := fmt.Scan(&userInput)
		if count == 0 {
			largestNumber = userInput
		}
		if err != nil {
			println("Error reading input")
		}
		if userInput > largestNumber {
			largestNumber = userInput
		}
	}
	fmt.Printf("the Largest number is %d", largestNumber)
}
