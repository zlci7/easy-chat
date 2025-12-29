// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"easy-chat/apps/gateway/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws() error {
	// todo: add your logic here and delete this line

	return nil
}
