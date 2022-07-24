CREATE DATABASE if not exists tmcit_quiz_database;

use tmcit_quiz_database;


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
	choice4 VARCHAR(255),
	PRIMARY KEY (quiz_id)
);

INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("5400", "海上先生", "broccoli", "2022/07/24", "後期期末試験");	
INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("5400", "黒木さん", "broccoli", "2022/07/24", "後期期末試験");	

INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "蓄電池のスペルを答えよ 3-1", "storage battery", "storge battery", "storage battery", "strage battery", "storege battery");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "発電のスペルを答えよ 3-2", "power generation", "power generate", "power generating", "power generation", "power generated");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "変電所のスペルを答えよ 3-3", "substation", "substation", "life span", "rectify", "fluctuations");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "OSPFの正式名称 4-1", "Open Shortest Path First", "Open Start Protocol Fuck", "Open Short Protocol Friends", "Opening Sample Prototype Fast", "Open Shortest Path First");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "STPの正式名称 4-2", "Spanning Tree Protocol", "Spanning Tree Protocol", "Special Thanks Price", "Smart Telephone Protocol", "Step Trun Pumpkin");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "DHCPの正式名称 4-3", "Dynamic Host Configuration Protocol", "Dynamic Host Configuration Protocol", "Dance Hop Control Protocol", "Dream Hyper Clone Proctocol", "Dokidoki Hassuru Coreppoi Protocol");
