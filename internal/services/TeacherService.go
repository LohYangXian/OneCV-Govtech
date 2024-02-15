package services

import "github.com/lohyangxian/OneCV-Govtech/internal/models"

func CreateTeacher() error {
	return nil
}

func GetTeacher() error {
	return nil
}

func GetTeachers() error {
	return nil
}

func CheckTeacherExist() error {
	return nil
}

func RegisterStudentsToTeachers() error {
	//	Check if the request is a POST request (DO IN CONTROLLER LAYER)
	//  From the request body, extract the teacher's email and the students' emails
	//  Check if the teacher's email is already registered
	// Return an error if teacher's email is not found
	//  Check if the students' emails are already registered
	// Return an error if students' email are not found
	// Update teacher's students list with the students' emails
	// Update students' teachers list with the teacher's email
	return nil
}

func GetCommonStudents() ([]models.Student, error) {
	//	Check if the request is a GET request (DO IN CONTROLLER LAYER)
	//  From the request body, extract the teachers' emails
	//  Check if the teachers' emails are already registered
	// Return an error if teachers' email are not found
	//  Get the students' emails from the teachers' emails
	// Remove duplicates
	// Return the students' emails
	var students []models.Student

	return students, nil
}

func RetrieveForNotifications() ([]models.Student, error) {
	//	Check if the request is a POST request (DO IN CONTROLLER LAYER)
	//  From the request body, extract the teacher's email and the notification message
	//  Check if the teacher's email is already registered
	// Return an error if teacher's email is not found
	//  Get the students' emails from the teacher's email
	//  Get the students' emails from the notification message
	//  Check if the students' emails are already registered
	// Return an error if students' email are not found
	//  Check if students' emails are suspended
	// Return an error if students' email are suspended
	// Check if teacher's email is in the students' teachers list
	// Return an error if teacher's email is not found
	// Remove duplicates
	// Return the students' emails
	var students []models.Student

	return students, nil
}
