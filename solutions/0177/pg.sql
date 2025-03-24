CREATE OR REPLACE FUNCTION NthHighestSalary (n integer)
RETURNS TABLE (
    salary integer
)
AS $$
BEGIN
    IF n IS NULL OR n < 1 THEN
        RETURN;
    END IF;

    RETURN QUERY (
        SELECT DISTINCT(e.salary)
        FROM public.Employee AS e
        ORDER BY e.salary DESC
        LIMIT 1
        OFFSET n - 1
    );
END;
$$ LANGUAGE plpgsql;
