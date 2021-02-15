CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	encrypted_password varchar(255) NOT NULL,
	email varchar(255) NOT NULL UNIQUE,
	token_jwt varchar(255) NULL
);