package mq

// Redis 频道名称，所有网关都订阅这个 Key
const PushMsgKey = "easychat_msg_push_topic"

// BroadcastMsg 推送消息体
type BroadcastMsg struct {
	Type       int    `json:"type"`       // 消息类型: 1-单聊, 2-群聊
	FromUserId int64  `json:"fromUserId"` // 发送人ID ← 新增
	ToUserId   int64  `json:"toUserId"`   // 接收人ID（改名）
	Content    string `json:"content"`    // 消息内容
	Timestamp  int64  `json:"timestamp"`  // 时间戳 ← 新增
	MsgId      string `json:"msgId"`      // 消息ID ← 新增
}
