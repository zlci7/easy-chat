package svc

import (
	"easy-chat/apps/msg/models"
	"easy-chat/apps/msg/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	//声明Model接口
	MsgModel models.MsgModel

	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.DataSource)
	rds := redis.MustNewRedis(c.Redis)

	return &ServiceContext{
		Config: c,

		//实例化Model并且赋值，通过MsgModel变量，可以调用MsgModel接口中的方法。
		//第一个参数是数据库连接，第二个是Redis缓存配置
		MsgModel: models.NewMsgModel(conn, c.Cache),

		//实例化Redis客户端
		RedisClient: rds,
	}
}
