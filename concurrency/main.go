package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutine()
	channels()
	raceCondition()
	mutex()
}

func goroutine() {
	fmt.Println("\n\ngoroutine()")
	for i := 0; i < 5; i++ {
		go fmt.Println("Hello")
		go fmt.Println("World")
	}

	time.Sleep(time.Second * 1)
}

func channels() {
	fmt.Println("\n\nchannels()")
	ch := make(chan int)

	go func() {
		time.AfterFunc(time.Second*1, func() {
			ch <- 1
		})
		close(ch)
	}()

	fmt.Println("Do some work here")
	// waiting value from another goroutine
	val := <-ch
	fmt.Printf("Got %v from other goroutine\n", val)
}

func raceCondition() {
	fmt.Println("\n\nraceCondition()")
	counter := 0
	n := 1000
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println("Is counter equals to", n, "?")
	fmt.Println(counter == n)
	if counter != n {
		fmt.Println("race condition!!!")
	}
}

func mutex() {
	fmt.Println("\n\nmutex()")
	counter := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("counter :", counter)
	fmt.Println("is counter equals to 1000 ?", counter == 1000)
	if counter == 1000 {
		fmt.Println("no race condition!!!")
	}
}
