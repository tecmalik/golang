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

		digit := int(binaryInput[counter] - 48)

		decimal = decimal + (float64(digit) * (math.Pow(2, float64(count))))
		count++
	}
	fmt.Printf("decimal = %d\n", int(decimal))
}
