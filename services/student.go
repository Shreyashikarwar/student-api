package services

import (
	"student-api/models"
	"student-api/repositories"
)

func GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := repositories.GetAllStudents(&students)
	return students, err
}

func CreateStudent(student models.Student) error {
	return repositories.CreateStudent(&student)
}

func GetStudentByID(id int) (models.Student, error) {
	var student models.Student
	err := repositories.GetStudentByID(&student, id)
	return student, err
}

func UpdateStudent(student models.Student) error {
	return repositories.UpdateStudent(&student)
}

func DeleteStudent(id int) error {
	return repositories.DeleteStudent(id)
}
