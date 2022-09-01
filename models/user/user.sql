DROP TABLE IF EXISTS `account`;

CREATE TABLE `account`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT '表ID',
    `uid` varchar(255) NOT NULL COMMENT '用户ID',
    `name` varchar(255) NOT NULL COMMENT '用户名',
    `email` varchar(255) NOT NULL COMMENT '邮箱',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `avatar_url` varchar(255) NOT NULL COMMENT '头像'
        PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;