package jonsephu

import "fmt"

type Person struct {
	No   int
	Next *Person
}

func AddPerson(no int) *Person {

	first := &Person{}
	current := &Person{}
	if no < 1 {
		fmt.Println("no的值不对")
		return first
	}

	for i := 1; i <= no; i++ {
		person := &Person{
			No: i,
		}
		if i == 1 {
			first = person
			current = person
			current.Next = first
		} else {
			current.Next = person
			current = person
			current.Next = first
		}
	}
	return first
}
