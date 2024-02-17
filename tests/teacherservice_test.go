package tests

import (
	"errors"
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"github.com/lohyangxian/OneCV-Govtech/internal/services"
	"github.com/lohyangxian/OneCV-Govtech/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTeacher(t *testing.T) {
	existingTeacherEmail := "teacher1@example.com"

	var existingTeacher models.Teacher
	existingTeacher, err = teacherService.GetTeacher(testDB, existingTeacherEmail)
	if err != nil {
		t.Errorf("Error getting student: %v", err)
	}

	assert.Equal(t, existingTeacher.Email, existingTeacherEmail, "Emails should be the same")

	nonExistingTeacherEmail := "teacherNOTFOUND@example.com"
	_, err = teacherService.GetTeacher(testDB, nonExistingTeacherEmail)
	assert.Error(t, err, "Error should be returned when teacher is not found")
}

func TestRegisterStudentsToTeacher(t *testing.T) {
	teacherEmail := "teacher1@example.com"
	studentEmails := []string{
		"student5@example.com",
		"student6@example.com",
		"student7@example.com",
		"student8@example.com",
		"student9@example.com",
	}

	//	Mock the student service
	mockStudentService := &mocks.MockStudentService{}
	teacherServiceWithMockedStudentService := &services.TeacherServiceImpl{
		StudentService: mockStudentService,
	}

	mockStudentService.On("GetStudents", testDB, studentEmails).Return([]models.Student{
		{ID: 5, Email: studentEmails[0], Status: "active"},
		{ID: 6, Email: studentEmails[1], Status: "active"},
		{ID: 7, Email: studentEmails[2], Status: "suspended"},
		{ID: 8, Email: studentEmails[3], Status: "active"},
		{ID: 9, Email: studentEmails[4], Status: "suspended"},
	}, nil)

	err = teacherServiceWithMockedStudentService.RegisterStudentsToTeacher(testDB, studentEmails, teacherEmail)
	assert.NoError(t, err)
	mockStudentService.AssertExpectations(t)

	var registeredStudentsEmails []string
	//	Verify that the students are registered to the teacher
	registeredStudentsEmails, err = teacherServiceWithMockedStudentService.GetCommonStudentEmails(testDB, []string{teacherEmail})
	if err != nil {
		t.Errorf("Error getting students: %v", err)
	}
	assert.NoError(t, err)

	// Verify that each student email is in the registeredStudentsEmails
	for _, studentEmail := range studentEmails {
		assert.Contains(t, registeredStudentsEmails, studentEmail, "Number of students should be the same")
	}

	// Test for error when student is not found
	studentEmailsNotFound := []string{
		"student5@example.com",
		"student6@example.com",
		"student7@example.com",
		"student8@example.com",
		"studentNOTFOUND@example.com",
	}

	mockStudentService.On("GetStudents", testDB, studentEmailsNotFound).Return([]models.Student{
		{ID: 5, Email: studentEmails[0], Status: "active"},
		{ID: 6, Email: studentEmails[1], Status: "active"},
		{ID: 7, Email: studentEmails[2], Status: "suspended"},
		{ID: 8, Email: studentEmails[3], Status: "active"},
	}, errors.New("not all students found"))

	err = teacherServiceWithMockedStudentService.RegisterStudentsToTeacher(testDB, studentEmailsNotFound, teacherEmail)

	mockStudentService.AssertExpectations(t)

	assert.Error(t, err, "Error should be returned when student is not found")

	// Test for error when teacher is not found
	teacherEmailNotFound := "teacherNotFound"
	err = teacherServiceWithMockedStudentService.RegisterStudentsToTeacher(testDB, studentEmailsNotFound, teacherEmailNotFound)
	assert.Error(t, err, "Error should be returned when teacher is not found")

	//Clean up the database
	TearDownTestDB(testDB)
	SetUpTestDB(testDB)
}

func TestGetCommonStudentEmails(t *testing.T) {
	teacherEmailsNotFound := []string{"teacherNOTFOUND@example.com"}
	_, err = teacherService.GetCommonStudentEmails(testDB, teacherEmailsNotFound)
	assert.Error(t, err, "Error should be returned when teacher is not found")

	teacherEmails := []string{
		"teacher1@example.com",
		"teacher2@example.com",
	}

	commonStudentEmails, err := teacherService.GetCommonStudentEmails(testDB, teacherEmails)
	expectedCommonStudentEmails := []string{
		"student1@example.com",
		"student2@example.com",
		"student3@example.com",
	}

	assert.NoError(t, err)
	assert.ElementsMatch(t, commonStudentEmails, expectedCommonStudentEmails, "Common student emails should be the same")

	teacherEmails2 := []string{
		"teacher1@example.com",
	}

	commonStudentEmails2, err := teacherService.GetCommonStudentEmails(testDB, teacherEmails2)
	expectedCommonStudentEmails2 := []string{
		"student1@example.com",
		"student2@example.com",
		"student3@example.com",
		"student4@example.com",
	}

	assert.NoError(t, err)
	assert.ElementsMatch(t, commonStudentEmails2, expectedCommonStudentEmails2, "Common student emails should be the same")

	teacherEmailsDuplicated := []string{
		"teacher1@example.com",
		"teacher2@example.com",
		"teacher2@example.com",
	}

	commonStudentEmailsDuplicated, err := teacherService.GetCommonStudentEmails(testDB, teacherEmailsDuplicated)
	expectedCommonStudentEmailsDuplicated := []string{
		"student1@example.com",
		"student2@example.com",
		"student3@example.com",
	}

	assert.NoError(t, err)
	assert.ElementsMatch(t, commonStudentEmailsDuplicated, expectedCommonStudentEmailsDuplicated, "Common student emails should be the same")
}

func TestRetrieveForNotifications(t *testing.T) {
	teacherEmailNotFound := "teacherNOTFOUND@example.com"
	_, err = teacherService.RetrieveForNotifications(testDB, teacherEmailNotFound, "Hello students")
	assert.Error(t, err, "Error should be returned when teacher is not found")

	teacherEmail := "teacher1@example.com"
	notificationStudentNotFound := "Hello students! @student1@example.com @studentNOTFOUND@example.com @student9@example.com"
	_, err = teacherService.RetrieveForNotifications(testDB, teacherEmail, notificationStudentNotFound)
	assert.Error(t, err, "Error should be returned when student is not found")

	notification := "Hello students! @student1@example.com @student2@example.com @student9@example.com"
	expectedStudentEmailsFromNotification := []string{
		"student1@example.com",
		"student2@example.com",
		"student9@example.com",
	}
	assert.ElementsMatch(t, expectedStudentEmailsFromNotification, expectedStudentEmailsFromNotification, "Student emails from notification should be the same")

	var studentEmails []string
	studentEmails, err = teacherService.RetrieveForNotifications(testDB, teacherEmail, notification)
	assert.NoError(t, err)

	expectedStudentEmails := []string{
		"student1@example.com",
		"student2@example.com",
		"student3@example.com",
	}

	assert.ElementsMatch(t, studentEmails, expectedStudentEmails, "Student emails should be the same, with no duplicates or suspended students")
}
