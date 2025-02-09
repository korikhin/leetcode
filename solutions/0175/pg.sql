-- PostgreSQL

select
    p.firstName,
    p.lastName,
    a.city,
    a.state
from public.Person as p
left join public.Address as a using(personId)
