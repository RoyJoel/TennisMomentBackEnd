CREATE TABLE `Player`
(
	`id`          int(11) NOT NULL AUTO_INCREMENT,
	`icon`        MEDIUMBLOB NOT NULL,
	`age`         int(11)  NOT NULL,
    `name`        varchar(255) DEFAULT NULL,
    `info_key`    varchar(255) NOT NULL,
	`yearsPlayed` int(11)  NOT NULL,
	`height`      DECIMAL(20) NOT NULL,
	`width`       DECIMAL(20) NOT NULL,
	`grip`        VARCHAR(255) NOT NULL,
	`backHand`    VARCHAR(255) NOT NULL,
	
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

CREATE TABLE `Stats`
{
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
	`returnPointsWon`INT(11) NOT NULL,
	`netPoints` INT(11) NOT NULL,
	`unforcedErrors` INT(11) NOT NULL,
	`forehandWinners` INT(11) NOT NULL,
	`backhandWinners` INT(11) NOT NULL,

	Foreign Key (player_id) REFERENCES Player(id)

	UNIQUE(player_id)
} ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

CREATE TABLE `Game`
{
	`id`         INT(11) NOT NULL AUTO_INCREMENT,
	`date`       DATE NOT NULL
	`place`      VARCHAR(255) NOT NULL,

	Foreign Key (`winner`) REFERENCES Player(`id`)
	Foreign Key (winnerStats) REFERENCES Stats(id)
	Foreign Key (`loser`) REFERENCES Player(`id`)
	Foreign Key (loserStats) REFERENCES Stats(id)

	UNIQUE(`winner`)
	UNIQUE('loser')
} ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

CREATE TABLE `Player_Game_Relation`
{
	Foreign Key (Player_id) REFERENCES Player(id)
	Foreign Key (Game_id) REFERENCES Game(id)
}


CREATE TABLE `Relation_Between_Player_Relation`
{
	Foreign Key () REFERENCES ()
}


 var name: String
    var icon: UIImage
    var sex: String
    var age: Int
    var yearsPlayed: Int
    var height: Float
    var width: Float
    var grip: String
    var backHand: String
    var careerStats: Stats
    var friends: [Player]
    var gamesPlayed: [Game]