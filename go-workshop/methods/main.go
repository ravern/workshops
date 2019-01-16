package main

type person struct {
	name string
	age  int
}

// Returns the date of birth of the given person.
func dateOfBirthFunction(p person, currentYear int) int {
	return currentYear - p.age
}

// Returns the date of birth of the given person.
func (p person) dateOfBirthMethod(currentYear int) int {
	return currentYear - p.age
}

func main() {
	myPerson := person{name: "John", age: 20}
	dateOfBirthFunction(myPerson, 2019)
	myPerson.dateOfBirthMethod(2019)
}
