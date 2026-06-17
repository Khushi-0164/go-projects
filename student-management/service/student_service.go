package service

import (
	"fmt"
	"student-management/models"
	"student-management/storage"
)

type StudentService struct {
	Storage storage.JSONStorage
}

func (s StudentService) AddStudent(name string, age int) error {
	students, err := s.Storage.Load()

	if err != nil {
		return err
	}

	maxID := 0

	for _, student := range students {
		if student.ID > maxID {
			maxID = student.ID
		}
	}

	newID := maxID + 1
	student := models.Student{
		ID:   newID,
		Name: name,
		Age:  age,
	}
	students = append(students, student)
	return s.Storage.Save(students)
}

func (s StudentService) ListStudents() ([]models.Student, error) {
	return s.Storage.Load()
}

func (s *StudentService) DeleteStudent(id int) error {
	students, err := s.Storage.Load()
	if err != nil {
		return err
	}
	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return s.Storage.Save(students)
		}
	}
	return fmt.Errorf("student with id %d not found", id)

}
func (s *StudentService) UpdateStudent(id int, name string, age int) error {
	students, err := s.Storage.Load()
	if err != nil {
		return err
	}
	for i := range students {
		student := &students[i]
		if student.ID == id {
			student.Name = name
			student.Age = age
			return s.Storage.Save(students)
		}
	}
	return fmt.Errorf("student with id %d not found", id)
}
