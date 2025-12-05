package main

import "fmt"

const age int = 30 //constant and variables can be declared outside the variables
func main() {
	const pi float64 = 3.14
	//pi = 737 cant change value because it is an constant
	fmt.Println(pi)
	//multiple constanr (grouping)
	const (
		port = 5050
		host = "localhost"
	)
	fmt.Println(port,host)


}
