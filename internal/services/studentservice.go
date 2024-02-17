package services

import (
	"errors"
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"gorm.io/gorm"
)

type StudentService interface {
	GetStudent(db *gorm.DB, studentEmail string) (models.Student, error)
	GetStudents(db *gorm.DB, studentEmails []string) ([]models.Student, error)
	CheckSuspension(db *gorm.DB, studentEmail string) bool
	SuspendStudent(db *gorm.DB, studentEmail string) error
}

type StudentServiceImpl struct {
	DB *gorm.DB
}

func NewStudentService(db *gorm.DB) StudentService {
	return &StudentServiceImpl{DB: db}
}

func (s *StudentServiceImpl) GetStudent(db *gorm.DB, studentEmail string) (models.Student, error) {
	var student models.Student
	err := db.Model(models.Student{}).Where("email = ?", studentEmail).First(&student).Error

	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *StudentServiceImpl) GetStudents(db *gorm.DB, studentEmails []string) ([]models.Student, error) {
	var students []models.Student
	err := db.Model(models.Student{}).Where("email IN ?", studentEmails).Find(&students).Error
	if len(students) != len(studentEmails) {
		return students, errors.New("not all students found")
	}
	return students, err
}

func (s *StudentServiceImpl) CheckSuspension(db *gorm.DB, studentEmail string) bool {
	var status string
	db.Model(&models.Student{}).Select("status").Where("email = ?", studentEmail).First(&status)

	return status == "suspended"
}

func (s *StudentServiceImpl) SuspendStudent(db *gorm.DB, studentEmail string) error {
	student, err := s.GetStudent(db, studentEmail)

	// Suspend the student's email
	err = db.Model(&student).Update("status", "suspended").Error

	if err != nil {
		return err
	}

	return nil
}
