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
		// 基于 seq 的双向查询（推荐使用）
		FindListBySeq(ctx context.Context, uid1, uid2 int64, msgType int64, anchorSeq int64, direction int32, limit int64) ([]*Msg, error)
		// 兼容旧版本的分页查询（已废弃）
		FindListBySession(ctx context.Context, uid1, uid2 int64, msgType int64, page, pageSize int64) ([]*Msg, error)
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

// 基于 seq 的双向查询（支持单聊和群聊）
func (m *customMsgModel) FindListBySeq(ctx context.Context, uid1, uid2 int64, msgType int64, anchorSeq int64, direction int32, limit int64) ([]*Msg, error) {
	var query string
	var args []interface{}

	if msgType == 1 {
		// 单聊：查询条件 (from_uid=A and to_uid=B) or (from_uid=B and to_uid=A)
		baseCondition := "type = ? and ((from_uid = ? and to_uid = ?) or (from_uid = ? and to_uid = ?))"

		if anchorSeq == 0 || direction == 0 {
			// 向前查询（查历史）或首次加载：seq < anchor_seq（anchor_seq=0时查最新）
			if anchorSeq == 0 {
				// 首次加载，获取最新消息
				query = fmt.Sprintf(
					"select %s from %s where %s order by seq desc limit ?",
					msgRows, m.table, baseCondition,
				)
				args = []interface{}{msgType, uid1, uid2, uid2, uid1, limit}
			} else {
				// 向前查询历史
				query = fmt.Sprintf(
					"select %s from %s where %s and seq < ? order by seq desc limit ?",
					msgRows, m.table, baseCondition,
				)
				args = []interface{}{msgType, uid1, uid2, uid2, uid1, anchorSeq, limit}
			}
		} else {
			// direction == 1: 向后查询（查新消息）：seq > anchor_seq
			query = fmt.Sprintf(
				"select %s from %s where %s and seq > ? order by seq asc limit ?",
				msgRows, m.table, baseCondition,
			)
			args = []interface{}{msgType, uid1, uid2, uid2, uid1, anchorSeq, limit}
		}
	} else {
		// 群聊：只需要 group_id 匹配（uid2 作为 group_id）
		baseCondition := "type = ? and group_id = ?"

		if anchorSeq == 0 || direction == 0 {
			if anchorSeq == 0 {
				// 首次加载群聊消息
				query = fmt.Sprintf(
					"select %s from %s where %s order by seq desc limit ?",
					msgRows, m.table, baseCondition,
				)
				args = []interface{}{msgType, uid2, limit}
			} else {
				// 向前查询群聊历史
				query = fmt.Sprintf(
					"select %s from %s where %s and seq < ? order by seq desc limit ?",
					msgRows, m.table, baseCondition,
				)
				args = []interface{}{msgType, uid2, anchorSeq, limit}
			}
		} else {
			// 向后查询群聊新消息
			query = fmt.Sprintf(
				"select %s from %s where %s and seq > ? order by seq asc limit ?",
				msgRows, m.table, baseCondition,
			)
			args = []interface{}{msgType, uid2, anchorSeq, limit}
		}
	}

	var resp []*Msg
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 查询单聊历史消息（已废弃，保留兼容性）
func (m *customMsgModel) FindListBySession(ctx context.Context, uid1, uid2 int64, msgType int64, page, pageSize int64) ([]*Msg, error) {
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
