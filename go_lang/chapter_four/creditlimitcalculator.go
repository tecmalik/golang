package main

import (
	"fmt"
)

func main() {
	fmt.Println("Credit Limit Calculator")
	fmt.Println("======================")
	fmt.Print("Enter account number: \n")
	var accountNumber string
	_, err := fmt.Scan(&accountNumber)
	if err != nil {
		fmt.Println("Error reading input")
	}
	var beginningBalance int32
	fmt.Print("Enter beginning balance :\n")
	_, err = fmt.Scan(&beginningBalance)
	if err != nil {
		fmt.Println("Error reading input")
	}
	fmt.Print("Enter ending balance :\n")

}
