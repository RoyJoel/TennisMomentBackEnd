-- Active: 1677501433605@@localhost@3306@tennismoment_app
CREATE TABLE `Player`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`login_name`   VARCHAR(255) NOT NULL,
	`name`        varchar(255) NOT NULL,
	`icon`        MEDIUMBLOB NOT NULL,
	`sex`         VARCHAR(255) NOT NULL,
	`age`         int(11)  NOT NULL,
	`years_played` int(11)  NOT NULL,
	`height`      DOUBLE NOT NULL,
	`width`       DOUBLE NOT NULL,
	`grip`        VARCHAR(255) NOT NULL,
	`backhand`    VARCHAR(255) NOT NULL,
	`points` int(11) NOT null,
	`is_adult` BOOLEAN NOT NULL,
	`career_stats_id` INT(11) NOT NULL
);

CREATE TABLE `Commodity`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`intro` varchar(255) NOT NULL,
	`price`  DOUBLE NOT NULL,
	`limit` int(11)  NOT NULL,
	`cag` int(11)  NOT NULL
);

CREATE TABLE `Bill`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `com_id` BIGINT  NOT NULL,
    `quantity` int(11)  NOT NULL,
    `opinion_id` BIGINT  NOT NULL,
	`order_id` BIGINT  NOT NULL
);

CREATE TABLE `Option`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`com_id` BIGINT  NOT NULL,
	`image` VARCHAR(255) NOT NULL,
	`intro` varchar(255) NOT NULL,
);

CREATE TABLE `Order`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`shipping_address_id` BIGINT  NOT NULL,
	`delivery_address_id` BIGINT  NOT NULL,
	`create_time` DOUBLE NOT NULL,
	`payed_time` DOUBLE NOT NULL,
	`completed_time` DOUBLE NOT NULL,
	`state` int(11)  NOT NULL
);

CREATE TABLE `Address`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`name` varchar(255) NOT NULL,
	`sex` varchar(255) NOT NULL,
	`phone_number` VARCHAR(255) NOT NULL,
	`province` VARCHAR(255) NOT NULL,
	`city` VARCHAR(255) NOT NULL,
	`area` VARCHAR(255) NOT NULL,
	`detailed_address` VARCHAR(255) NOT NULL
);

CREATE TABLE `Cart`
(
	`player_id` BIGINT NOT NULL,
	`order_id` BIGINT NOT NULL
);

CREATE TABLE `Relationship`
(
	`player1_id` BIGINT NOT NULL,
	`player2_id` BIGINT NOT NULL,
	INDEX (`player1_id`),
	INDEX (`player2_id`)
);

CREATE TABLE `Stats`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL, 
	`aces`         INT(11) NOT NULL,
	`double_faults` INT(11) NOT NULL,
	`serve_points` INT(11) NOT NULL,
	`first_serve_points` INT(11) NOT NULL,
	`first_serve_points_in` INT(11) NOT NULL,
	`first_serve_points_won` INT(11) NOT NULL,
	`second_serve_points_in` INT(11) NOT NULL,
	`second_serve_points_won` INT(11) NOT NULL,
	`break_points_faced` INT(11) NOT NULL,
	`break_points_saved` INT(11) NOT NULL,
	`serve_games_played` INT(11) NOT NULL,
	`serve_games_won` INT(11) NOT NULL,
	`return_aces` INT(11) NOT NULL,
	`return_serve_points` INT(11) NOT NULL,
	`first_serve_return_points` INT(11) NOT NULL,
	`first_serve_return_points_in` INT(11) NOT NULL,
	`first_serve_return_points_won` INT(11) NOT NULL,
	`second_serve_return_points_in` INT(11) NOT NULL,
	`second_serve_return_points_won` INT(11) NOT NULL,
	`break_points_opportunities` INT(11) NOT NULL,
	`break_points_converted` INT(11) NOT NULL,
	`return_games_played` INT(11) NOT NULL,
	`return_games_won` INT(11) NOT NULL,
	`net_points` INT(11) NOT NULL,
	`unforced_errors` INT(11) NOT NULL,
	`forehand_winners` INT(11) NOT NULL,
	`backhand_winners` INT(11) NOT NULL
 );

 CREATE TABLE `PlayerStats`
(
`player_id` INT(11) NOT NULL,
	`stats_id` INT(11) NOT NULL,
	INDEX (`player_id`),
	INDEX (`stats_id`)
);

CREATE TABLE `Game`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`date`       DOUBLE NOT NULL,
	`place`      VARCHAR(255) NOT NULL,
	`surface` VARCHAR(255) NOT NULL,
	`set_num` int(11) NOT NULL,
	`game_num` INT(11) NOT NULL,
	`round` INT(11) NOT NULL,
	`is_golden_goal` BOOLEAN NOT NULL,
	`is_player1_serving` BOOLEAN NOT NULL,
	`is_player1_left` BOOLEAN NOT NULL,
	`is_change_position` BOOLEAN NOT NULL,
	`is_completed` BOOLEAN NOT NULL,
	`player1_id` INT(11) NOT NULL,
	`player1_stats_id` INT(11) NOT NULL,
	`player2_id` INT(11) NOT NULL,
	`player2_stats_id` INT(11) NOT NULL,
	`is_player1_first_serve` INT(11) NOT NULL,
	`is_player2_first_serve` INT(11) NOT NULL,
	`result` JSON not null
);

CREATE TABLE `Club`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`name`   VARCHAR(255) NOT NULL,
	`icon`        MEDIUMBLOB NOT NULL,
	`intro`         VARCHAR(255) NOT NULL,
	`owner_id`         BIGINT  NOT NULL,
	`address`        VARCHAR(255) NOT NULL
);

CREATE TABLE `Schedule`
(
	`date`       DOUBLE NOT NULL,
	`place`      VARCHAR(255) NOT NULL,
	`player1_id` INT(11) NOT NULL,
	`player2_id` INT(11) NOT NULL
);

CREATE TABLE `Event`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`name`   VARCHAR(255) NOT NULL,
	`icon`        MEDIUMBLOB NOT NULL,
	`start_date`         DOUBLE NOT NULL,
	`end_date`         DOUBLE  NOT NULL,
	`level`        INT(11) NOT NULL,
	`draw`  JSON not null,
	`schedule`  JSON not null
);

CREATE TABLE `ClubEvent`
(
`club_id` INT(11) NOT NULL,
	`event_id` INT(11) NOT NULL,
	INDEX (`club_id`),
	INDEX (`event_id`)
);

 CREATE TABLE `EventGame`
(
`game_id` INT(11) NOT NULL,
	`event_id` INT(11) NOT NULL,
	INDEX (`game_id`),
	INDEX (`event_id`)
);

 CREATE TABLE `EventPlayer`
(
`player_id` INT(11) NOT NULL,
	`event_id` INT(11) NOT NULL,
	INDEX (`player_id`),
	INDEX (`event_id`)
);

 CREATE TABLE `ClubManager`
(
`club_id` INT(11) NOT NULL,
	`manager_id` INT(11) NOT NULL,
	INDEX (`club_id`),
	INDEX (`manager_id`)
);

 CREATE TABLE `ClubMember`
(
`club_id` INT(11) NOT NULL,
	`member_id` INT(11) NOT NULL,
	INDEX (`club_id`),
	INDEX (`member_id`)
);