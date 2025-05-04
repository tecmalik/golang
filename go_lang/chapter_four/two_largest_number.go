package main

import "fmt"

func main() {
	fmt.Print("Enter 10 numbers: ")
	var userInput int32
	var largestNumber int32
	var secondLargestNumber int32
	for count := 0; count < 10; count++ {

		fmt.Printf("Enter number %d number out of 10 :\n", count+1)
		_, err := fmt.Scan(&userInput)
		if count == 0 {
			largestNumber = userInput
			secondLargestNumber = largestNumber
		}
		if err != nil {
			println("Error reading input")
		}
		if userInput > largestNumber {
			secondLargestNumber = largestNumber
			largestNumber = userInput
		}
		if userInput > secondLargestNumber && userInput < largestNumber {
			secondLargestNumber = userInput
		}

	}
	fmt.Printf("the Largest number is %d \n", largestNumber)
	fmt.Printf("the Second Largest number is %d", secondLargestNumber)

}
