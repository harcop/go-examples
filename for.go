package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { //while loop
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range [3]int{} {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("Hello", i)
		break
	}

	for n := range [6]int{} {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
