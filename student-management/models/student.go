package models

import "fmt"

type Student struct {
	ID   int
	Name string
	Age  int
}

func (s Student) Display() {

	fmt.Println("ID:", s.ID)
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)

}
