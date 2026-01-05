-- 群成员表
CREATE TABLE `group_members` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '群ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '成员ID',
  `role` tinyint(4) NOT NULL DEFAULT '1' COMMENT '角色:1-普通,2-管理员,3-群主',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:0-已退出,1-正常,2-已被踢,3-已禁言',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '群内昵称',
  `inviter_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀请人ID',
  `last_ack_msg_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后已读消息ID',
  `mute_end_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '禁言结束时间',
  `join_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '加入时间',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `leave_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '离开时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_group_user` (`group_id`, `user_id`),
  KEY `idx_user` (`user_id`),
  KEY `idx_group_status` (`group_id`, `status`),
  KEY `idx_role` (`group_id`, `role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群成员表';