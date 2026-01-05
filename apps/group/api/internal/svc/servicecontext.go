// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"easy-chat/apps/group/api/internal/config"
	"easy-chat/apps/group/rpc/groupclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	GroupRpc groupclient.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GroupRpc: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRpc)),
	}
}
