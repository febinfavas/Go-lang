SELECT * FROM students;

SELECT * FROM students WHERE cgpa = 10.00 AND student_id =2;
SELECT * FROM students WHERE cgpa = 10.00 OR cgpa = 9.00;
SELECT * FROM students WHERE NOT student_id =3;

SELECT * FROM students WHERE cgpa > 9.00;
SELECT * FROM students WHERE cgpa < 8.00;
SELECT * FROM students WHERE cgpa >= 9.00;
SELECT * FROM students WHERE cgpa <= 9.00;
SELECT * FROM students WHERE cgpa != 9.00;

SELECT 2+5;
SELECT 5-5;
SELECT 2*5;
SELECT 4/2;
SELECT * FROM students WHERE student_id = 1+2;
SELECT * FROM students WHERE student_id = 10/2;
SELECT 4 & 2;
SELECT 5!;




