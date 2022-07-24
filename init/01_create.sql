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
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "蓄電池のスペルを答えよ 3-1", "storage battery", "storge battery", "storage battery", "strage battery", "storege battery");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "発電のスペルを答えよ 3-2", "power generation", "power generate", "power generating", "power generation", "power generated");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(1, "変電所のスペルを答えよ 3-3", "substation", "substation", "life span", "rectify", "fluctuations");

INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("5400", "黒木さん", "broccoli", "2022/07/24", "後期期末試験");	
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "OSPFの正式名称 4-1", "Open Shortest Path First", "Open Start Protocol Fuck", "Open Short Protocol Friends", "Opening Sample Prototype Fast", "Open Shortest Path First");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "STPの正式名称 4-2", "Spanning Tree Protocol", "Spanning Tree Protocol", "Special Thanks Price", "Smart Telephone Protocol", "Step Trun Pumpkin");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(2, "DHCPの正式名称 4-3", "Dynamic Host Configuration Protocol", "Dynamic Host Configuration Protocol", "Dance Hop Control Protocol", "Dream Hyper Clone Proctocol", "Dokidoki Hassuru Coreppoi Protocol");

INSERT INTO rooms(title, subtitle, author, date, comment) 
			VALUES("日本語検定", "5級", "えらいひと", "2022/07/24", "生まれは宇宙");	
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(3, "朝起きたら？", "パンツ", "パン", "ごはん", "パンツ", "フライパン");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(3, "ここはどこ？わたしはだれ？", "チャーリーとチョコレート工場", "チャーリーとチョコレート工場", "穴と雪の女王", "すもも", "諦めたらここで試合終了ですよ");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(3, "人類は進化した", "ネッコ", "ワンコ", "ニャンコ", "ネッコ", "イッヌ");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(3, "某で前日招集がかかったよ", "私は行かないがお前は行け", "いく", "いかない", "私は行かないがお前は行け", "私は行くがお前は来るな");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(3, "10月の天気がいい日（気温22℃前後）で東京朝9時出発、車2時間以内で出かける場所と言えば？", "752", "宇都宮", "北海道", "ノストラダムス", "752");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
VALUES(3, "魔人拳でガークラしたよ", "帰る", "帰る", "帰る", "帰る", "帰る");
INSERT INTO quizzes(room_id, statement, answer, choice1, choice2, choice3, choice4)
