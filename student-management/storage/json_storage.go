package storage

import (
	"encoding/json"
	"os"
	"student-management/models"
)

type JSONStorage struct {
	FileName string
}

func (j JSONStorage) Load() ([]models.Student, error) {
	data, err := os.ReadFile(j.FileName)
	if err != nil {
		return nil, err
	}
	var students []models.Student
	err = json.Unmarshal(data, &students)
	return students, err
}

func (j JSONStorage) Save(students []models.Student) error {

	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(j.FileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
