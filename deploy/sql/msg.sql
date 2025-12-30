CREATE TABLE `msg` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `msg_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务ID',
  `from_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送者ID',
  `to_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '接收者ID',
  `group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '群ID',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '消息类型',
  `content` varchar(2048) NOT NULL DEFAULT '' COMMENT '消息内容',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送时间(ms)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_msg_id` (`msg_id`),
  KEY `idx_to_uid` (`to_uid`), 
  KEY `idx_from_uid` (`from_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息表';