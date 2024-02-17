package tests

//func TestRegister_NotFound(t *testing.T) {
//
//	// Set up a mock db
//	testDB, err := SetUpTestDBConnection()
//	if err != nil {
//		t.Errorf("failed to connect to the test database: %v", err)
//	}
//
//	// Seed the test database
//	SetUpTestDB(testDB)
//
//	router, mockServer := SetUpMockServer(testDB)
//
//	// Make HTTP request to your API endpoint
//	router.POST("/api/register", func(c *gin.Context) {
//		mockServer.Register(c)
//	})
//
//	// Test case 1: Teacher not found
//	requestBodyTeacherNotFound := `{
//		"teacher": "teacherNotFound@example.com",
//		"students": ["student1@example.com", "student2@example.com"]
//	}`
//
//	requestBodyTeacherNotFoundReader := strings.NewReader(requestBodyTeacherNotFound)
//
