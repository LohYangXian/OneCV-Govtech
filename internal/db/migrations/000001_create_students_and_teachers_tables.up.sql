CREATE TABLE IF NOT EXISTS teachers(
   teacher_id serial PRIMARY KEY,
   email VARCHAR (300) UNIQUE NOT NULL,
   students INT[]
);

CREATE TABLE IF NOT EXISTS students(
   student_id serial PRIMARY KEY,
   email VARCHAR (300) UNIQUE NOT NULL,
   status VARCHAR (50) NOT NULL,
   teachers INT[]
);

-- Define a many-to-many relationship table
CREATE TABLE IF NOT EXISTS teacher_students (
   teacher_id INT REFERENCES teachers(teacher_id),
   student_id INT REFERENCES students(student_id),
   CONSTRAINT teacher_student_pk PRIMARY KEY (teacher_id, student_id)
);