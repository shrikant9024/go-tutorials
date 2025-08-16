package main

import "fmt"

type paymenter interface {
	pay(amount int)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount int) {
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount int) {
	fmt.Println("sending money through razorpay ", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount int) {
	fmt.Println("Paying through fakepayment mehthod", amount)
}
func main() {
	newPayment := payment{
		gateway: razorpay{},
	}

	newFakePayemnt := payment{
		gateway: fakepayment{},
	}

	newPayment.makePayment(100)
	newFakePayemnt.makePayment(200)

}

//interface is an contract who so ever signs it should follow the method it provides , so here we are implementing two method for two structs, razorpay and fakepayment
//now we have a payementer type and , gateway is payementer type in payment struct so , two pay methods we have in razorpay and fake payement signed the contract as they have pay methods
//and in main function we created the instance of the razorpay and fake payement and used the gateway  to implement particular payemnt for razorpay and fakepayment

//implements polymorphism
