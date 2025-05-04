package main

import "fmt"

func main() {
	fmt.Println("Enter a number to simulate the sum of series from 1....n ")
	var number = 0
	var sumOfSeries int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("error:", err)
	}
	for count := 1; count <= number; count++ {
		sumOfSeries = number + count
	}
	fmt.Println("Sum of series is:", sumOfSeries)
}
