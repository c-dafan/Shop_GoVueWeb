/*
Navicat MySQL Data Transfer

Source Server         : test
Source Server Version : 50715
Source Host           : localhost:3306
Source Database       : myshop

Target Server Type    : MYSQL
Target Server Version : 50715
File Encoding         : 65001

Date: 2019-04-04 18:51:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `aid` int(11) NOT NULL AUTO_INCREMENT,
  `address_jd` varchar(50) NOT NULL,
  `address_details` varchar(50) NOT NULL,
  `mail` varchar(30) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `name` varchar(20) NOT NULL,
  `uid` int(11) NOT NULL,
  PRIMARY KEY (`aid`),
  KEY `user` (`uid`),
  CONSTRAINT `user` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `cartid` int(11) NOT NULL AUTO_INCREMENT,
  `num` int(11) NOT NULL,
  `gid` int(11) NOT NULL,
  `uid` int(11) NOT NULL,
  PRIMARY KEY (`cartid`),
  KEY `userid` (`uid`),
  KEY `goodid` (`gid`),
  CONSTRAINT `goodid` FOREIGN KEY (`gid`) REFERENCES `good` (`gid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `userid` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `cid` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(50) NOT NULL,
  `time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `socre` tinyint(4) NOT NULL,
  `otid` int(11) NOT NULL,
  `gid` int(11) NOT NULL,
  PRIMARY KEY (`cid`),
  KEY `order_item` (`otid`),
  KEY `gid` (`gid`),
  CONSTRAINT `gid` FOREIGN KEY (`gid`) REFERENCES `good` (`gid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `order_item` FOREIGN KEY (`otid`) REFERENCES `order_item` (`otid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for good
-- ----------------------------
DROP TABLE IF EXISTS `good`;
CREATE TABLE `good` (
  `gid` int(11) NOT NULL AUTO_INCREMENT,
  `price` decimal(5,2) NOT NULL,
  `num` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `img` varchar(100) DEFAULT NULL,
  `sid` int(11) NOT NULL,
  `kind` tinyint(4) NOT NULL,
  PRIMARY KEY (`gid`),
  KEY `seller` (`sid`),
  CONSTRAINT `seller` FOREIGN KEY (`sid`) REFERENCES `seller` (`sid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `oid` int(11) NOT NULL AUTO_INCREMENT,
  `time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `uid` int(11) NOT NULL,
  `aid` int(11) NOT NULL,
  PRIMARY KEY (`oid`),
  KEY `uid` (`uid`),
  KEY `aid` (`aid`),
  CONSTRAINT `aid` FOREIGN KEY (`aid`) REFERENCES `address` (`aid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `uid` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item` (
  `otid` int(11) NOT NULL AUTO_INCREMENT,
  `num` int(11) NOT NULL,
  `oid` int(11) NOT NULL,
  `gid` int(11) NOT NULL,
  `cid` int(11) DEFAULT NULL,
  PRIMARY KEY (`otid`),
  KEY `order` (`oid`),
  KEY `good` (`gid`),
  KEY `comment` (`cid`),
  CONSTRAINT `comment` FOREIGN KEY (`cid`) REFERENCES `comment` (`cid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `good` FOREIGN KEY (`gid`) REFERENCES `good` (`gid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `order` FOREIGN KEY (`oid`) REFERENCES `order` (`oid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for seller
-- ----------------------------
DROP TABLE IF EXISTS `seller`;
CREATE TABLE `seller` (
  `sid` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `begin_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `rank` tinyint(3) NOT NULL,
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `uid` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `aid` int(11) DEFAULT NULL,
  `phone` varchar(20) NOT NULL,
  PRIMARY KEY (`uid`),
  KEY `address` (`aid`),
  CONSTRAINT `address` FOREIGN KEY (`aid`) REFERENCES `address` (`aid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
