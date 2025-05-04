package main

import "fmt"

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
	fmt.Print("Enter beginning balance of the month :\n")
	_, err = fmt.Scan(&beginningBalance)
	if err != nil {
		fmt.Println("Error reading input")
	}
	fmt.Print("Enter ending balance :\n")
	var totalChargedItems int32
	fmt.Print("Enter total charged items of the month :\n")
	_, err = fmt.Scan(&totalChargedItems)
	if err != nil {
		fmt.Println("Error reading input")
	}
	var totalCredits int32
	fmt.Print("Enter total credits of the month :\n")
	_, err = fmt.Scan(&totalCredits)
	if err != nil {
		fmt.Println("Error reading input")
	}
	var allowedCreditLimits int32
	fmt.Print("Enter allowed credits of the month :\n")
	_, err = fmt.Scan(&allowedCreditLimits)
	if err != nil {
		fmt.Println("Error reading input")
	}

	newBalance := beginningBalance + totalChargedItems - totalCredits
	if newBalance < allowedCreditLimits {
		fmt.Println("You do not have enough credits on this month ")
		fmt.Println("new balance = ", newBalance)
		fmt.Println("allowed credit limit = ", allowedCreditLimits)
	} else {
		fmt.Println("You do not have enough credits on this month ")
		fmt.Println("new balance = ", newBalance)
		fmt.Println("allowed credit limit = ", allowedCreditLimits)

	}
}
