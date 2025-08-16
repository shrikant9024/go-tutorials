package main

import "fmt"

type OrderStatus int

const (
	Recieved OrderStatus = iota
	Confirmed
	Prepared
	Delivered
	Completed
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order stats to", status)
}

func main() {

	changeOrderStatus(Recieved)

}
