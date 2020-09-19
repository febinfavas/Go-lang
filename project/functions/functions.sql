SELECT * FROM students;

SELECT * FROM students WHERE student_id BETWEEN 3 AND 5;
SELECT * FROM students WHERE student_id NOT BETWEEN 3 AND 5;

SELECT * FROM students ORDER BY student_id;
SELECT * FROM students ORDER BY student_id DESC;

SELECT * FROM students WHERE cgpa IS DISTINCT FROM 9.00;
SELECT * FROM students WHERE cgpa IS NOT DISTINCT FROM 9.00;

SELECT abs(-17.3);
SELECT sqrt(25);

SELECT ascii('A');
SELECT ascii('1');

SELECT * FROM students LIMIT 5;
SELECT * FROM students OFFSET 3 LIMIT 2;
SELECT * FROM students OFFSET 3 LIMIT 5;

SELECT * FROM students WHERE name LIKE 'f%';

SELECT cgpa FROM students GROUP BY cgpa;
SELECT cgpa,COUNT(*) FROM students GROUP BY cgpa;
SELECT cgpa,COUNT(*) FROM students Group By cgpa HAVING COUNT(*)>1;
SELECT cgpa,COUNT(*) FROM students Group By cgpa HAVING COUNT(*)>1 ORDER BY cgpa;

SELECT MAX(cgpa) FROM students ;
SELECT MIN(cgpa) FROM students ;
SELECT AVG(cgpa) FROM students ;
SELECT ROUND(AVG(cgpa)) FROM students;
SELECT ROUND(SUM(cgpa)) FROM students;

SELECT NOW();
SELECT EXTRACT(MONTH FROM NOW());


