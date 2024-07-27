package main

import "fmt"

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

// stringer is like toString in Java

// Person type
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Implementing the Stringer interface for Person
func (p Person) String() string {
    return fmt.Sprintf("%s %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

func main() {
    person := Person{FirstName: "John", LastName: "Doe", Age: 30}
    fmt.Println(person)  // This will call the String() method automatically


		fmt.Println(Aspirin)
		fmt.Println(Placebo)
		fmt.Println(Acetaminophen)
		fmt.Println(Paracetamol)
}
