package query

const UpsertUserName = `
	insert into users(first_name) 
	values ($1)
	on conflict (first_name) do update set first_name = excluded.first_name
	where id = $2;
`

const UpsertUserSurname = `
	insert into users(surname) 
	values ($1)
	on conflict (surname) do update set surname = excluded.surname
	where id = $2;
`

const UpsertUserPatronymic = `
	insert into users(patronymic) 
	values ($1)
	on conflict (patronymic) do update set patronymic = excluded.patronymic
    where id = $2;
`

const UpsertUserCity = `
	insert into users(city) 
	values ($1)
	on conflict (city) do update set city = excluded.city
    where id = $2;
`

const UpsertUserAge = `
	insert into users(age) 
	values ($1)
	on conflict (age) do update set age = excluded.age
    where id = $2;
`

const UpsertUserUniversity = `
	insert into users(university) 
	values ($1)
	on conflict (university) do update set university = excluded.university
    where id = $2;
`

const UpsertUserEducation = `
	insert into users(education) 
	values ($1)
	on conflict (education) do update set education = excluded.education
    where id = $2;
`
