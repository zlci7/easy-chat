package svc

import (
	"easy-chat/apps/group/models"
	"easy-chat/apps/group/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	//声明Model接口
	GroupModel        models.GroupsModel
	GroupMembersModel models.GroupMembersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config: c,

		//实例化GroupModel并且赋值，通过GroupModel变量，可以调用GroupModel接口中的方法。
		//第一个参数是数据库连接，第二个是Redis缓存配置
		GroupModel: models.NewGroupsModel(conn, c.Cache),

		//实例化GroupMembersModel并且赋值，通过GroupMembersModel变量，可以调用GroupMembersModel接口中的方法。
		//第一个参数是数据库连接，第二个是Redis缓存配置
		GroupMembersModel: models.NewGroupMembersModel(conn, c.Cache),
	}
}
