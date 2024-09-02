package main

import "fmt"

func main() {
	v1 := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	ch1 := make(chan int, 10)
	go func () {
		for i := range(v1) {
			fmt.Println("Hello", i)
			ch1 <- i
		}
	}()
	fmt.Println(<-ch1)
	fmt.Println("bye bye")
	// processChannel(ch1)
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <- ch
			results <- process(v)
		}()
	}

	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
 return v * 2
}
