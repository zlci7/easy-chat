/*
鉴权与握手处理
*/

package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"

	"easy-chat/apps/gateway/internal/svc"
	"easy-chat/pkg/jwtx" // 复用阶段一写的 JWT 工具
)

var upgrader = websocket.Upgrader{
	// 允许跨域（测试方便）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 获取 Token
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "No token found", http.StatusUnauthorized)
			return
		}

		// 2. 解析 Token (使用阶段一封装的工具)
		// 注意：ParseToken 需要你自己去 jwtx 包里实现一个解析方法，或者用 jwt 库直接解
		claims, err := jwtx.ParseToken(token, svcCtx.Config.Jwt.AccessSecret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 从 claims 拿到 UserID (假设你的 claim key 是 "uid")
		uid := int64(claims["uid"].(float64))

		// 3. 升级为 WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error("Upgrade failure:", err)
			return
		}

		// 4. 加入连接管理器
		svcCtx.ConnMgr.AddUserConnection(uid, conn)
		logx.Infof("User %d connected", uid)

		// 5. 开启读取循环 (KeepAlive)
		// 必须有一个循环读取消息，否则连接会断开
		go func() {
			defer func() {
				svcCtx.ConnMgr.RemoveUserConnection(uid)
				conn.Close()
			}()

			for {
				// 这里目前只处理 Ping/Pong 或者简单消息
				// 后续上行消息逻辑写在这里
				_, _, err := conn.ReadMessage()
				if err != nil {
					break
				}
			}
		}()
	}
}
