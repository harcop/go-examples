package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 32
	m["k2"] = 13

	fmt.Println("maps", m)
	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("maps", m)

	clear(m)
	fmt.Println("maps", m)

	vv, prs := m["k2"]
	fmt.Println("prs", prs)
	fmt.Println("_", vv)

	n1 := make(map[string]int{"foo": 1, "bar": 2})
	n2 := make(map[string]int{"foo": 1, "bar": 2})
	if maps.Equal(n1, n2) {
		fmt.Println("n1 == n2")
	}
}
