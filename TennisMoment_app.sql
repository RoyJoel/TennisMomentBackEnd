-- Active: 1677501433605@@116.62.121.33@3306@tennismoment_app
CREATE TABLE `Player`
(
	`login_name`   VARCHAR(255) PRIMARY KEY NOT NULL,
	`name`        varchar(255) NOT NULL,
	`icon`        MEDIUMBLOB NOT NULL,
	`sex`         VARCHAR(255) NOT NULL,
	`age`         int(11)  NOT NULL,
	`years_played` int(11)  NOT NULL,
	`height`      DECIMAL(20) NOT NULL,
	`width`       DECIMAL(20) NOT NULL,
	`grip`        VARCHAR(255) NOT NULL,
	`backhand`    VARCHAR(255) NOT NULL,
	`point` int(11) NOT null,
	`is_adult` BOOLEAN NOT NULL,
	`career_stats_id` INT(11) NOT NULL,
	`friends` VARCHAR(255) NOT null,

	Foreign Key (`friends`) REFERENCES Relationship(`login_name`)
);

CREATE TABLE `Relationship`
(
	`login_name` VARCHAR(255) NOT NULL,
	`friend_login_name` VARCHAR(255) NOT NULL,
	INDEX (`friend_login_name`)
);

CREATE TABLE `Stats`
(
	`id` INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL, 
	`aces`         INT(11) NOT NULL,
	`double_faults` INT(11) NOT NULL,
	`first_serve_points` INT(11) NOT NULL,
	`first_serve_points_in` INT(11) NOT NULL,
	`first_serve_points_won` INT(11) NOT NULL,
	`second_serve_points` INT(11) NOT NULL,
	`second_serve_points_won` INT(11) NOT NULL,
	`break_points_faced` INT(11) NOT NULL,
	`break_points_saved` INT(11) NOT NULL,
	`serve_games_played` INT(11) NOT NULL,
	`serve_games_won` INT(11) NOT NULL,
	`total_serve_points_won` INT(11) NOT NULL,
	`first_serve_return_points` INT(11) NOT NULL,
	`first_serve_return_aces` INT(11) NOT NULL,
	`first_serve_return_points_won` INT(11) NOT NULL,
	`second_serve_return_points` INT(11) NOT NULL,
	`second_serve_return_aces` INT(11) NOT NULL,
	`second_serve_return_points_won` INT(11) NOT NULL,
	`break_points_opportunities` INT(11) NOT NULL,
	`break_points_converted` INT(11) NOT NULL,
	`return_games_played` INT(11) NOT NULL,
	`return_games_won` INT(11) NOT NULL,
	`total_return_points_won`INT(11) NOT NULL,
	`total_points_won` INT(11) NOT NULL,
	`net_points` INT(11) NOT NULL,
	`unforced_errors` INT(11) NOT NULL,
	`forehand_winners` INT(11) NOT NULL,
	`backhand_winners` INT(11) NOT NULL,
	`player_login_name` VARCHAR(255) NOT NULL UNIQUE,

	Foreign Key (`player_login_name`) REFERENCES Player(`login_name`)

 )AUTO_INCREMENT = 0;

 CREATE TABLE `PlayerStats`
(
`id` INT(11) PRIMARY KEY NOT NULL,
	`login_name` VARCHAR(255) NOT NULL,
	INDEX (`id`)
);

CREATE TABLE `Game`
(
	`date`       BIGINT NOT NULL,
	`place`      VARCHAR(255) NOT NULL,
	`surface` VARCHAR(255) NOT NULL,
	`set_num` int(11) NOT NULL,
	`game_num` INT(11) NOT NULL,
	`is_golden_goal` BOOLEAN NOT NULL,
	`is_completed` BOOLEAN NOT NULL,
	`player1_login_name` VARCHAR(255) NOT NULL,
	`player1_stats_id` INT(11) NOT NULL,
	`player2_login_name` VARCHAR(255) NOT NULL,
	`player2_stats_id` INT(11) NOT NULL,
	`result` JSON not null,

	Foreign Key (`player1_login_name`) REFERENCES Player(`login_name`),
	Foreign Key (`player1_stats_id`) REFERENCES Stats(`id`),
	Foreign Key (`player2_login_name`) REFERENCES Player(`login_name`),
	Foreign Key (`player2_stats_id`) REFERENCES Stats(`id`)
);

/* CREATE TABLE `Player_Game_Relation`
{
	Foreign Key (`Player_id`) REFERENCES Player(`id`)
	Foreign Key (`Game_id`) REFERENCES Game(`id`)
}


CREATE TABLE `Relation_Between_Player_Relation`
{
	Foreign Key () REFERENCES ()
} */
