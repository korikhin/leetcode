select
    max(salary) as SecondHighestSalary
from
    public.Employee
where
    salary < (
        select
            max(salary)
        from
            public.Employee
    )
