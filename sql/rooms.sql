CREATE TABLE rooms (
	id serial NOT NULL PRIMARY KEY,
    user_id integer NOT NULL,
	name varchar(255) NOT NULL UNIQUE
);