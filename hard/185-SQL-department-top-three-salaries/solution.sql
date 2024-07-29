SELECT 
    dep.name AS Department, 
    emp.name AS Employee, 
    emp.salary
    -- emp.salaryRank 
FROM (
    SELECT 
        employee.*, 
        DENSE_RANK() OVER (PARTITION BY departmentid ORDER BY salary DESC) AS salaryRank
    FROM employee
) AS emp
LEFT JOIN department AS dep ON dep.id = emp.departmentid
WHERE emp.salaryRank <= 3