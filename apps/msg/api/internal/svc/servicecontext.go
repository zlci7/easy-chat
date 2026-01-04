// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"easy-chat/apps/msg/api/internal/config"
	"easy-chat/apps/msg/rpc/msgclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	MsgRpc msgclient.Msg
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// 这里的 c.MsgRpc 就是从 YAML -> Config 读进来的配置
		// zrpc.MustNewClient 会根据配置自动连接 Etcd 并发现服务
		MsgRpc: msgclient.NewMsg(zrpc.MustNewClient(c.MsgRpc)),
	}
}
