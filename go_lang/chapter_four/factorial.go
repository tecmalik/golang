package main

import "fmt"

func main() {
	fmt.Print("Enter a number to get is Factorial")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading input : ")
	}
	factorial := 1
	for number > 0 {
		factorial *= number
		number = number - 1

	}
	fmt.Println(factorial)
}
