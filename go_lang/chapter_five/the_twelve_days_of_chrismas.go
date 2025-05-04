package main

import "fmt"

func main() {
	var day = 0
	for day != -1 {
		fmt.Print(" enter Day from 1 to 12 or 13 to quit : ")
		_, err := fmt.Scan(day)
		if err != nil {
			fmt.Println("Error reading input")
		}
		switch day {
		case 1:
			fmt.Println("On the first day of Christmas\n\tMy true love gave to me\n\tA partridge in a pear tree.")
		case 2:
			fmt.Println("On the second day of Christmas\n\tMy true love gave to me\n\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 3:
			fmt.Println("On the third day of Christmas\n\tMy true love gave to me\n\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 4:
			fmt.Println("On the fourth day of Christmas\n\tMy true love gave to me\n\tFour calling birds\n\tThree French hens\n\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 5:
			fmt.Println("On the fifth day of Christmas\n\tMy true love gave to me\n\tFive golden rings\n\tFour calling birds\n\tThree French hens\n\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 6:
			fmt.Println("On the sixth day of Christmas\n\tMy true love gave to me\n\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 7:
			fmt.Println("On the seventh day of Christmas\n\tMy true love gave to me\n\tSeven swans a-swimming,\n\t\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 8:
			fmt.Println("On the eighth day of Christmas\n\tMy true love gave to me\n\tEight maids a-milking,\n\t\tSeven swans a-swimming,\n\n\t\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 9:
			fmt.Println("On the ninth day of Christmas\n\tMy true love gave to me\n\tNine ladies dancing,\n\t\tEight maids a-milking,\n\t\tSeven swans a-swimming,\n\t\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 10:
			fmt.Println("On the tenth day of Christmas\n\tMy true love gave to me\n\tTen lords a-leaping,\n\t\tNine ladies dancing,\n\t\tEight maids a-milking,\n\t\tSeven swans a-swimming,\n\t\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 11:
			fmt.Print("On the eleventh day of Christmas\n\tMy true love gave to me Eleven pipers piping,\n\t\tTen lords a-leaping,\n\t\tNine ladies dancing,\n\t\tEight maids a-milking,\n\t\tSeven swans a-swimming,\n\t\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 12:
			fmt.Println("On the twelfth day of Christmas\n\tMy true love gave to me\n\tTwelve drummers drumming,\n\t\tEleven pipers piping,\n\t\tTen lords a-leaping,\n\t\tNine ladies dancing,\n\t\tEight maids a-milking,\n\t\tSeven swans a-swimming,\n\t\tSix geese a-laying,\n\t\tFive golden rings,\n\t\tFour calling birds,\n\t\tThree French hens,\n\t\tTwo turtle doves\n\tAnd a partridge in a pear tree.")
		case 13:
			day = -1
		default:
			fmt.Println("invalid input")
		}
	}
}
