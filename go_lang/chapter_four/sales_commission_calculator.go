package main

import "fmt"

func main() {
	var itemValue float32
	var totalValue float32
	count := 1
	for itemValue != -1 {
		fmt.Printf("Enter item %d value or enter -1 to exit: ", count)
		_, err := fmt.Scan(&itemValue)
		if err != nil {
			fmt.Println("Error reading input")
		}
		count++
		if itemValue != -1 {
			totalValue = +itemValue
		}
	}

	salesCommission := 200 + ((totalValue / 100) * 9)
	fmt.Println("======================")
	fmt.Println("your total sales commission is : ", salesCommission)
	fmt.Println("======================")
	fmt.Print("Thank you")
}
