package main

import "fmt"

func main() {
	fmt.Print("Enter a number to calculate estimates the value \n of the mathematical constant e :")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
	}
	var factorial = 1
	eConstant := 1.0
	count := 1
	counter := 1
	for count < number {
		eConstant += 1 / float64(factorial)
		factorial = 1
		counter++
		count++

		for counter < 0 {
			factorial *= counter
		}
	}
	fmt.Printf("e constant is %f\n", eConstant)

}
