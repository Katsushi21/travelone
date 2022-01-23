CREATE TABLE IF NOT EXISTS `article` (
  `id` int NOT NULL,
  `uid` smallint NOT NULL,
  `title` varchar(256) NOT NULL,
  `thumbnail` varchar(256) NOT NULL,
  `address` varchar(256) NOT NULL,
  `conditions` json NOT NULL,
  `title` varchar(256) NOT NULL,
  `thumbnail` varchar(256) NOT NULL,
  `comment1` text NOT NULL,
  `image1` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `account` (
  `id` smallint NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(256) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `birth` varchar(12) NOT NULL NOT NULL,
  `history` varchar(4) NOT NULL,
  `avatar` varchar(256) NOT NULL NOT NULL,
  `favorite` text NOT NULL,
  `introduction` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `marker` (
  `id` smallint NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(256) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `birth` varchar(12) NOT NULL NOT NULL,
  `history` varchar(4) NOT NULL,
  `avatar` varchar(256) NOT NULL NOT NULL,
  `favorite` text NOT NULL,
  `introduction` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;