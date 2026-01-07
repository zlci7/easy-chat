//心跳机制实现

package server

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

// 心跳配置
const (
	PingInterval = (PongWait * 9) / 10 //心跳间隔
	PongWait     = 600 * time.Second   //pong等待时间
)

func StartHeartbeat(uid int64, userConn *UserConnection, onClose func()) {
	//定时发送Ping
	go func() {
		ticker := time.NewTicker(PingInterval)
		defer ticker.Stop()

		for {
			// time.Sleep(PingInterval)
			// 使用 range ticker.C 或 select
			// 这样更优雅，且 select 方便未来扩展退出信号
			<-ticker.C

			//发送ping消息（使用带锁的Write方法，避免并发写入）
			err := userConn.Write(websocket.PingMessage, nil)
			if err != nil {
				logx.Errorf("send ping message to user %d error: %v", uid, err)
				if onClose != nil {
					onClose()
				}
				return
			}
			logx.Infof("send ping message to user %d success", uid)
		}
	}()
}

// SetupPongHandler 设置 Pong 处理器
func SetupPongHandler(uid int64, conn *websocket.Conn) {
	// 1. 初始化死刑时间
	// 如果接下来 60s 内啥都没发生，ReadMessage 就会报错，连接断开
	conn.SetReadDeadline(time.Now().Add(PongWait))

	// 2. 设置“缓刑”回调
	conn.SetPongHandler(func(string) error {
		logx.Infof("User %d pong received", uid)

		// 关键动作：重置读取截止时间
		// 也就是：只要你理我了，我就再多等你 60s
		conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})
}
