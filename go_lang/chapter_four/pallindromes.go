package main

import "fmt"

func main() {
	fmt.Println("enter 5 digit number")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	if len(input) > 5 || len(input) < 5 {
		fmt.Println("input not 5 digit number.")
		return
	}
	if input[0] == input[4] && input[1] == input[3] {
		fmt.Printf(" %s is a palladium", input)
	}
}
