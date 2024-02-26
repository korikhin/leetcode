-- PostgreSQL

create or replace function NthHighestSalary(n int) returns table (salary int) as $$
begin
    if n < 1 then
        return;
    end if;

    return query (
        select distinct(e.salary)
        from public.Employee as e
        order by e.salary desc
        limit 1 
        offset n - 1
    );
end;
$$ language plpgsql;
