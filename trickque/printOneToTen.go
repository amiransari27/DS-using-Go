package trickque

import (
	"fmt"
	"sync"
)

func PrintOneToTen() {

	evenCh := make(chan bool, 1)
	oddCh := make(chan bool, 1)

	var wg sync.WaitGroup

	oddCh <- true

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i += 2 {
			<-oddCh
			fmt.Println(i)
			evenCh <- true
		}

	}()

	go func() {
		defer wg.Done()

		for i := 2; i <= 10; i += 2 {
			<-evenCh
			fmt.Println(i)
			if i < 10 {
				oddCh <- true
			}
		}

	}()

	wg.Wait()
}
