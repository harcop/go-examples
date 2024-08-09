package main

type Person struct {
	FirstName string
	LastName string
	Age int
}

func MakePerson(firstName, lastName string, age int) Person {
	person := Person {
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
	return person
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
		person := Person {
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
	return &person
}

func main() {
	MakePerson("Bomber", "Blower", 24)
	MakePersonPointer("Bomber", "Blower", 24)
}
