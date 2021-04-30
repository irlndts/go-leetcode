package problems

// https://leetcode.com/problems/second-highest-salary/

// select ifnull((select distinct Salary from Employee group by Salary order by Salary desc limit 1, 1), null) as SecondHighestSalary
