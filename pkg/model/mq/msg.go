package mq

// Redis 频道名称，所有网关都订阅这个 Key
const PushMsgKey = "easychat_msg_push_topic"

// BroadcastMsg 推送消息体
type BroadcastMsg struct {
	Type    int    `json:"type"`    // 消息类型: 1-单聊, 2-群聊, 3-系统通知
	UserId  int64  `json:"userId"`  // 目标用户ID
	Content string `json:"content"` // 消息内容 (或者用 []byte)
}
