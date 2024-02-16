package services

import (
	"errors"
	"github.com/lohyangxian/OneCV-Govtech/internal/models"
	"gorm.io/gorm"
	"log"
	"regexp"
)

// TODO: Extract SQL Queries to another class if possible
func CreateTeacher() error {
	return nil
}

func GetTeacher(db *gorm.DB, teacherEmail string) (models.Teacher, error) {
	teacher := models.Teacher{}

	err := db.Model(teacher).Where("email = ?", teacherEmail).First(&teacher).Error

	//TODO: Create Custom Error Message
	//TODO: Improve Error Handling
	return teacher, err
}

func CheckTeacherExist() bool {
	return false
}

func removeDuplicates(users []string) []string {
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

func addTeachersAndStudents(db *gorm.DB, teacher models.Teacher, students []models.Student) error {
	/*
		Quite a costly operation. GORM might need to update an intermediate left join table to reflect the changes
	*/
	err := db.Model(&teacher).Association("Students").Append(students)

	if err != nil {
		return err
	}

	log.Println(teacher.Students)

	return nil
}

func GetStudentsFromNotification(notification string) ([]string, error) {
	var studentEmails []string

	// Define the regular expression pattern to match email addresses
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

	return studentEmails, nil
}

func RegisterStudentsToTeacher(db *gorm.DB, studentEmails []string, teacherEmail string) error {
	//	Check if the request is a POST request (DO IN CONTROLLER LAYER)
	//  From the request body, extract the teacher's email and the students' emails
	//  Check if the teacher's email is already registered
	// Return an error if teacher's email is not found
	//  Check if the students' emails are already registered
	// Return an error if students' email are not found
	// Update teacher's students list with the students' emails
	// Update students' teachers list with the teacher's email

	studentEmails = removeDuplicates(studentEmails)
	teacher, err := GetTeacher(db, teacherEmail)

	//TODO: Create Custom Error Message
	if err != nil {
		return errors.New("teacher not found")
	}

	students, err := GetStudents(db, studentEmails)

	//TODO: Create Custom Error Message
	if err != nil {
		return errors.New("students not found")
	}

	//TODO: Create Custom Error Message
	if len(students) != len(studentEmails) {
		return errors.New("some students not found")
	}

	//add students to teacher's students list
	//add teacher to students' teachers list
	//TODO: Create SQL Query
	err = addTeachersAndStudents(db, teacher, students)

	if err != nil {
		return err
	}

	return nil
}

func GetCommonStudentEmails(db *gorm.DB, teacherEmails []string) ([]string, error) {
	//	Check if the request is a GET request (DO IN CONTROLLER LAYER)
	//  From the request body, extract the teachers' emails
	//  Check if the teachers' emails are already registered
	// Return an error if teachers' email are not found
	//  Get the students' emails from the teachers' emails
	// Remove duplicates
	// Return the students' emails
	var studentEmails []string

	//TODO: Add Error handling if teacher is not found

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

	var studentEmails []string

	_, err := GetTeacher(db, teacherEmail)

	if err != nil {
		return studentEmails, errors.New("teacher not found")
	}

	//DB query

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
	studentEmailsFromNotification, err := GetStudentsFromNotification(notification)

	for _, email := range studentEmailsFromNotification {
		_, err = CheckStudentExist(db, email)
		if err != nil {
			return studentEmails, errors.New("student not found")
		}
		if CheckSuspension(db, email) == false {
			studentEmails = append(studentEmails, email)
		}
	}

	if err != nil {
		return studentEmails, err
	}

	//Step 3: Remove duplicates
	studentEmails = removeDuplicates(studentEmails)

	return studentEmails, nil
}
