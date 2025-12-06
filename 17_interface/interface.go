package main

import "fmt"

type paymentGateway interface {
	getAmount() int
	getRemarks() string
}

//eSewa Struct

type eSewa struct {
	amount  int
	remarks string
}

func (e eSewa) getAmount() int {
	return e.amount
}

func (e eSewa) getRemarks() string {
	return e.remarks
}

//khalti struct

type khalti struct {
	amount  int
	remarks string
}

func (k khalti) getAmount() int {
	return k.amount
}

func (k khalti) getRemarks() string {
	return k.remarks
}

type cash struct {
	amount int
	remarks string
}


func (c cash) getAmount() int {
	return c.amount
}

func (c cash) getRemarks() string {
	return c.remarks
}

func main() {
	py1 := eSewa{amount: 150, remarks: "For study materials."}
	py2 := khalti{amount: 200, remarks: "For foods."}
	py3 := cash{amount: 333, remarks: "Service Charge"}

	var gateway int
	fmt.Println("Please choose your payment method: \n 1.eSewa \n 2.Khalti \n Cash")
	fmt.Scanln(&gateway)
	
	var choice paymentGateway
	switch gateway{
	case 1:
		choice = py1
	case 2:
		choice = py2

	case 3:
		choice = py3
	default:
		fmt.Println("Invalid payment gateway choosen")
		return
	}

	fmt.Println("You paid :",choice.getAmount())
	fmt.Println("Remarks :",choice.getRemarks())

}
