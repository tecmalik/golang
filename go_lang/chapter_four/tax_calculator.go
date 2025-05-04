package main

import "fmt"

func main() {
	fmt.Print("# citizen tax calculator per earnings")
	for count := 0; count < 3; count++ {
		var tax float32
		fmt.Printf("Enter ernings of citizen %d: \n", count+1)
		var earnings float32
		_, err := fmt.Scan(&earnings)
		if err != nil {
			fmt.Println("Error reading input")
		}
		if earnings <= 30000 {
			tax = (earnings / 100) * 15
		}
		if earnings > 30000 {
			tax = (earnings / 100) * 20
		}
		fmt.Printf("%d citizen tax is %f \n", count+1, tax)
	}
}
