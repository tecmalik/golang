package main

import "fmt"

func main() {
	var milesDriven int
	var gallonsUsed int
	var milesPerGallon float64

	for gallonsUsed != -1 {

		fmt.Print("Enter the value of miles_driven or -1 to exit: ")
		_, err := fmt.Scan(&milesDriven)
		if err != nil {
			return
		}
		if milesDriven == -1 {
			break
		}

		fmt.Print("Enter the value of gallons_used or -1 tp exit: ")
		_, err = fmt.Scan(&gallonsUsed)
		if err != nil {
			return
		}

		if gallonsUsed != -1 {
			milesPerGallon = float64(milesDriven) / float64(gallonsUsed)
			fmt.Printf("calculated miles per gallon is : %f\n: ", milesPerGallon)
		}
	}
	print("Thank you")
}
