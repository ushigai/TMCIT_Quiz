CREATE DATABASE if not exists go_database;

use go_database;


CREATE TABLE if not exists users (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(255),
	email VARCHAR(255),
	PRIMARY KEY (id)
);

CREATE TABLE if not exists rooms (
	id INT not null AUTO_INCREMENT,
	title VARCHAR(255),
	subtitle VARCHAR(255),
	author VARCHAR(255),
	date VARCHAR(255),
	comment VARCHAR(255),
	PRIMARY KEY (id)
);

INSERT INTO users(id, name, email) VALUES(1, "Yamada", "yamada@example.com");
INSERT INTO users(id, name, email) VALUES(2, "Tanaka", "tanaka@example.com");

INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("1300", "seiji", "ushigai", "2022/07/01", "nandayona");
INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("1300", "kids", "shoma", "2022/07/02", "shoma kami");
