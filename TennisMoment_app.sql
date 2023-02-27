CREATE TABLE Player
(
	`id`          int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`login_name`   VARCHAR(255) NOT NULL,
	`name`        varchar(255) NOT NULL,
	`icon`        MEDIUMBLOB NOT NULL,
	`sex`         VARCHAR(255) NOT NULL,
	`age`         int(11)  NOT NULL,
	`years_played` int(11)  NOT NULL,
	`height`      DECIMAL(20) NOT NULL,
	`width`       DECIMAL(20) NOT NULL,
	`grip`        VARCHAR(255) NOT NULL,
	`backHand`    VARCHAR(255) NOT NULL
);

CREATE TABLE `Stats`
(
	`aces`         INT(11) NOT NULL,
	`doubleFaults` INT(11) NOT NULL,
	`firstServePoints` INT(11) NOT NULL,
	`firstServePointsIn` INT(11) NOT NULL,
	`firstServePointsWon` INT(11) NOT NULL,
	`secondServePoints` INT(11) NOT NULL,
	`secondServePointsWon` INT(11) NOT NULL,
	`breakPointsFaced` INT(11) NOT NULL,
	`breakPointsSaved` INT(11) NOT NULL,
	`serveGamesPlayed` INT(11) NOT NULL,
	`serveGamesWon` INT(11) NOT NULL,
	`totalServePointsWon` INT(11) NOT NULL,
	`firstServeReturnPoints` INT(11) NOT NULL,
	`firstServeReturnAces` INT(11) NOT NULL,
	`firstServeReturnPointsWon` INT(11) NOT NULL,
	`secondServeReturnPoints` INT(11) NOT NULL,
	`secondServeReturnAces` INT(11) NOT NULL,
	`secondServeReturnPointsWon` INT(11) NOT NULL,
	`breakPointsOpportunities` INT(11) NOT NULL,
	`breakPointsConverted` INT(11) NOT NULL,
	`returnGamesPlayed` INT(11) NOT NULL,
	`returnGamesWon` INT(11) NOT NULL,
	`totalReturnPointsWon`INT(11) NOT NULL,
	`totalPointsWon` INT(11) NOT NULL,
	`netPoints` INT(11) NOT NULL,
	`unforcedErrors` INT(11) NOT NULL,
	`forehandWinners` INT(11) NOT NULL,
	`backhandWinners` INT(11) NOT NULL,
	`player_id` int(11) NOT NULL UNIQUE,

	Foreign Key (`player_id`) REFERENCES Player(`id`)

 );

/* CREATE TABLE `Game`
{
	`id`         int(11) NOT NULL AUTO_INCREMENT,
	`date`       DATE NOT NULL,
	`place`      VARCHAR(255) NOT NULL,

	Foreign Key (`winner_id`) REFERENCES Player(`id`)
	Foreign Key (`winnerStats`) REFERENCES Stats(`id`)
	Foreign Key (`loser`) REFERENCES Player(`id`)
	Foreign Key (`loserStats`) REFERENCES Stats(`id`)

	UNIQUE(`winner`)
	UNIQUE('loser')
} ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8; */

/* CREATE TABLE `Player_Game_Relation`
{
	Foreign Key (`Player_id`) REFERENCES Player(`id`)
	Foreign Key (`Game_id`) REFERENCES Game(`id`)
}


CREATE TABLE `Relation_Between_Player_Relation`
{
	Foreign Key () REFERENCES ()
} */
