package tests

import (
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudent(t *testing.T) {
	//Create a new instance of student service

	var existingStudent models.Student
	existingStudentEmail := "student1@example.com"
	existingStudent, err = studentService.GetStudent(testDB, existingStudentEmail)
	if err != nil {
		t.Errorf("Error getting student: %v", err)
	}

	assert.Equal(t, existingStudent.Email, existingStudentEmail, "Emails should be the same")

	nonExistingStudentEmail := "studentNOTFOUND@example.com"
	_, err = studentService.GetStudent(testDB, nonExistingStudentEmail)
	assert.Error(t, err, "Error should be returned when student is not found")
}

func TestGetStudents(t *testing.T) {
	existingStudentEmails := []string{
		"student1@example.com",
		"student2@example.com",
		"student3@example.com",
		"student4@example.com",
		"student5@example.com",
		"student6@example.com",
		"student7@example.com",
	}
	var existingStudents []models.Student
	existingStudents, err = studentService.GetStudents(testDB, existingStudentEmails)
	if err != nil {
		t.Errorf("Error getting students: %v", err)
	}

	assert.Equal(t, len(existingStudents), len(existingStudentEmails), "Number of students should be the same")

	nonExistingStudentEmails := []string{
		"student1@example.com",
		"student2@example.com",
		"student3@example.com",
		"student4@example.com",
		"studentNOTFOUND@example.com",
		"student6NOTFOUND@example.com",
		"student7@example.com",
	}

	_, err = studentService.GetStudents(testDB, nonExistingStudentEmails)
	assert.Error(t, err, "Error should be returned when student is not found")
}

func TestCheckSuspension(t *testing.T) {
	suspendedStudentEmail := "student4@example.com"
	nonSuspendedStudentEmail := "student1@example.com"

	assert.True(t, studentService.CheckSuspension(testDB, suspendedStudentEmail), "Student should be suspended")
	assert.False(t, studentService.CheckSuspension(testDB, nonSuspendedStudentEmail), "Student should not be suspended")
}

func TestSuspendStudent(t *testing.T) {
	studentEmail := "student1@example.com"
	err = studentService.SuspendStudent(testDB, studentEmail)
	if err != nil {
		t.Errorf("Error suspending student: %v", err)
	}

	assert.True(t, studentService.CheckSuspension(testDB, studentEmail), "Student should be suspended")

	// Unsuspend the student
	err = testDB.Model(&models.Student{}).Where("email = ?", studentEmail).Update("status", "active").Error
	if err != nil {
		t.Errorf("Error updating database: %v", err)
	}

	assert.False(t, studentService.CheckSuspension(testDB, studentEmail), "Student should not be suspended")
	nonExistingStudentEmail := "studentNOTFOUND@example.com"
	err = studentService.SuspendStudent(testDB, nonExistingStudentEmail)
	assert.Error(t, err, "Error should be returned when student is not found")
}
