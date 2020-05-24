-- Find all current managers of each department and display his/her title, first name, last name, current salary

SELECT t.title, e.first_name, e.last_name, s.salary
	FROM employees e
		LEFT JOIN titles t ON t.emp_no = e.emp_no
		LEFT JOIN salaries s ON s.emp_no = e.emp_no
	WHERE t.to_date > CURDATE()
		AND s.to_date > CURDATE();


-- Find all employees (department, title, first name, last name, hire date, how many years they have been working)
-- to congratulate them on their hire anniversary this month.

SELECT d.dept_name, t.title, e.first_name, e.last_name, hire_date,  YEAR(CURDATE()) - YEAR(hire_date) work_years
	FROM employees e
		LEFT JOIN titles t ON t.emp_no = e.emp_no
        LEFT JOIN dept_emp de ON de.emp_no = e.emp_no
        LEFT JOIN departments d ON de.dept_no = d.dept_no
    WHERE MONTH(hire_date) = MONTH(CURDATE())
		AND t.to_date > CURDATE()
		AND de.to_date > CURDATE();

-- Find all departments, their current employee count, their current sum salary.

SELECT d.dept_name, COUNT(de.emp_no) employees, SUM(s.salary) salaries
	FROM departments d
		LEFT JOIN dept_emp de ON de.dept_no = d.dept_no
        LEFT JOIN employees e ON e.emp_no = de.emp_no
        LEFT JOIN salaries s ON s.emp_no = e.emp_no
	WHERE s.to_date > CURDATE()
    GROUP BY d.dept_no;