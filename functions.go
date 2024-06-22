package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	r := plus(1, 2)
	fmt.Println("1+2 =", r)

	res := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals();
	fmt.Println(a, b)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
