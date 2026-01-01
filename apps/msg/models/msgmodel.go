package models

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MsgModel = (*customMsgModel)(nil)

type (
	// MsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMsgModel.
	MsgModel interface {
		msgModel
		FindListBySession(ctx context.Context, uid1, uid2 int64, page, pageSize int64) ([]*Msg, error)
	}

	customMsgModel struct {
		*defaultMsgModel
	}
)

// NewMsgModel returns a model for the database table.
func NewMsgModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MsgModel {
	return &customMsgModel{
		defaultMsgModel: newMsgModel(conn, c, opts...),
	}
}

// 查询单聊历史消息de1
func (m *customMsgModel) FindListBySession(ctx context.Context, uid1, uid2 int64, page, pageSize int64) ([]*Msg, error) {
	// 查询条件：(from_uid=A and to_uid=B) or (from_uid=B and to_uid=A)
	// 并且 type=1（单聊），group_id=0
	query := fmt.Sprintf(
		"select %s from %s where type = 1 and ((from_uid = ? and to_uid = ?) or (from_uid = ? and to_uid = ?)) order by seq desc limit ?, ?",
		msgRows, m.table,
	)

	var resp []*Msg
	offset := (page - 1) * pageSize

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid1, uid2, uid2, uid1, offset, pageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
