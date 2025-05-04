package main

import (
	"fmt"
)

func main() {
	for count := 1; count < 6; count++ {
		fmt.Printf("%d\t%d\t%d\t%d \n", count, count*count, count*count*count, count*count*count*count)
	}

}
