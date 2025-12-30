// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// JWT 配置
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}

	// userRPC 客户端配置（用于服务发现）
	UserRpc zrpc.RpcClientConf

	//msgRPC 客户端配置（用于服务发现）
	MsgRpc zrpc.RpcClientConf

	// Redis 配置
	Redis redis.RedisConf
}
