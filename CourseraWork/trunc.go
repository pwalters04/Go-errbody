package main

import (
	"fmt"
)

var userInput float64

func Trunc() int {

	fmt.Println("Please Input a Float Number and it will return an Int of your number")

	_, err := fmt.Scanf("%f", &userInput)
	if err != nil {
		fmt.Println("[ERROR] Input Error: ", err)
	}

	return int(userInput)
}
func main() {

	output := Trunc()
	println("Your Number: ", output)

}
