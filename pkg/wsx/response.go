package wsx

import (
	"easy-chat/pkg/xerr"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

// WsMessage WebSocket 统一消息格式
type WsMessage struct {
	Action string      `json:"action"`         // 消息类型：error, msgAck, newMsg, typing, online, etc.
	Code   uint32      `json:"code,omitempty"` // 错误码（仅在 action=error 时使用）
	Msg    string      `json:"msg,omitempty"`  // 提示信息
	Data   interface{} `json:"data,omitempty"` // 数据载荷
	Time   int64       `json:"time,omitempty"` // 时间戳（可选）
}

// SendJSON 发送 JSON 消息（带错误处理）
func SendJSON(conn *websocket.Conn, msg *WsMessage) error {
	if conn == nil {
		return xerr.NewErrMsg("connection is nil")
	}
	return conn.WriteJSON(msg)
}

// SendSuccess 发送成功消息
// action: 动作类型，如 "msgAck", "operationSuccess"
// data: 可选的数据载荷
func SendSuccess(conn *websocket.Conn, action string, data interface{}) error {
	return SendJSON(conn, &WsMessage{
		Action: action,
		Code:   xerr.OK,
		Msg:    "success",
		Data:   data,
	})
}

// SendError 发送错误消息
// err: 错误对象，会自动解析 CodeError
func SendError(conn *websocket.Conn, err error) error {
	var errCode uint32 = xerr.SERVER_COMMON_ERROR
	var errMsg = "操作失败"

	// 尝试解析自定义错误
	if e, ok := err.(*xerr.CodeError); ok {
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	} else if err != nil {
		errMsg = err.Error()
	}

	return SendJSON(conn, &WsMessage{
		Action: "error",
		Code:   errCode,
		Msg:    errMsg,
	})
}

// SendErrorWithCode 发送指定错误码的错误消息
func SendErrorWithCode(conn *websocket.Conn, code uint32, msg string) error {
	return SendJSON(conn, &WsMessage{
		Action: "error",
		Code:   code,
		Msg:    msg,
	})
}

// SendMessage 发送新消息通知（用于接收方）
func SendMessage(conn *websocket.Conn, msgData interface{}) error {
	return SendJSON(conn, &WsMessage{
		Action: "newMsg",
		Data:   msgData,
	})
}

// SendTyping 发送正在输入状态
func SendTyping(conn *websocket.Conn, userId int64) error {
	return SendJSON(conn, &WsMessage{
		Action: "typing",
		Data: map[string]interface{}{
			"userId": userId,
		},
	})
}

// SendOnlineStatus 发送在线状态
func SendOnlineStatus(conn *websocket.Conn, userId int64, online bool) error {
	return SendJSON(conn, &WsMessage{
		Action: "onlineStatus",
		Data: map[string]interface{}{
			"userId": userId,
			"online": online,
		},
	})
}

// BroadcastMessage 向多个连接广播消息（带错误日志）
func BroadcastMessage(conns []*websocket.Conn, msg *WsMessage) {
	for _, conn := range conns {
		if err := SendJSON(conn, msg); err != nil {
			logx.Errorf("Broadcast message failed: %v", err)
		}
	}
}
