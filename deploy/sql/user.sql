CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
    `password` varchar(128) NOT NULL DEFAULT '' COMMENT '加密后的密码',
    `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0:未知 1:男 2:女',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';