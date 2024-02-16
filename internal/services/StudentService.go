package services

import (
	"errors"
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"gorm.io/gorm"
	"log"
)

func CreateStudent() error {
	return nil
}

// Theres CheckStudentExist already
// TODO: Remove this function
func GetStudent() error {
	return nil
}

// TODO: Combine GetStudents and GetUnsuspendedStudents, with a flag to check if the student is suspended
func GetStudents(db *gorm.DB, studentEmails []string) ([]models.Student, error) {
	var students []models.Student
	err := db.Model(models.Student{}).Where("email IN ?", studentEmails).Find(&students).Error
	return students, err
}

func GetUnsuspendedStudents(db *gorm.DB, studentEmails []string) ([]models.Student, error) {
	var students []models.Student
	err := db.Model(models.Student{}).Where("email IN ?", studentEmails).Where("status = ?", "active").Find(&students).Error
	return students, err
}

func CheckStudentExist(db *gorm.DB, studentEmail string) (models.Student, error) {
	var student models.Student
	err := db.Model(models.Student{}).Where("email = ?", studentEmail).First(&student).Error

	if err != nil {
		return student, err
	}

	return student, nil
}

func SuspendStudent(db *gorm.DB, studentEmail string) error {
	//	Check if the request is a POST request (DO IN CONTROLLER LAYER)
	//  From the request body, extract the student's email
	//  Check if the student's email is already registered
	// Return an error if student's email is not found
	//  Check if the student's email is already suspended
	// Return an error if student's email is already suspended
	//  Suspend the student's email

	student, err := CheckStudentExist(db, studentEmail)

	//TODO: Create Custom Error Message
	if err != nil {
		return errors.New("student does not exist")
	}

	//TODO: Change Status to Constant String
	if CheckSuspension(db, studentEmail) {
		return errors.New("student is already suspended")
	}

	// Suspend the student's email
	if err = db.Model(&student).Update("status", "suspended").Error; err != nil {
		return err
	}

	return nil
}

func CheckSuspension(db *gorm.DB, studentEmail string) bool {
	var status string
	db.Model(&models.Student{}).Select("status").Where("email = ?", studentEmail).First(&status)

	log.Println(status)
	//TODO: Change Status to Constant String
	return status == "suspended"
}

func CheckIsTeacherRegistered() error {
	return nil
}
