select
    e.name as "Employee"
from
    public.Employee as e
where
    exists (
        select
            null
        from
            public.Employee as m
        where
            e.managerId = m.id
            and e.salary > m.salary
    );
