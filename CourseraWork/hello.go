package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	//testing("me", 2020,04,30)

	var number string
	fmt.Printf("Please Enter An Integer To Append ( Enter -1 to exit): ")
	_, err := fmt.Scanf("%v", &number)
	if err != nil {
		fmt.Println(err)
	}
	var charChecker, _ = strconv.Atoi(number)
	fmt.Println("Is the a Number:", charChecker)
}

func testing(name string, year int, month int, day int) {

	allTogether := "release." + strconv.Itoa(year) + strconv.Itoa(month) + strconv.Itoa(day)
	fmt.Println("Does this work: ", allTogether)

}
