
-- Run in madast database

create schema madast;

create table madast.users
(
	id text not null
		constraint users_pk
			primary key,
	email text
		constraint users_email_key
			unique,
	firstname text,
	lastname text
);

alter table madast.users owner to postgres;

