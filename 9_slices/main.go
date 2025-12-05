package main

import (
	"fmt"
	"slices"
)

// slices are dynamic arrays
func main() {
	// uninitialized slice is nil
	var nums []int
	// fmt.Println(nums)
	fmt.Println(len(nums))

	var item = make([]int,2,3) // initial size of the slice is 2 and 3 is its capacity
	fmt.Println(item)
	fmt.Println(cap(item)) //the max capacity of slice is 3 cap--> to check the capacity
	//to add element is slices
	item = append(item, 1)
	item = append(item, 2)
	item = append(item, 3)
	item = append(item, 4)
	fmt.Println(item)
	fmt.Println(cap(item)) // automatically increase the capacity
	fmt.Println(len(item))

	//indexing
	item[1]=3
	item[0]=2
	fmt.Println(item)

	// to copy slice
	var copy_item = make([]int,len(item))
	copy(copy_item,item)
	fmt.Println(copy_item)

	// slice operator
	var ope = []int{1,2,3}
	fmt.Println(ope[0:2])//it will retrurn value from 0 index to second index  (it include itself)
	fmt.Println(ope[:2]) 


	//slice package
	var product1 = []int{1,2}
	var product2 = []int{1,2}
	fmt.Println(slices.Equal(product1, product2))

	//2d slice
	var twod = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(twod)
}
