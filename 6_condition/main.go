package main

import "fmt"

var age int

func main() {
	fmt.Println("Enter your age")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("adult")
	} else if age <= 0 {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Teenager")
	}

	//logical condition
	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("User is authorized")
	}else{
		fmt.Println("Not Authorized")
	}

	//declaring variable inside if
	
	if hasRole := false; hasRole == true{
		fmt.Println("User has role")
	}else{
		fmt.Println("User dont have role.")
	}

	// go doesnt have any ternary operators in go version 1.25
}
