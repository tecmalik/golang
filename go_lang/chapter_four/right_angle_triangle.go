package main

import "fmt"

func main() {
	var base int
	fmt.Print("Enter a base length between 1 to 10: ")
	_, err := fmt.Scan(&base)
	if err != nil {
		fmt.Println("Error reading input")
	}
	for triangle := 0; triangle < base; triangle++ {
		for count := 0; count < triangle; count++ {
			fmt.Print("*")
		}
		println("p")
	}

}
