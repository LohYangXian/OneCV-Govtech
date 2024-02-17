# OneCV-Govtech API

This API allows teachers to perform administrative functions for their students. Teachers and students are identified by their email addresses. It provides endpoints for teachers to manage students, retrieve common students, suspend students, and send notifications to students.

## Hosted API

- Currently, the API is not hosted. You need to run it locally to test the functionality.

## Running the API Locally

To run the API locally, follow these instructions:

1. **Clone the Repository**: Clone the OneCV-Govtech repository to your local machine.

    ```bash
    git clone <repository_url>
    ```

2. **Navigate to the Project Directory**: Move into the cloned project directory.

    ```bash
    cd OneCV-Govtech
    ```

3. **Set Up Configuration**:
    - Update the database configuration in `config.yml` according to your local setup.


4. **Build and Run the API**:
    - Ensure you have Go installed on your system.
    - Run the following command to build and start the API server:

    ```bash
   make build
    make run
    ```

5. **Access the API**:
    - The API will be available at `http://localhost:3000`.


6. **Run Tests**:
    - Run the following command to execute the tests:

    ```bash
    make test
    ```


## Project Folder Structure

```
.
├── bin
|   └── OneCV-Govtech.exe
├── config
│   └── config.go
├── internal
│   ├── api
│   │    ├── controller.go
│   │    └── server.go
│   ├── database
│   │    ├── db.go
│   │    └── migrations
│   │        └── ...
│   ├── errors
│   │    ├── badrequesterror.go
│   │    └── notfounderror.go
│   ├── models
│   │   ├── student.go
│   │   └── teacher.go
│   └── services
│       ├── studentservice.go
│       └── teacherservice.go
├── tests
│   ├── main_test.go
│   ├── studentservice_test.go
│   ├── teacherservice_test.go
│   ├── controller_test.go
│   ├── mocks
│   │   ├── mockstudentservice.go
│   │   └── mockteacherservice.go
│   └── test_db
│       └── ...
├── gitignore
├── config.yml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

This Markdown representation provides a visual representation of your project's folder structure. You can paste it into your README.md file to help users understand the organization of your project's codebase.

## API Endpoints

### 1. Register Students to a Teacher

- **Endpoint**: `POST /api/register`
- **Request Body**:

    ```json
    {
        "teacher": "teacherken@gmail.com",
        "students": ["studentjon@gmail.com", "studenthon@gmail.com"]
    }
    ```

- **Success Response**:
    - HTTP Status Code: `204 No Content`
- **Error Responses**:
    - HTTP Status Code: `400 Bad Request`
      - When the request payload is invalid or missing required fields.
    - HTTP Status Code: `404 Not Found`
      - When the teacher or student does not exist in the database.

### 2. Retrieve Common Students

- **Endpoint**: `GET /api/commonstudents`
- **Request Example**:

    ```
    GET /api/commonstudents?teacher=teacherken%40gmail.com&teacher=teacherjoe%40gmail.com
    ```

- **Success Response**:
    - HTTP Status Code: `200 OK`
    - Response Body:

    ```json
    {
        "students": ["commonstudent1@gmail.com", "commonstudent2@gmail.com"]
    }
    ```
- **Error Responses**:
- HTTP Status Code: `400 Bad Request`
  - When the request payload is invalid or missing required fields.
- HTTP Status Code: `404 Not Found`
  - When the teacher does not exist in the database.

### 3. Suspend a Student

- **Endpoint**: `POST /api/suspend`
- **Request Body**:

    ```json
    {
        "student": "studentmary@gmail.com"
    }
    ```

- **Success Response**:
    - HTTP Status Code: `204 No Content`
- **Error Responses**:
    - HTTP Status Code: `400 Bad Request`
      - When the request payload is invalid or missing required fields.
    - HTTP Status Code: `404 Not Found`
      - When the student does not exist in the database.

### 4. Retrieve Students for Notifications

- **Endpoint**: `POST /api/retrievefornotifications`
- **Request Body**:

    ```json
    {
        "teacher":  "teacherken@gmail.com",
        "notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
    }
    ```

- **Success Response**:
    - HTTP Status Code: `200 OK`
    - Response Body:

    ```json
    {
        "recipients": ["studentbob@gmail.com", "studentagnes@gmail.com", "studentmiche@gmail.com"]
    }
    ```
- **Error Responses**:
- HTTP Status Code: `400 Bad Request`
  - When the request payload is invalid or missing required fields.
- HTTP Status Code: `404 Not Found`
  - When the teacher does not exist in the database.

## Future Improvements

- ***Improve Validation of Request Payloads***: The API should have better validation for e.g. valid emails to ensure that the data is in the correct format and contains all the required fields.
- ***Improve Error Handling***: The error handling can be improved to provide more detailed error messages to the client.
- ***Improve Efficiency of Retrieving Common Students***: The current implementation of retrieving common students is not efficient. It can be improved to reduce the number of database queries.
- ***Potential Refactoring of Code***: The data structure of the models can be refactored to allow more efficient SQL queries
- ***Hosting the API***: The API should be hosted on a cloud platform to make it accessible to the public.