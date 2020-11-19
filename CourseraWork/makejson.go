package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Addr Address
}
type Address struct {
	Street string
	City   string
	State  string
	Zip    int
}

var (
	firstName        string
	lastName         string
	userStreetNumber string
	userStreetName   string
	userStreetType   string
	userCity         string
	userState        string
	userZip          int
)

func main() {

	fmt.Println("SPACES Not ALLOWED / Press ENTER For Each Input")

	fmt.Println("Please Enter Your First Name")
	_, _ = fmt.Scan(&firstName)

	fmt.Println("Please Enter Your Last Name")
	_, _ = fmt.Scan(&lastName)

	var userName = firstName + " " + lastName

	fmt.Println("Please Enter Your Street Address Number:")
	_, _ = fmt.Scan(&userStreetNumber)

	fmt.Println("Please Enter Your Street Address Name ONLY:")
	_, _ = fmt.Scan(&userStreetName)

	fmt.Println("Is Your Address ( ST / DR / BLVD / HWY:")
	_, _ = fmt.Scan(&userStreetType)

	var userStreet = userStreetNumber + " " + userStreetName + " " + userStreetType

	fmt.Println("Please Enter Your City: ")
	_, _ = fmt.Scanf("%s", &userCity)

	fmt.Println("Please Enter Your State: ")
	_, _ = fmt.Scanf("%s", &userState)

	fmt.Println("Please Enter Your Zip Code: ")
	_, _ = fmt.Scanf("%d", &userZip)

	var person = User{
		Name: userName,
		Addr: Address{
			userStreet,
			userCity,
			userState,
			userZip,
		},
	}
	makeItJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User Information: \n")
	fmt.Println(string(makeItJSON))

}
