package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{name: "Toluwap", age: 27})

	fmt.Println(newPerson("John"))

	s := person{name: "Toluwap", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 51

	fmt.Println(sp.age)
	fmt.Println(s.age)
}
