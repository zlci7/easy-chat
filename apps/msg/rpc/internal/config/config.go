package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	//Mysql配置
	DataSource string
	//Mysql连接Redis缓存配置
	Cache cache.CacheConf

	//Redis连接配置
	Redis redis.RedisConf
}
