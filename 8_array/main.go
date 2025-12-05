package main

import "fmt"

func main() {
	var nums [4]int
	// to get the length of array
	fmt.Println(len(nums))
	// to add elements
	nums[0] /*0 is the index*/ = 1
	fmt.Println(nums[1])
	//to print whole array
	fmt.Println(nums)

	var vals[4]bool
	vals[2]=true
	fmt.Println(vals) //the default is false in case of bool and 0 in type of int

	var names[3]string
	names[1]="amish"
	fmt.Println(names) //it prints empty string

	//to declare it in single line
	items := [3]string{"pencil","bag","pen"}
	fmt.Println(items)


	//to make 2d array
	foods := [2][2]string{{"apple","banana"},{"noodles","chips"}}
	fmt.Println(foods)
}

// why to use array?
// - fixed size , -- memory optimization , --constant time access