package server

/*
	连接管理器，管理用户连接
	主要功能：
	1. 添加用户连接
	2. 移除用户连接
	3. 获取用户连接
	4. 向用户发送消息
	5. 向所有用户发送消息
	6. 向指定用户发送消息
	7. 向指定用户发送消息
	8. 向指定用户发送消息
*/

import (
	"sync"

	"github.com/gorilla/websocket"
)

// 用户连接
type UserConnection struct {
	*websocket.Conn
	wLock sync.Mutex //加互斥锁，避免并发写入 WebSocket 导致的 Panic
}

// 安全写入协程安全的websocket连接
func (c *UserConnection) Write(messageType int, data []byte) error {
	c.wLock.Lock()
	defer c.wLock.Unlock()
	return c.Conn.WriteMessage(messageType, data)
}

// 连接管理器
type ConnectionManager struct {
	rwLock      sync.RWMutex              //读写锁
	connections map[int64]*UserConnection //用户id与连接的映射,map的key是用户id,value是连接
}

// 创建连接管理器
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[int64]*UserConnection),
	}
}

// 添加用户连接
func (m *ConnectionManager) AddUserConnection(uid int64, conn *websocket.Conn) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.connections[uid] = &UserConnection{
		Conn: conn,
	}
}

// 移除连接
func (m *ConnectionManager) RemoveUserConnection(uid int64) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//这里简单移除，后续扩展移除特定Session
	delete(m.connections, uid)
}

// 获取连接，返回一个conn指针对象
func (m *ConnectionManager) GetUserConnection(uid int64) *UserConnection {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	return m.connections[uid]
}

// 向指定用户发送消息
func (m *ConnectionManager) SendMsg(uid int64, data []byte) error {
	conn := m.GetUserConnection(uid)
	if conn == nil {
		//用户不在线，不做处理，由RPC服务中的离线消息处理逻辑
		return nil
	}
	return conn.Write(websocket.TextMessage, data)
}
