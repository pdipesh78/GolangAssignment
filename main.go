package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserData struct {
	UID      int64 `json:"username"`
	Password int64 `json:"password"`
}

var selectedItem string
var username int64
var password int64

func userLogin() {
	var option string
	fmt.Println("Please Enter Your User ID: ")
	fmt.Scan(&username)
	fmt.Println("Please Enter Your Password: ")
	fmt.Scan(&password)
	userdata := UserData{UID: username, Password: password}

	f, err := os.Open("userdata.json")
	if err != nil {
		fmt.Println("Error while validating user")
	}
	defer f.Close()

	var userData UserData
	err = json.NewDecoder(f).Decode(&userData)
	if err != nil {
		fmt.Println("Error in reading file")
	}

	if userData.UID == userdata.UID && userData.Password == userData.Password {
		fmt.Println("User authenticated successfully...")
		fmt.Println("1.Withdraw money.\n2.Deposit money.\n3.Request balance.\n4.Quit the program.")
		fmt.Scan(&option)
	} else {
		fmt.Println("\n Incorrect UserId or Password")
		main()
	}
	switch option {
	case "1":
		fmt.Println("Please withdraw money")
		main()
	case "2":
		fmt.Println("Please deposit money")
		main()
	case "3":
		fmt.Println("Please check balance")
		main()
	case "4":
		fmt.Println("exiting...")
		os.Exit(3)
	}
}

func createAccount() {
	newUser := UserData{}
	fmt.Println("Please enter userId:")
	fmt.Scan(&newUser.UID)
	fmt.Println("Please enter password:")
	fmt.Scan(&newUser.Password)
	fmt.Println("User Created..", newUser)
}

func main() {
	fmt.Println("Hi! Welcome to Mr. Mahesh ATM Machine! \n\nPlease select an option from the menu below: ")
	fmt.Println("\nl -> Login \nc -> Create New Account \nq -> Quit")
	fmt.Scan(&selectedItem)
	fmt.Println("You have selected: ", selectedItem)

	switch selectedItem {
	case "l":
		userLogin()
	case "c":
		createAccount()
	case "q":
		os.Exit(3)
	default:
		fmt.Println("Please Enter Valid Option")
		main()
	}
}
