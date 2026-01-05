package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	//Mysql配置
	DataSource string
	//Mysql连接Redis缓存配置
	Cache cache.CacheConf

	SnowflakeNodeId int64 `json:"SnowflakeNodeId" default:"1"`
}
