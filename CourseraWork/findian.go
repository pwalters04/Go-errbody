package main

import (
	"fmt"
	"strings"
)

var (
	letter_i = "i"
	letter_I = "I"
	letter_a = "a"
	letter_A = "A"
)

func Seek(your_word string) {
	var letter_a_check = strings.ContainsAny(your_word, letter_a)
	var letter_i_check = strings.ContainsAny(your_word, letter_i)
	var letter_A_check = strings.ContainsAny(your_word, letter_A)
	var letter_I_check = strings.ContainsAny(your_word, letter_I)

	if (letter_a_check == true && letter_i_check == true) || (letter_A_check == true || letter_I_check == true) {
		if strings.LastIndex(your_word, "n") > -1 || strings.LastIndex(your_word, "N") > -1 {
			fmt.Println("Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}

}

func main() {
	var userWord string

	fmt.Printf("Please Enter Your Word For The WordChecker: ")
	_, err := fmt.Scanf("%s", &userWord)
	if err != nil {
		fmt.Println(err)
	}
	Seek(userWord)

	fmt.Println("*******************************")

	fmt.Println("Now Running Automated Test Cases... ")

	var testStringsSlice = []string{"ian", "Ian", "iuiygaygn", "ihhhhhn", "xian"}
	for testItem := range testStringsSlice {
		fmt.Println("Test Word: ", testStringsSlice[testItem])
		Seek(testStringsSlice[testItem])
	}
}
