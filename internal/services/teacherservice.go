package services

import (
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"gorm.io/gorm"
	"regexp"
)

// TODO: Check if this function is required
func CreateTeacher() error {
	return nil
}

func GetTeacher(db *gorm.DB, teacherEmail string) (models.Teacher, error) {
	teacher := models.Teacher{}

	err := db.Model(teacher).Where("email = ?", teacherEmail).First(&teacher).Error

	return teacher, err
}

func RemoveDuplicates(users []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, user := range users {
		if !seen[user] {
			result = append(result, user)
			seen[user] = true
		}
	}
	return result
}

func GetStudentsFromNotification(notification string) []string {
	var studentEmails []string

	// Define the regular expression pattern to match email addresses
	// This is assuming that the notification string is well-formed, where the email addresses
	// starts with a '@' symbol
	emailRegex := `@[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}`

	// Compile the regular expression pattern
	re := regexp.MustCompile(emailRegex)

	// Find all email addresses in the notification string
	emails := re.FindAllString(notification, -1)

	for _, email := range emails {
		// Remove the leading '@' symbol
		email = email[1:]
		studentEmails = append(studentEmails, email)
	}

	return studentEmails
}

func RegisterStudentsToTeacher(db *gorm.DB, studentEmails []string, teacherEmail string) error {
	studentEmails = RemoveDuplicates(studentEmails)
	teacher, err := GetTeacher(db, teacherEmail)

	if err != nil {
		return err
	}

	students, err := GetStudents(db, studentEmails)

	if err != nil || len(students) != len(studentEmails) {
		return err
	}

	err = db.Model(&teacher).Association("Students").Append(students)

	if err != nil {
		return err
	}

	return nil
}

func GetCommonStudentEmails(db *gorm.DB, teacherEmails []string) ([]string, error) {
	var studentEmails []string

	for _, email := range teacherEmails {
		_, err := GetTeacher(db, email)
		if err != nil {
			return studentEmails, err
		}
	}

	//This Query consists of 2 Nested Loops, which is not efficient
	//TODO: Create more efficient SQL Query
	err := db.
		Select("DISTINCT students.email").
		Table("students").
		Joins("JOIN teacher_student ON students.id = teacher_student.student_id").
		Joins("JOIN teachers ON teacher_student.teacher_id = teachers.id").
		Where("teachers.email IN ?", teacherEmails).
		Group("students.id").
		Having("COUNT(DISTINCT teachers.id) = ?", len(teacherEmails)).
		Find(&studentEmails).Error

	if err != nil {

		return studentEmails, err
	}

	return studentEmails, nil
}

func RetrieveForNotifications(db *gorm.DB, teacherEmail string, notification string) ([]string, error) {
	var studentEmails []string

	_, err := GetTeacher(db, teacherEmail)

	if err != nil {
		return studentEmails, err
	}

	//Step 1: Get students registered with teacher, and not suspended
	err = db.
		Select("DISTINCT students.email").
		Table("students").
		Joins("JOIN teacher_student ON students.id = teacher_student.student_id").
		Joins("JOIN teachers ON teacher_student.teacher_id = teachers.id").
		Where("teachers.email = ?", teacherEmail).
		Where("students.status = ?", "active").
		Find(&studentEmails).Error

	if err != nil {
		return studentEmails, err
	}

	//Step 2: Get students from notification
	studentEmailsFromNotification := GetStudentsFromNotification(notification)

	if err != nil {
		return studentEmails, err
	}

	for _, email := range studentEmailsFromNotification {
		_, err = GetStudent(db, email)
		if err != nil {
			return studentEmails, err
		}
		if CheckSuspension(db, email) == false {
			studentEmails = append(studentEmails, email)
		}
	}

	//Step 3: Remove duplicates
	studentEmails = RemoveDuplicates(studentEmails)

	return studentEmails, nil
}
