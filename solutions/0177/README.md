## 177. Nth Highest Salary

Write a solution to find the n-th highest salary from the `Employee` table. If there is no n-th highest salary, return `null`.

The result format is in the following example.

Table: `Employee`

```
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+

The id column is the primary key (column with unique values) for this table.
Each row of this table contains information about the salary of an employee.
```

<br>

### Example 1

```
Input: n = 2

Employee table:

+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+

Output:

+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
```

### Example 2

```
Input: n = 2

Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+

Output:

+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| null                   |
+------------------------+
```
