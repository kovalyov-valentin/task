package main

import "fmt"

func push(name string, ch chan string) {
	msg := "Hey, " + name 
	ch <- msg
}

func main() {
	ch := make(chan string) 

	go push("Moe", ch)
	go push("Larry", ch)
	go push("Curly", ch)

	fmt.Println(<-ch, <-ch, <-ch)
}