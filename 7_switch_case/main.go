package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 7
	switch i {
	case 1:
		fmt.Println("one") // if i is 1
	case 2:
		fmt.Println("two") //if i is 2
	case 3:
		fmt.Println("three") //if i is 3
	default: // if nothing is true
		fmt.Println("seven")
	}


	// multiple condition switch
	switch time.Now().Weekday(){ // using standerd time package library
	case time.Saturday , time.Sunday:
		fmt.Println("Its weekend!")
	default:
		fmt.Println("Its weekday!")
	}

	// type switch
	whoAmI := func(j interface{}){
		switch t := j.(type){
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Other",t)
		}
	}

	whoAmI(true)
}
