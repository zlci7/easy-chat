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
		FindUserIdsByGroupId(ctx context.Context, groupId int64, status int32) ([]int64, error)
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

// FindUserIdsByGroupId 根据群ID和状态查询用户ID列表
// status: 0-所有, 1-仅正常成员, 3-排除禁言成员
func (m *customGroupMembersModel) FindUserIdsByGroupId(ctx context.Context, groupId int64, status int32) ([]int64, error) {
	var resp []int64
	var query string
	var args []interface{}

	// 根据 status 构建不同的查询
	switch status {
	case 0:
		// 查询所有成员（不常用）
		query = fmt.Sprintf("select user_id from %s where group_id = ?", m.table)
		args = []interface{}{groupId}

	case 1:
		// 仅查询正常成员（status = 1）
		query = fmt.Sprintf("select user_id from %s where group_id = ? and status = 1", m.table)
		args = []interface{}{groupId}

	case 3:
		// 排除禁言成员（status != 3）
		query = fmt.Sprintf("select user_id from %s where group_id = ? and status in (1, 2)", m.table)
		args = []interface{}{groupId}

	default:
		// 默认只返回正常成员
		query = fmt.Sprintf("select user_id from %s where group_id = ? and status = 1", m.table)
		args = []interface{}{groupId}
	}

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		// 没有成员时返回空切片，而不是错误
		return []int64{}, nil
	default:
		return nil, err
	}
}
