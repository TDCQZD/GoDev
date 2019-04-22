/*
Navicat MySQL Data Transfer

Source Server         : MySql
Source Server Version : 50711
Source Host           : localhost:3306
Source Database       : godb

Target Server Type    : MYSQL
Target Server Version : 50711
File Encoding         : 65001

Date: 2018-09-13 13:56:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `books`
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL,
  `author` varchar(100) DEFAULT NULL,
  `price` int(100) DEFAULT NULL,
  `sales` int(100) DEFAULT NULL,
  `stock` int(100) DEFAULT NULL,
  `img_path` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES ('48', '解忧杂货店', '东野圭吾', '27', '102', '98', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('49', '边城', '沈从文', '23', '102', '98', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('50', '中国哲学史', '冯友兰', '45', '101', '99', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('51', '忽然七日', ' 劳伦', '19', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('52', '苏东坡传', '林语堂', '19', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('53', '百年孤独', '马尔克斯', '30', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('54', '扶桑', '严歌苓', '20', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('55', '给孩子的诗', '北岛', '22', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('56', '为奴十二年', '所罗门', '17', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('57', '平凡的世界', '路遥', '55', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('58', '悟空传', '今何在', '14', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('59', '硬派健身', '斌卡', '31', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('60', '从晚清到民国', '唐德刚', '40', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('61', '三体', '刘慈欣', '57', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('62', '看见', '柴静', '20', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('63', '活着', '余华', '11', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('64', '小王子', '安托万', '19', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('65', '我们仨', '杨绛', '11', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('66', '生命不息,折腾不止', '罗永浩', '25', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('67', '皮囊', '蔡崇达', '24', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('68', '恰到好处的幸福', '毕淑敏', '16', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('69', '大数据预测', '埃里克', '37', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('70', '人月神话', '布鲁克斯', '56', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('71', 'C语言入门经典', '霍尔顿', '45', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('72', '数学之美', '吴军', '30', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('73', 'Java编程思想', '埃史尔', '71', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('74', '设计模式之禅', '秦小波', '20', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('75', '图解机器学习', '杉山将', '34', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('76', '艾伦图灵传', '安德鲁', '47', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('77', '教父', '马里奥普佐', '29', '100', '100', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('78', '完美世界', '辰东', '30', '11100', '200', 'static/img/default.jpg');
INSERT INTO `books` VALUES ('80', '盗墓笔记', '南派三叔', '40', '1000', '11', '/static/img/default.jpg');

-- ----------------------------
-- Table structure for `cart`
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `ID` varchar(100) NOT NULL,
  `ZSM` int(100) DEFAULT NULL COMMENT '用户购物车的总数量',
  `ZJE` double(100,2) DEFAULT NULL COMMENT '用户购物车的总金额',
  `YHID` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `yhid` (`YHID`),
  CONSTRAINT `yhid` FOREIGN KEY (`YHID`) REFERENCES `shopusers` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES ('69613058-b995-40ea-4ea5-d086b7fbf462', '1', '23.00', 'db1162d5-296a-4092-7da3-08f3886bd0a4');

-- ----------------------------
-- Table structure for `cart_item`
-- ----------------------------
DROP TABLE IF EXISTS `cart_item`;
CREATE TABLE `cart_item` (
  `ID` int(100) NOT NULL AUTO_INCREMENT,
  `BOOKID` int(100) DEFAULT NULL COMMENT '关联表books中bookid',
  `SM` int(100) DEFAULT NULL COMMENT 'book数量',
  `ZJE` double(100,2) DEFAULT NULL COMMENT 'book总金额',
  `CARTID` varchar(100) DEFAULT NULL COMMENT 'cartid 属于哪个购物车的购物项',
  PRIMARY KEY (`ID`),
  KEY `bookid` (`BOOKID`),
  KEY `cartid` (`CARTID`),
  CONSTRAINT `bookid` FOREIGN KEY (`BOOKID`) REFERENCES `books` (`id`),
  CONSTRAINT `cartid` FOREIGN KEY (`CARTID`) REFERENCES `cart` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of cart_item
-- ----------------------------
INSERT INTO `cart_item` VALUES ('1', '49', '1', '23.00', '69613058-b995-40ea-4ea5-d086b7fbf462');

-- ----------------------------
-- Table structure for `orders`
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `ID` varchar(100) NOT NULL,
  `STATE` int(11) DEFAULT NULL,
  `SJ` varchar(100) DEFAULT NULL,
  `COUNT` int(100) DEFAULT NULL,
  `AMOUNT` double(100,2) DEFAULT NULL,
  `YHID` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `accountid` (`YHID`),
  CONSTRAINT `accountid` FOREIGN KEY (`YHID`) REFERENCES `shopusers` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES ('28842619-7aed-4dca-5d75-08fad12a58fe', '1', '2018/09/12 09:34:35', '2', '50.00', 'db1162d5-296a-4092-7da3-08f3886bd0a4');
INSERT INTO `orders` VALUES ('3caa7670-6b38-4649-6196-c8168883194e', '1', '2018/09/12 09:35:06', '3', '95.00', '7c6b6ded-0dc1-4171-660c-9cbb21e0f206');

-- ----------------------------
-- Table structure for `order_item`
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `TITLE` varchar(100) DEFAULT NULL,
  `AUTHOR` varchar(100) DEFAULT NULL,
  `PRICE` double(10,2) DEFAULT NULL,
  `IMG` varchar(100) DEFAULT NULL,
  `COUNT` int(100) DEFAULT NULL,
  `AMOUNT` double(100,2) DEFAULT NULL,
  `ORDERID` varchar(100) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `orderid` (`ORDERID`),
  CONSTRAINT `orderid` FOREIGN KEY (`ORDERID`) REFERENCES `orders` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order_item
-- ----------------------------
INSERT INTO `order_item` VALUES ('1', '解忧杂货店', '东野圭吾', '27.00', 'static/img/default.jpg', '1', '27.00', '28842619-7aed-4dca-5d75-08fad12a58fe');
INSERT INTO `order_item` VALUES ('2', '边城', '沈从文', '23.00', 'static/img/default.jpg', '1', '23.00', '28842619-7aed-4dca-5d75-08fad12a58fe');
INSERT INTO `order_item` VALUES ('3', '中国哲学史', '冯友兰', '45.00', 'static/img/default.jpg', '1', '45.00', '3caa7670-6b38-4649-6196-c8168883194e');
INSERT INTO `order_item` VALUES ('4', '边城', '沈从文', '23.00', 'static/img/default.jpg', '1', '23.00', '3caa7670-6b38-4649-6196-c8168883194e');
INSERT INTO `order_item` VALUES ('5', '解忧杂货店', '东野圭吾', '27.00', 'static/img/default.jpg', '1', '27.00', '3caa7670-6b38-4649-6196-c8168883194e');

-- ----------------------------
-- Table structure for `session`
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `UUID` varchar(255) NOT NULL,
  `USERID` varchar(255) NOT NULL,
  `USERNAME` varchar(255) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `userid` (`USERID`),
  KEY `userName` (`USERNAME`),
  CONSTRAINT `userName` FOREIGN KEY (`USERNAME`) REFERENCES `shopusers` (`NAME`),
  CONSTRAINT `userid` FOREIGN KEY (`USERID`) REFERENCES `shopusers` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of session
-- ----------------------------

-- ----------------------------
-- Table structure for `shopusers`
-- ----------------------------
DROP TABLE IF EXISTS `shopusers`;
CREATE TABLE `shopusers` (
  `ID` varchar(100) NOT NULL,
  `NAME` varchar(100) DEFAULT NULL,
  `PWD` varchar(100) DEFAULT NULL,
  `EMAIL` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `name` (`NAME`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of shopusers
-- ----------------------------
INSERT INTO `shopusers` VALUES ('7c6b6ded-0dc1-4171-660c-9cbb21e0f206', 'qqqq', '111111', '11@11.com');
INSERT INTO `shopusers` VALUES ('db1162d5-296a-4092-7da3-08f3886bd0a4', 'tttt', '111111', '11@11.com');
