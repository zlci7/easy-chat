package models

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupMembersModel = (*customGroupMembersModel)(nil)

type (
	// GroupMembersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupMembersModel.
	GroupMembersModel interface {
		groupMembersModel
		// 新增：根据群ID查询用户ID列表
		FindUserIdsByGroupId(ctx context.Context, groupId int64) ([]int64, error)
	}

	customGroupMembersModel struct {
		*defaultGroupMembersModel
	}
)

// NewGroupMembersModel returns a model for the database table.
func NewGroupMembersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupMembersModel {
	return &customGroupMembersModel{
		defaultGroupMembersModel: newGroupMembersModel(conn, c, opts...),
	}
}

// 根据群ID查询用户ID列表
func (m *customGroupMembersModel) FindUserIdsByGroupId(ctx context.Context, groupId int64) ([]int64, error) {

	var resp []int64
	query := fmt.Sprintf("select user_id from %s where group_id =?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, groupId)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}
