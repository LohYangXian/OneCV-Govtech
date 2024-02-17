package mocks

import (
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockStudentService is a mock implementation of the StudentService interface
type MockStudentService struct {
	mock.Mock
}

// GetStudent mocks the GetStudent method of StudentService interface
func (m *MockStudentService) GetStudent(db *gorm.DB, studentEmail string) (models.Student, error) {
	args := m.Called(db, studentEmail)
	return args.Get(0).(models.Student), args.Error(1)
}

// GetStudents mocks the GetStudents method of StudentService interface
func (m *MockStudentService) GetStudents(db *gorm.DB, studentEmails []string) ([]models.Student, error) {
	args := m.Called(db, studentEmails)
	return args.Get(0).([]models.Student), args.Error(1)
}

// CheckSuspension mocks the CheckSuspension method of StudentService interface
func (m *MockStudentService) CheckSuspension(db *gorm.DB, studentEmail string) bool {
	args := m.Called(db, studentEmail)
	return args.Bool(0)
}

// SuspendStudent mocks the SuspendStudent method of StudentService interface
func (m *MockStudentService) SuspendStudent(db *gorm.DB, studentEmail string) error {
	args := m.Called(db, studentEmail)
	return args.Error(0)
}
