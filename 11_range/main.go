package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8}
	// s:=0
	// for _, num := range nums{
	// 	s+=num
	// 	fmt.Println(s)
		
	// }

	m:=map[string]string{"name":"amish","field":"devops"}
	for index , v := range m{
		fmt.Println(index,v)
	}


	//range in string

	for i, c := range "golang"{
		fmt.Println(i,c)
	}
}
