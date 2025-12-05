package main
import "fmt"
func main() {
	a  := 20 //set value in code
	var b int
	fmt.Println("Input value of b :")
	fmt.Scanln(&b) //ask user for value
	s := a+b
	fmt.Println("Sum is:",s)
}