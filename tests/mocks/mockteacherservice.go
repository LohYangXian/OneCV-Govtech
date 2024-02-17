package mocks

import (
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockTeacherService is a mock implementation of the TeacherService interface.
type MockTeacherService struct {
	mock.Mock
}

// GetTeacher mocks the GetTeacher method.
func (m *MockTeacherService) GetTeacher(db *gorm.DB, teacherEmail string) (models.Teacher, error) {
	args := m.Called(db, teacherEmail)
	return args.Get(0).(models.Teacher), args.Error(1)
}

// RegisterStudentsToTeacher mocks the RegisterStudentsToTeacher method.
func (m *MockTeacherService) RegisterStudentsToTeacher(db *gorm.DB, studentEmails []string, teacherEmail string) error {
	args := m.Called(db, studentEmails, teacherEmail)
	return args.Error(0)
}

// GetCommonStudentEmails mocks the GetCommonStudentEmails method.
func (m *MockTeacherService) GetCommonStudentEmails(db *gorm.DB, teacherEmails []string) ([]string, error) {
	args := m.Called(db, teacherEmails)
	return args.Get(0).([]string), args.Error(1)
}

// RetrieveForNotifications mocks the RetrieveForNotifications method.
func (m *MockTeacherService) RetrieveForNotifications(db *gorm.DB, teacherEmail string, notification string) ([]string, error) {
	args := m.Called(db, teacherEmail, notification)
	return args.Get(0).([]string), args.Error(1)
}
