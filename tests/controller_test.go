package tests

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lohyangxian/OneCV-Govtech/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//Reminder: NO TESTING OF DB IN HANDLER LAYER
//Only test validation of request and response

func TestRegister_Success(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/register", func(c *gin.Context) {
		mockServer.Register(c)
	})

	// Create a new HTTP request
	requestBody := `{
		"teacher": "teacher1@example.com",
		"students": ["student1@example.com", "student2@example.com"]
	}`

	mockServer.TeacherService.(*mocks.MockTeacherService).On("RegisterStudentsToTeacher", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create a reader from the JSON string
	requestBodyReader := strings.NewReader(requestBody)

	req, err := http.NewRequest("POST", "/api/register", requestBodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code, "Expected status code %d, but got %d", http.StatusNoContent, w.Code)
}

func TestRegister_NotFound(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/register", func(c *gin.Context) {
		mockServer.Register(c)
	})

	// Create a new HTTP request
	requestBody := `{
		"teacher": "teacherNotFound@example.com",
		"students": ["student1@example.com", "student2@example.com"]
	}`

	mockServer.TeacherService.(*mocks.MockTeacherService).On("RegisterStudentsToTeacher", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("teacher / student not found"))

	// Create a reader from the JSON string
	requestBodyReader := strings.NewReader(requestBody)

	req, err := http.NewRequest("POST", "/api/register", requestBodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d, but got %d", http.StatusNotFound, w.Code)
}

func TestRegister_IncompleteRequestBody(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/register", func(c *gin.Context) {
		mockServer.Register(c)
	})

	// Create a new HTTP request
	requestBodyMissingTeacher := `{
    	"students": ["student1@example.com", "student2@example.com"]
	}`

	// Create a reader from the JSON string
	requestBodyMissingTeacherReader := strings.NewReader(requestBodyMissingTeacher)

	req, err := http.NewRequest("POST", "/api/register", requestBodyMissingTeacherReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, but got %d", http.StatusBadRequest, w.Code)

	// Create a new HTTP request
	requestBodyMissingStudent := `{
    	"teacher": "teacher1@example.com"
	}`

	// Create a reader from the JSON string
	requestBodyMissingStudentReader := strings.NewReader(requestBodyMissingStudent)

	req, err = http.NewRequest("POST", "/api/register", requestBodyMissingStudentReader)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
}

func TestCommonStudents_Success(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.GET("/api/commonstudents", func(c *gin.Context) {
		mockServer.CommonStudents(c)
	})

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/commonstudents?teacher=teacher1@example.com", nil)

	mockServer.TeacherService.(*mocks.MockTeacherService).On("GetCommonStudentEmails", mock.Anything, mock.Anything).Return([]string{"student1@example.com"}, nil)

	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, but got %d", http.StatusOK, w.Code)
}

func TestCommonStudents_NotFound(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.GET("/api/commonstudents", func(c *gin.Context) {
		mockServer.CommonStudents(c)
	})

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/commonstudents?teacher=teacherNOTFOUND@example.com", nil)

	mockServer.TeacherService.(*mocks.MockTeacherService).On("GetCommonStudentEmails", mock.Anything, mock.Anything).Return([]string{""}, errors.New("teacher not found"))

	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d, but got %d", http.StatusNotFound, w.Code)
}

func TestCommonStudents_IncompleteRequestBody(t *testing.T) {

	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.GET("/api/commonstudents", func(c *gin.Context) {
		mockServer.CommonStudents(c)
	})

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/commonstudents?", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, but got %d", http.StatusOK, w.Code)
}

func TestSuspend_Success(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/suspend", func(c *gin.Context) {
		mockServer.Suspend(c)
	})

	// Create a new HTTP request
	requestBody := `{
		"student": "student1@example.com"
	}`

	mockServer.StudentService.(*mocks.MockStudentService).On("SuspendStudent", mock.Anything, mock.Anything).Return(nil)

	// Create a reader from the JSON string
	requestBodyReader := strings.NewReader(requestBody)

	req, err := http.NewRequest("POST", "/api/suspend", requestBodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code, "Expected status code %d, but got %d", http.StatusNoContent, w.Code)
}

func TestSuspend_NotFound(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/suspend", func(c *gin.Context) {
		mockServer.Suspend(c)
	})

	// Create a new HTTP request
	requestBody := `{
		"student": "studentNOTFOUND@example.com"
	}`

	mockServer.StudentService.(*mocks.MockStudentService).On("SuspendStudent", mock.Anything, mock.Anything).Return(errors.New("student not found"))

	// Create a reader from the JSON string
	requestBodyReader := strings.NewReader(requestBody)

	req, err := http.NewRequest("POST", "/api/suspend", requestBodyReader)

	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d, but got %d", http.StatusNotFound, w.Code)
}

func TestSuspend_IncompleteRequestBody(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/suspend", func(c *gin.Context) {
		mockServer.Suspend(c)
	})

	// Create a new HTTP request
	requestBodyMissingStudent := `{
	}`

	// Create a reader from the JSON string
	requestBodyMissingStudentReader := strings.NewReader(requestBodyMissingStudent)

	req, err := http.NewRequest("POST", "/api/suspend", requestBodyMissingStudentReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
}

func TestRetrieveForNotifications_Success(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/retrievefornotifications", func(c *gin.Context) {
		mockServer.RetrieveForNotifications(c)
	})

	// Create a new HTTP request
	requestBody := `{
		"teacher": "teacher1@example.com",
		"notification": "Hello students! @student1@example.com @student2@example.com"
	}`

	mockServer.TeacherService.(*mocks.MockTeacherService).On("RetrieveForNotifications", mock.Anything, mock.Anything, mock.Anything).
		Return([]string{"student1@example.com", "student2@example.com", "student3@example.com"}, nil)

	// Create a reader from the JSON string
	requestBodyReader := strings.NewReader(requestBody)

	req, err := http.NewRequest("POST", "/api/retrievefornotifications", requestBodyReader)

	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, but got %d", http.StatusOK, w.Code)
}

func TestRetrieveForNotifications_NotFound(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/retrievefornotifications", func(c *gin.Context) {
		mockServer.RetrieveForNotifications(c)
	})

	// Create a new HTTP request
	requestBody := `{
		"teacher": "teacherNOTFOUND@example.com",
		"notification": "Hello students! @student1@example.com @student2@example.com"
	}`

	mockServer.TeacherService.(*mocks.MockTeacherService).On("RetrieveForNotifications", mock.Anything, mock.Anything, mock.Anything).
		Return([]string{""}, errors.New("teacher not found"))

	// Create a reader from the JSON string
	requestBodyReader := strings.NewReader(requestBody)

	req, err := http.NewRequest("POST", "/api/retrievefornotifications", requestBodyReader)

	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d, but got %d", http.StatusNotFound, w.Code)
}

func TestRetrieveForNotifications_IncompleteRequestBody(t *testing.T) {
	router, mockServer := SetUpMockServer(nil)

	// Make HTTP request to your API endpoint
	router.POST("/api/retrievefornotifications", func(c *gin.Context) {
		mockServer.RetrieveForNotifications(c)
	})

	// Create a new HTTP request
	requestBodyMissingTeacher := `{
		"notification": " "
	}`
	// Create a reader from the JSON string
	requestBodyMissingTeacherReader := strings.NewReader(requestBodyMissingTeacher)

	req, err := http.NewRequest("POST", "/api/retrievefornotifications", requestBodyMissingTeacherReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, but got %d", http.StatusBadRequest, w.Code)

	// Create a new HTTP request
	requestBodyMissingNotification := `{
		"teacher": "teacher1@example.com"
	}`
	// Create a reader from the JSON string
	requestBodyMissingNotificationReader := strings.NewReader(requestBodyMissingNotification)

	req, err = http.NewRequest("POST", "/api/retrievefornotifications", requestBodyMissingNotificationReader)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
}
