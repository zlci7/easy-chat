CREATE TABLE `msg` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `msg_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务全局唯一ID(UUID)',
  `from_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送者ID',
  `to_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '接收者ID',
  `group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '群ID',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '消息类型',
  `content` varchar(2048) NOT NULL DEFAULT '' COMMENT '消息内容',
  
  `seq` bigint(20)  NOT NULL DEFAULT '0' COMMENT '会话内消息序列号(单聊/群聊内递增)',
  
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送时间(ms)',
  
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_msg_id` (`msg_id`),
  
  -- 优化：用于接收方拉取离线消息 (Sync)
  KEY `idx_to_uid_seq` (`to_uid`, `seq`), 
  
  KEY `idx_from_uid` (`from_uid`),
  KEY `idx_group_id` (`group_id`) -- 如果做群聊，建议加上这个索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息表';