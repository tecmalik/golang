package main

import (
	"fmt"
	"math"
)

func main() {
	var binaryInput string
	fmt.Println("enter binary number (1 and 0's) :")
	_, err := fmt.Scan(&binaryInput)
	if err != nil {
		fmt.Println("Error reading input")
	}
	count := 0
	decimal := 0.0
	for counter := len(binaryInput) - 1; counter >= 0; counter-- {
		decimal = +(float64(binaryInput[counter]) * (math.Pow(2.0, float64(count))))
		count++
	}
	fmt.Printf("decimal = %d\n", int(decimal))
}
