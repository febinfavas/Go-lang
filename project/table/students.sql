CREATE TABLE students (student_id SERIAl,name VARCHAR(50),CGPA DECIMAL(4,2));

INSERT INTO students(NAME,CGPA) VALUES ('febin',8.52);
INSERT INTO students(NAME,CGPA) VALUES ('favas',10.00);

SELECT * FROM students;

INSERT INTO students(NAME,CGPA) VALUES ('ashiq',9.52);
INSERT INTO students(NAME,CGPA) VALUES ('saleem',9.00);
INSERT INTO students(NAME,CGPA) VALUES ('salin',9.80);
INSERT INTO students(NAME,CGPA) VALUES ('farha',7.80);
INSERT INTO students(NAME,CGPA) VALUES ('fahmi',8.80);

select * from students;