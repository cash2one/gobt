#
# SQL Export
# Created by Querious (962)
# Created: 2016年8月8日 GMT+8下午12:27:05
# Encoding: Unicode (UTF-8)
#

-- CREATE DATABASE IF NOT EXISTS torrent DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
-- USE torrent

CREATE TABLE `history` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `keyword` varchar(45) DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent0` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent1` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent2` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent3` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent4` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent5` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent6` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent7` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent8` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrent9` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrenta` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrentb` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrentc` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrentd` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrente` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `torrentf` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `data` varchar(1024) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `infohash` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `infohash` varchar(40) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `infohash` (`infohash`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


CREATE TABLE `recommend` (
  `id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
