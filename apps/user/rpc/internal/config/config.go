package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	//告诉代码，配置文件里会有MySQL数据库地址和 Redis 配置。
	DataSource string          //对应mysql连接字符串
	Cache      cache.CacheConf //对应Redis缓存配置
}
