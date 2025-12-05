package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	m := make(map[string]string)

	// setting element
	m["name"] = "Amish"
	m["field"] = "DevOPS"


	//get an element
	fmt.Println(m["name"])
	fmt.Println(m["field"])
	fmt.Println(m["email"]) //if key value doesn't exist in map then it return 0 values


	new_map := make(map[int]string)
	new_map[1]="hasPerm=Admin"
	fmt.Println(new_map[1])

	//to delete element from map
	delete(m,"field")
	fmt.Println(m)

	//to delete whole map
	clear(m)
	fmt.Println(m)

	//to check if an element is in map or not
	ak := map[string]int{"price":100,"phones":2}
	_, ok :=ak["phones"]
	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

	//to check two maps are equal or not
	ak1 := map[string]int{"price":100,"phones":2}
	ak2 := map[string]int{"price":100,"phones":2}

	fmt.Println(maps.Equal(ak1,ak2))

}
