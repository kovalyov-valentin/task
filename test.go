package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(c *chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	i := 0

	for {
		time.Sleep(time.Duration(time.Millisecond * 10))
		*c <- i
		i++
		if i == 100 {
			close(*c)
			return
		}
	}
}

func readChan(ch *chan int, wg *sync.WaitGroup) {
wg.Add(1)
defer wg.Done()

loop:
	for {
		select {
		case i, ok := <-*ch:
			{
				if !ok {
				 break loop 
				}
				fmt.Println(i)

			}
		default:
			time.Sleep(time.Duration(time.Microsecond * 10))
		}
	}
}

func main() {
	c := make(chan int, 10)


	wg := sync.WaitGroup{}

	go greet(&c, &wg)
	for i := 0; i < 10;i++{
		go readChan(&c, &wg)
	}
	fmt.Println("main() started")

	time.Sleep(time.Duration(time.Second * 1))
	fmt.Println("main() sleap stop ")

	wg.Wait()
	fmt.Println("main() stopped")
}
