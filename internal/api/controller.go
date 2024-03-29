package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lohyangxian/OneCV-Govtech/internal/errors"
)

// Register registers one or more students to a specified teacher.
//
// Method: POST
// Endpoint: /api/register
// Headers: Content-Type: application/json
// Success response status: HTTP 204
// Request body example:
//
//	{
//	  "teacher": "teacherken@gmail.com",
//	  "students": [
//	    "studentjon@gmail.com",
//	    "studenthon@gmail.com"
//	  ]
//	}
//
// @Summary Register students to a teacher
// @Description Register one or more students to a specified teacher.
// @Accept json
// @Produce json
// @Param teacher body string true "Email address of the teacher"
// @Param students body []string true "List of student email addresses"
// @Success 204 {string} string "Success"
// @Router /api/register [post]
func (s *Server) Register(c *gin.Context) {
	var requestBody struct {
		TeacherEmail  string   `json:"teacher" binding:"required"`
		StudentEmails []string `json:"students" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		responseError := errors.NewBadRequestError("missing required fields: teacher / students")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	err = s.TeacherService.RegisterStudentsToTeacher(s.Database, requestBody.StudentEmails, requestBody.TeacherEmail)

	if err != nil {
		responseError := errors.NewNotFoundError("student or teacher not found")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	c.JSONP(204, "Success")
}

// CommonStudents retrieves a list of students common to a given list of teachers.
//
// Method: GET
// Endpoint: /api/commonstudents
// Success response status: HTTP 200
// Request example 1: GET /api/commonstudents?teacher=teacherken%40gmail.com
// Success response body 1:
//
//	{
//	  "students" : [
//	    "commonstudent1@gmail.com",
//	    "commonstudent2@gmail.com",
//	    "student_only_under_teacher_ken@gmail.com"
//	  ]
//	}
//
// Request example 2: GET /api/commonstudents?teacher=teacherken%40gmail.com&teacher=teacherjoe%40gmail.com
// Success response body 2:
//
//	{
//	  "students" : [
//	    "commonstudent1@gmail.com",
//	    "commonstudent2@gmail.com"
//	  ]
//	}
//
// @Summary Retrieve common students
// @Description Retrieve a list of students common to a given list of teachers.
// @Produce json
// @Param teacher query string true "Teacher email address"
// @Success 200 {object} string "Success"
// @Router /api/commonstudents [get]
func (s *Server) CommonStudents(c *gin.Context) {
	var teacherEmails = c.QueryArray("teacher")

	if len(teacherEmails) == 0 {
		responseError := errors.NewBadRequestError("missing required field: teacher")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	students, err := s.TeacherService.GetCommonStudentEmails(s.Database, teacherEmails)

	if err != nil {
		responseError := errors.NewNotFoundError("teacher not found")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	c.JSONP(200, gin.H{"students": students})
}

// Suspend suspends a specified student.
//
// Method: POST
// Endpoint: /api/suspend
// Headers: Content-Type: application/json
// Success response status: HTTP 204
// Request body example:
//
//	{
//	  "student" : "studentmary@gmail.com"
//	}
//
// @Summary Suspend a student
// @Description Suspend a specified student.
// @Accept json
// @Produce json
// @Param student body string true "Email address of the student to suspend"
// @Success 204 {string} string "Success"
// @Router /api/suspend [post]
func (s *Server) Suspend(c *gin.Context) {
	var requestBody struct {
		StudentEmail string `json:"student" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		responseError := errors.NewBadRequestError("missing required field: student")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	err = s.StudentService.SuspendStudent(s.Database, requestBody.StudentEmail)

	if err != nil {
		responseError := errors.NewNotFoundError("student not found")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	c.JSONP(204, "Success")
}

// RetrieveForNotifications retrieves a list of students who can receive a given notification.
//
// Method: POST
// Endpoint: /api/retrievefornotifications
// Headers: Content-Type: application/json
// Success response status: HTTP 200
// Request body example 1:
//
//	{
//	  "teacher":  "teacherken@gmail.com",
//	  "notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
//	}
//
// Success response body 1:
//
//	{
//	  "recipients": [
//	    "studentbob@gmail.com",
//	    "studentagnes@gmail.com",
//	    "studentmiche@gmail.com"
//	  ]
//	}
//
// Request body example 2:
//
//	{
//	  "teacher":  "teacherken@gmail.com",
//	  "notification": "Hey everybody"
//	}
//
// Success response body 2:
//
//	{
//	  "recipients": [
//	    "studentbob@gmail.com"
//	  ]
//	}
//
// @Summary Retrieve students for notifications
// @Description Retrieve a list of students who can receive a given notification.
// @Accept json
// @Produce json
// @Param teacher body string true "Email address of the teacher"
// @Param notification body string true "Notification message"
// @Success 200 {object} string "Success"
// @Router /api/retrievefornotifications [post]
func (s *Server) RetrieveForNotifications(c *gin.Context) {
	var requestBody struct {
		TeacherEmail string `json:"teacher" binding:"required"`
		Notification string `json:"notification" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		responseError := errors.NewBadRequestError("missing required fields: teacher / notification")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	recipients, err := s.TeacherService.RetrieveForNotifications(s.Database, requestBody.TeacherEmail, requestBody.Notification)

	if err != nil {
		responseError := errors.NewNotFoundError("student or teacher not found")
		c.AbortWithStatusJSON(responseError.Status(), gin.H{"message": responseError.Error()})
		return
	}

	c.JSONP(200, gin.H{"recipients": recipients})
}
