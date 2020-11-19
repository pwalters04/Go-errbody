package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbers := make([]int, 3)
	var string_number string
	var max = 100
	var init = 0

	for init <= max {

		fmt.Printf("Please Enter An Integer To Append ( Enter x to exit): ")
		_, err := fmt.Scanf("%v", &string_number)
		if err != nil {
			fmt.Println(err)
		}

		if string_number == "x" {
			break
		}
		var number, _ = strconv.Atoi(string_number)

		numbers = append(numbers, number)
		sort.Ints(numbers)
		fmt.Printf("len=%d slice=%v\n", len(numbers), numbers)

		fmt.Println("*****************")
		init = +2
	}
	fmt.Println("Thank You For Adding Numbers.")
	fmt.Println(" Sorted Final Slice: ", numbers)
	fmt.Println(" Exiting")

}
