package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	Fname string
	Lname string
}

var fileChecker string

func main() {

	fmt.Println("Please Enter File Name with file extension(.txt): ")
	_, err := fmt.Scan(&fileChecker)
	if err != nil {
		fmt.Println(err)
	}

	if fileChecker == "names.txt" {
		file, err := os.Open(fileChecker)
		if err != nil {
			fmt.Println(err)
		}
		read := bufio.NewReader(file)

		sliceStru := make([]Person, 0, 0)
		for {
			line, _, err := read.ReadLine()
			if err == io.EOF {
				break
			}
			word := strings.Split(string(line), " ")
			people := Person{
				Fname: word[0],
				Lname: word[1],
			}
			sliceStru = append(sliceStru, people)
		}

		for _, item := range sliceStru {
			fmt.Println(item)
		}

	} else {
		fmt.Println("Wrong File Name")
		os.Exit(1)
	}

}
