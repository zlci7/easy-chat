// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//添加UserRpc字段
	UserRpc zrpc.RpcClientConf

	//添加JWT鉴权
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}
}
