// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"easy-chat/apps/gateway/internal/config"
	"easy-chat/apps/gateway/internal/server"
	"easy-chat/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// 注册连接管理器
	ConnMgr *server.ConnectionManager

	// RPC 客户端（用于服务发现和调用）
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	connmgr := server.NewConnectionManager()

	subscriber := server.NewSubscriber(c.Redis, connmgr)
	subscriber.Start()

	return &ServiceContext{
		Config: c,

		//初始化连接管理器
		ConnMgr: connmgr,

		// 初始化 UserRpc 客户端（这里会从 Etcd 发现服务）
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
