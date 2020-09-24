SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE reddit;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `login` varchar(200) NOT NULL,
    `password` varchar(200) NOT NULL,
    PRIMARY KEY (`id`),
    INDEX (`login`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
    `token` varchar(200) NOT NULL,
    `id` varchar(200) NOT NULL,
    `login` varchar(200) NOT NULL,
    `expire` varchar(200) NOT NULL,
    PRIMARY KEY (`login`),
    INDEX (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;