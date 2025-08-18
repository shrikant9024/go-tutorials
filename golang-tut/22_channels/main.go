package main

import (
	"fmt"
	// "time"
)

//sending

// func processNum(numchan chan int) {

// 	for num := range numchan {

// 		fmt.Println("Processing Number", num)
// 		time.Sleep(time.Second * 2)

// 	}
// }

// func sum(result chan int, a int, b int) {

// 	numResult := a + b
// 	result <- numResult

// }

// func task(done chan bool) {

// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("processing")

// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	for email := range emailChan {
// 		fmt.Println("sending email to ", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {

	// numchan := make(chan int)

	// go processNum(numchan)

	// numchan <- 5
	// for {
	// 	numchan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 1)

	// messageChan := make(chan string)

	// messageChan <- "ping" 		//blocking
	// msg := <-messageChan

	// fmt.Println(msg)

	// result := make(chan int)
	// go sum(result, 4, 5)

	// res := <-result
	// fmt.Println(res)

	// done := make(chan bool)
	// go task(done)

	// <-done

	//unbuffered channels

	// emailChan := make(chan string, 100)
	// done := make(chan bool)
	// go emailSender(emailChan, done)

	// // emailChan <- "example@gmail.com"
	// // emailChan <- "example2@gmail.com"

	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan)
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")
	// <-done

	//recieve only channel -> email <-chan string
	//send only channel ->  done chan-> bool

	//listening multiple channels
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "ping"
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recived data from channel 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("recived data from channel 2", chan2Val)
		}
	}

}
