package tests

import (
	"github.com/lohyangxian/OneCV-Govtech/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudent(t *testing.T) {
	//Create a new instance of student service
	studentService := services.NewStudentService(testDB)

	existingStudentEmail := "student1@example.com"
	existingStudent, err := studentService.GetStudent(testDB, existingStudentEmail)
	if err != nil {
		t.Errorf("Error getting student: %v", err)
	}

	assert.Equal(t, existingStudent.Email, existingStudentEmail, "Emails should be the same")

	nonExistingStudentEmail := "studentNOTFOUND@example.com"
	_, err = studentService.GetStudent(testDB, nonExistingStudentEmail)
	assert.Error(t, err, "Error should be returned when student is not found")
}
