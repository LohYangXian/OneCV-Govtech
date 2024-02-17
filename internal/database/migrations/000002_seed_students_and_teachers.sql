CREATE TABLE IF NOT EXISTS students (
                                        id SERIAL PRIMARY KEY,
                                        email VARCHAR(255) UNIQUE,
                                        status VARCHAR(50) -- Change the data type as per your requirements
);

CREATE TABLE IF NOT EXISTS teachers (
                                        id SERIAL PRIMARY KEY,
                                        email VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS teacher_student (
                                               teacher_id INT,
                                               student_id INT,
                                               FOREIGN KEY (teacher_id) REFERENCES teachers(id),
                                               FOREIGN KEY (student_id) REFERENCES students(id),
                                               PRIMARY KEY (teacher_id, student_id)
);

-- Create 10 teachers
INSERT INTO teachers (email) VALUES
                                 ('teacher1@example.com'),
                                 ('teacher2@example.com'),
                                 ('teacher3@example.com'),
                                 ('teacher4@example.com'),
                                 ('teacher5@example.com'),
                                 ('teacher6@example.com'),
                                 ('teacher7@example.com'),
                                 ('teacher8@example.com'),
                                 ('teacher9@example.com'),
                                 ('teacher10@example.com');

-- Create 10 students
INSERT INTO students (email, status) VALUES
                                         ('student1@example.com', 'active'),
                                         ('student2@example.com', 'active'),
                                         ('student3@example.com', 'active'),
                                         ('student4@example.com', 'suspended'),
                                         ('student5@example.com', 'active'),
                                         ('student6@example.com', 'active'),
                                         ('student7@example.com', 'suspended'),
                                         ('student8@example.com', 'active'),
                                         ('student9@example.com', 'suspended'),
                                         ('student10@example.com', 'active');

-- Create relationships between students and teachers
INSERT INTO teacher_student (teacher_id, student_id) VALUES
                                                         (1, 1),
                                                         (1, 2),
                                                         (1, 3),
                                                         (1, 4),
                                                         (2, 1),
                                                         (2, 2),
                                                         (2, 3),
                                                         (2, 5),
                                                         (5, 5),
                                                         (6, 6),
                                                         (7, 7),
                                                         (8, 8),
                                                         (9, 9),
                                                         (10, 10);
