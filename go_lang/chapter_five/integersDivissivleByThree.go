package main

import "fmt"

func main() {
	for count := 1; count < 31; count++ {
		if count%3 == 0 {
			fmt.Printf("%d is divisible by 3\n", count)
		}
	}

}
