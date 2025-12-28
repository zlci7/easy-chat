package svc

import (
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	//声明Model接口
	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	//初始化mysql连接
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config: c,

		//实例化Model并且赋值，通过UserModel变量，可以调用UserModel接口中的方法。
		//第一个参数是数据库连接，第二个是Redis缓存配置
		UserModel: models.NewUserModel(conn, c.Cache),
	}
}
