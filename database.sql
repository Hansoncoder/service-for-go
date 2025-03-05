/*
 Navicat Premium Data Transfer

 Source Server         : gotest
 Source Server Type    : MySQL
 Source Server Version : 90200 (9.2.0)
 Source Host           : localhost:3306
 Source Schema         : gotest

 Target Server Type    : MySQL
 Target Server Version : 90200 (9.2.0)
 File Encoding         : 65001

 Date: 05/03/2025 15:19:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `username`, `password`) VALUES (1, 'test', '$2a$10$3ggTkpfNJrE3PNGN7Rk3JeL16WAWSvPxvMtYPph4RRbPdR4M8j782');
INSERT INTO `users` (`id`, `username`, `password`) VALUES (2, 'test1', '$2a$10$NvYHKlKyUtBVDlDv4EldOe5n9NkQ1EhOhT4u07XJf/mwj7Tm4AhVW');
INSERT INTO `users` (`id`, `username`, `password`) VALUES (4, 'test2', '$2a$10$fZUCik7H7ZUQjpHvMvcQ6eD/qjbUAobCeR2KiEiXYG7.vtyXb1oQW');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
