package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	rate := 0.05
	numberOfYears := 10
	fmt.Println("Year      Amount on Deposit ")
	for year := 0; year < numberOfYears; year++ {
		amount := principal * math.Pow(1+rate, float64(year))
		fmt.Printf("%4d%  20.2f\n", year, amount)
	}

}
