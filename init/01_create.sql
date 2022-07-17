CREATE DATABASE if not exists go_database;

use go_database;

DROP TABLE IF EXISTS users;

CREATE TABLE if not exists users (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(255),
	email VARCHAR(255),
	PRIMARY KEY (id)
);

CREATE TABLE if not exists RoomData (
	id INT not null AUTO_INCREMENT,
	title VARCHAR(255),
	subtitle VARCHAR(255),
	PRIMARY KEY (id)
);

INSERT INTO users(id, name, email) VALUES(1, "Yamada", "yamada@example.com");
INSERT INTO users(id, name, email) VALUES(2, "Tanaka", "tanaka@example.com");

INSERT INTO RoomData(title, subtitle) VALUES("1300", "seiji");
INSERT INTO RoomData(title, subtitle) VALUES("1300", "kids");
