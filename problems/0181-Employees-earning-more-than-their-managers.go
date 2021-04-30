package problems

// https://leetcode.com/problems/employees-earning-more-than-their-managers/

// select l.Name as Employee from Employee l join Employee r on l.ManagerId = r.Id AND l.Salary > r.Salary
