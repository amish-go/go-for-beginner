package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //
}
//reciver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	orderList := order{
		id:     "1",
		amount: 66.6,
		status: "paid",
	}

	orderList2 := order{
		id: "2",
		amount: 30.0,
		status: "unpaid",
	}

	//to add field in the struct

	orderList.createdAt = time.Now()

	//to get a field
	fmt.Println(orderList.amount)
	//to print whole struct
	fmt.Println("Order struct", orderList)

	//to change value of struct 
	orderList2.status="delivered"
	fmt.Println("Order struct", orderList2)

	
	orderList.changeStatus("confirmed")
	fmt.Println(orderList)


}
