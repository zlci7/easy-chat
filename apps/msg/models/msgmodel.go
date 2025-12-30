package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MsgModel = (*customMsgModel)(nil)

type (
	// MsgModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMsgModel.
	MsgModel interface {
		msgModel
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
