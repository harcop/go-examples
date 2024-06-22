package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Print("Write ", i, " as ")

	switch i {
	case 1: 
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3: 
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("I don't know me", t)
		}
	}

	whatIam(true)
	whatIam(2)
}
