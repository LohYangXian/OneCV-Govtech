CREATE TABLE IF NOT EXISTS teachers
(
    id    serial PRIMARY KEY,
    email VARCHAR(300) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS students
(
    id     serial PRIMARY KEY,
    email  VARCHAR(300) UNIQUE NOT NULL,
    status VARCHAR(50)         NOT NULL
);

-- Define a many-to-many relationship table
CREATE TABLE IF NOT EXISTS teacher_student
(
    teacher_id INT REFERENCES teachers (id),
    student_id INT REFERENCES students (id),
    CONSTRAINT teacher_student_pk PRIMARY KEY (teacher_id, student_id)
);