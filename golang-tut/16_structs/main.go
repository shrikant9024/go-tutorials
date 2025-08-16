package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    int
	status    string
	createdAt time.Time
	customer
}

// func newOrder(id string, amount int, status string) *order {

// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder

// }

// func (o *order) changeStatus(status string) {

// 	o.status = status

// 	fmt.Println(status)
// }

func main() {

	//instance
	// var myOrder order =

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50,
	// 	status: "Recieved",
	// }
	// myOrder.createdAt = time.Now()

	// 	fmt.Println("Printing the object", myOrder)

	// 	// myOrder.changeStatus("confirmed")
	// 	fmt.Println(myOrder.status)

	//using constructer
	// myOrder := newOrder("1", 300, "approved")

	// fmt.Println(myOrder)
	// fmt.Println(myOrder.id)

	// //if u want to use one time struct

	// language := struct {
	// 	name string
	// 	flag bool
	// }{"golang", true}

	// fmt.Println(language)

	customer1 := customer{
		name:  "John doe",
		phone: "988888888",
	}

	myOrder := order{
		id:        "1",
		amount:    50,
		createdAt: time.Now(),
		status:    "approved",
		customer:  customer1,
	}

	myOrder.customer.name = "strr"

	fmt.Println(myOrder.customer)

}
