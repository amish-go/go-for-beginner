package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("IN changeNum", num)
// }

// by refrence
func changeNum(num *int){
	*num = 5 
	fmt.Println("IN changeNum", *num)
}

func main() {
	num := 1
	changeNum(&num)

	fmt.Println("Memory address of num is",&num)
	

	fmt.Println("After change",num )
}
