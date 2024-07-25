/*
 Navicat Premium Data Transfer

 Source Server         : LIXIN
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : gowebtest

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 25/07/2024 10:49:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS `admin_info`;
CREATE TABLE `admin_info`  (
  `admin_id` int NOT NULL,
  `admin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `admin_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `lastlogin_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`admin_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_info
-- ----------------------------
INSERT INTO `admin_info` VALUES (10000001, 'lixin', '123456', NULL);

-- ----------------------------
-- Table structure for card_attribute
-- ----------------------------
DROP TABLE IF EXISTS `card_attribute`;
CREATE TABLE `card_attribute`  (
  `card_id` int NOT NULL,
  `card_probability` int NOT NULL,
  PRIMARY KEY (`card_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of card_attribute
-- ----------------------------

-- ----------------------------
-- Table structure for card_info
-- ----------------------------
DROP TABLE IF EXISTS `card_info`;
CREATE TABLE `card_info`  (
  `card_id` int NOT NULL,
  `card_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `card_level` int NOT NULL,
  `card_imageurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `card_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`card_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of card_info
-- ----------------------------
INSERT INTO `card_info` VALUES (10000001, '李子哥头像', 3, 'cardimage\\10000001.jpg', '李子哥常用头像，不值钱。');
INSERT INTO `card_info` VALUES (10000002, '秋哥头像', 3, 'cardimage\\10000002.jpg', '秋哥常用头像，不值钱');
INSERT INTO `card_info` VALUES (10000003, '周老师头像', 3, 'cardimage\\10000003.jpg', '周老师常用头像，不值钱');
INSERT INTO `card_info` VALUES (10000004, '二次元美图', 4, 'cardimage\\10000004.jpg', '电报上的二次元美图，好看。');

-- ----------------------------
-- Table structure for user_card_library
-- ----------------------------
DROP TABLE IF EXISTS `user_card_library`;
CREATE TABLE `user_card_library`  (
  `user_id` int NOT NULL,
  `card_id` int NOT NULL,
  `card_num` int NOT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_card_library
-- ----------------------------

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `user_id` int NOT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_level` int NOT NULL,
  `user_balance` int NOT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (10000001, 'lixin', '12345lixin', 0, 0);
INSERT INTO `user_info` VALUES (10000002, 'tom', '12345tom', 0, 0);
INSERT INTO `user_info` VALUES (10000003, 'Zhuo', '123456', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
