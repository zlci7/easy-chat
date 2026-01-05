package models

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupsModel = (*customGroupsModel)(nil)

type (
	// GroupsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupsModel.
	GroupsModel interface {
		groupsModel
		// 新增：事务创建群组和群主
		InsertWithGroupAndMember(ctx context.Context, group *Groups, member *GroupMembers) error
	}

	customGroupsModel struct {
		*defaultGroupsModel
	}
)

// NewGroupsModel returns a model for the database table.
func NewGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupsModel {
	return &customGroupsModel{
		defaultGroupsModel: newGroupsModel(conn, c, opts...),
	}
}

// 事务创建群组和群主
func (m *customGroupsModel) InsertWithGroupAndMember(ctx context.Context, group *Groups, member *GroupMembers) error {
	// 开启事务
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 1. 插入群组
		queryGroup := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, groupsRowsExpectAutoSet)
		// 注意：这里字段顺序要和你 sql 生成的 groupsRowsExpectAutoSet 对应
		// 假设顺序是: group_id, name, owner_uid, type, avatar, create_time
		_, err := session.ExecCtx(ctx, queryGroup, group.GroupId, group.Name, group.OwnerUid, group.Type, group.Avatar, group.CreateTime)
		if err != nil {
			return err
		}

		// 2. 插入群主作为成员
		// 注意：这里硬编码了 group_members 表名，实际项目中最好通过 config 或常量获取
		queryMember := "insert into group_members (group_id, user_id, role, join_time) values (?, ?, ?, ?)"
		_, err = session.ExecCtx(ctx, queryMember, member.GroupId, member.UserId, member.Role, member.JoinTime)
		return err
	})
}
