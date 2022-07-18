CREATE DATABASE if not exists go_database;

use go_database;


CREATE TABLE if not exists rooms (
	id INT not null AUTO_INCREMENT,
	title VARCHAR(255),
	subtitle VARCHAR(255),
	author VARCHAR(255),
	date VARCHAR(255),
	comment VARCHAR(255),
	PRIMARY KEY (id)
);

CREATE TABLE if not EXISTS quizzes (
	quiz_id INT NOT NULL AUTO_INCREMENT,
	room_id INT,
	statement VARCHAR(255),
	answer VARCHAR(255),
	choice1 VARCHAR(255),
	choice2 VARCHAR(255),
	choice3 VARCHAR(255),
	choice4 VARCHAR(255)
)

INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("1300", "seiji", "ushigai", "2022/07/01", "nandayona");
INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("1300", "kids", "shoma", "2022/07/02", "shoma kami");

INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "test statement 1-1", "test answer", 
"test choice1", "test choice2", "test choice1", "test choice2");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "test statement 1-2", "test answer", 
"test choice1", "test choice2", "test choice1", "test choice2");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "test statement 2-1", "test answer", 
"test choice1", "test choice2", "test choice1", "test choice2");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "test statement 2-2", "test answer", 
"test choice1", "test choice2", "test choice1", "test choice2");
