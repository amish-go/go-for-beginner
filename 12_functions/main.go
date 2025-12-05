package main

import "fmt"

// func add(a, b int) int {
// 	return a + b

// }

func getLanguages() (string, string , string , bool , int){
 	return "golang","java","c++",false,69
}




func main() {

	

	// result := add(2, 3)
	lang1 , lang2 , lang3 , final , _:= getLanguages()
	fmt.Println(lang1 , lang2 , lang3, final)


}
