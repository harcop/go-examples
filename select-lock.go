package main

import (
	"fmt"
)

func main2() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("Starting in Goroutine")
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <- ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()
	
	fmt.Println("Starting in Main")
	inMain := 2
	ch2 <- inMain
	fromGoroutine := <-ch1
	fmt.Println("main:", inMain, fromGoroutine)
} 

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("Starting in Goroutine")
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <- ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()
	
	fmt.Println("Starting in Main")
	inMain := 2
	var fromGoroutine int
	select {
	case ch2 <- inMain:
	case fromGoroutine = <-ch1:
	}
	fmt.Println("main:", inMain, fromGoroutine)
} 
