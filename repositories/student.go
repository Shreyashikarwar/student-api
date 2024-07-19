package repositories

import (
	"database/sql"
	"student-api/config"
	"student-api/models"
)

func GetAllStudents(students *[]models.Student) error {
	rows, err := config.DB.Query("SELECT id, name, age, grade FROM students")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Grade); err != nil {
			return err
		}
		*students = append(*students, student)
	}
	return nil
}

func CreateStudent(student *models.Student) error {
	query := "INSERT INTO students (name, age, grade) VALUES (?, ?, ?)"
	result, err := config.DB.Exec(query, student.Name, student.Age, student.Grade)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	student.ID = int(id)
	return nil
}

func GetStudentByID(student *models.Student, id int) error {
	query := "SELECT id, name, age, grade FROM students WHERE id = ?"
	return config.DB.QueryRow(query, id).Scan(&student.ID, &student.Name, &student.Age, &student.Grade)
}

func UpdateStudent(student *models.Student) error {
	query := "UPDATE students SET name = ?, age = ?, grade = ? WHERE id = ?"
	_, err := config.DB.Exec(query, student.Name, student.Age, student.Grade, student.ID)
	return err
}

func DeleteStudent(id int) error {
	query := "DELETE FROM students WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	return err
}
