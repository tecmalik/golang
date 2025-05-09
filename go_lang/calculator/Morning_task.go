package main

import "fmt"

func main() {
	valueString := "abcde"
	stringGoal := "cdeab"

	fmt.Print(check(valueString, stringGoal))
}
func check(firstValue string, goal string) bool {
	if len(firstValue) != len(goal) {
		return false
	}
	var digit = (len(firstValue) - 1) / 2
	finalValue := firstValue[digit:(len(firstValue))]
	finalValue += firstValue[1:digit]
	if finalValue != goal {
		return false
	}
	fmt.Print(finalValue + "23")
	return true
}
