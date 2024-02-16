package services

import (
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"gorm.io/gorm"
)

// TODO: Check if this function is required
func CreateStudent() error {
	return nil
}

func GetStudent(db *gorm.DB, studentEmail string) (models.Student, error) {
	var student models.Student
	err := db.Model(models.Student{}).Where("email = ?", studentEmail).First(&student).Error

	if err != nil {
		return student, err
	}

	return student, nil
}

func GetStudents(db *gorm.DB, studentEmails []string) ([]models.Student, error) {
	var students []models.Student
	err := db.Model(models.Student{}).Where("email IN ?", studentEmails).Find(&students).Error
	return students, err
}

func CheckSuspension(db *gorm.DB, studentEmail string) bool {
	var status string
	db.Model(&models.Student{}).Select("status").Where("email = ?", studentEmail).First(&status)

	return status == "suspended"
}

func SuspendStudent(db *gorm.DB, studentEmail string) error {
	student, err := GetStudent(db, studentEmail)

	// Suspend the student's email
	err = db.Model(&student).Update("status", "suspended").Error

	if err != nil {
		return err
	}

	return nil
}
