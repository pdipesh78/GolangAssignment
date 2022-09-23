package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "log"
	// "net/http"
)

type User struct {
	Id       int64 `json:"userId"`
	Password int64 `json:"Password"`
}

var selectedOption string
var userId int64
var password int64

func login() {
	var option string
	fmt.Println("Please enter your user id: ")
	fmt.Scan(&userId)
	fmt.Println("Please enter your password: ")
	fmt.Scan(&password)
	user := User{Id: userId, Password: password}

	f, err := os.Open("user.json")
	if err != nil {
		fmt.Println("Error while validating user")
	}
	defer f.Close()

	var userData User
	err = json.NewDecoder(f).Decode(&userData)
	if err != nil {
		fmt.Println("Error in reading file")
	}

	if userData.Id == user.Id && userData.Password == user.Password {
		fmt.Println("User authenticated successfully...")
		fmt.Println("1.Withdraw money.\n2.Deposit money.\n3.Request balance.\n4.Quit the program.")
		fmt.Scan(&option)
	} else {
		fmt.Println("\nIncorrect UserId or Password \n")
		main()
	}
	switch option {
	case "1":
		fmt.Println("Withdraw money")
		main()
	case "2":
		fmt.Println("Deposit money")
		main()
	case "3":
		fmt.Println("Check balance")
		main()
	case "4":
		fmt.Println("exiting...")
		os.Exit(3)
	}
}

func createAccount() {
	newUser := User{}
	fmt.Println("enter userId:")
	fmt.Scan(&newUser.Id)
	fmt.Println("enter password:")
	fmt.Scan(&newUser.Password)
	fmt.Println("User Created..", newUser)
}

func main() {
	fmt.Println("Hi! Welcome to Mr. Jagadeesh ATM Machine! \n\nPlease select an option from the menu below: ")
	fmt.Println("\nl -> Login \nc -> Create New Account \nq -> Quit")
	fmt.Scan(&selectedOption)
	fmt.Println("You have selected: ", selectedOption)

	switch selectedOption {
	case "l":
		login()

	case "c":
		createAccount()
	case "q":
		os.Exit(3)
	default:
		fmt.Println("Please enter valid option\n")
		main()
	}

}
