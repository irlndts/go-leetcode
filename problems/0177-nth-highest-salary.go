package problems

// https://leetcode.com/problems/nth-highest-salary/

/*
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  declare i int;
  set i = N - 1;
  RETURN (
      # Write your MySQL query statement below.
      select Salary from Employee group by Salary order by Salary desc limit i, 1
  );
END
*/
