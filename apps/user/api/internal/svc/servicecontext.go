// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// 定义接口，方便 Logic 层调用
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// 这里的 c.UserRpc 就是从 YAML -> Config 读进来的配置
		// zrpc.MustNewClient 会根据配置自动连接 Etcd 并发现服务
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
