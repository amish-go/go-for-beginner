package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _,num:= range nums {
		total+=num
	}
	return total
}

func main() {
	fmt.Println(sum(2,3,4,2,5,32,4234,23,23))


	//to send value from slice
	nums := []int{3,4,5,6}
	fmt.Println(sum(nums...))
}
