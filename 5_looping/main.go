package main

import "fmt"

// in go there is only for loop
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// infinite loop
	// for {
	// 	println("1")
	// }
	
	// normal for loop
	for j:=0;j< 3; j++ {
		// break
		if i==2{ //learning if in next chapter
			continue // to skip one iterration
		}
		fmt.Println(j)
	}

	// range
	for k := range 10{
		fmt.Println(k)
	}
}
