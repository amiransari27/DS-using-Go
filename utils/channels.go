package utils

import (
	"fmt"
	"sync"
	"time"
)

func SendHello() {

	msgCh := make(chan string)

	go func(localCh chan<- string) {
		fmt.Println("Sender: Sending message ")

		localCh <- "Hello from goroutine"

		fmt.Println("Sender: Message recieved by someone ")
	}(msgCh)

	go func(recCh <-chan string) {

		fmt.Println("Reciewer: Ready to recieve")

		msg := <-recCh

		fmt.Println("Reciewer: got messge ", msg)
	}(msgCh)

	time.Sleep(time.Second * 1)
}

func SendHelloAndWorld() {
	msgCh := make(chan string, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Sender : Send hello")
		msgCh <- "Hello"
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Sender : Send World")
		msgCh <- "World"
	}()

	fmt.Println("reciewer: ready")

	go func() {
		wg.Wait()
		close(msgCh)
	}()

	for msg := range msgCh {

		fmt.Println("Message recieved ", msg)
	}

	time.Sleep(time.Second * 1)

}
