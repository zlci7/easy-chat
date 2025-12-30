package logic

import (
	"context"

	"easy-chat/apps/msg/rpc/internal/svc"
	"easy-chat/apps/msg/rpc/msg"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryLogic {
	return &GetHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取消息历史
func (l *GetHistoryLogic) GetHistory(in *msg.GetHistoryReq) (*msg.GetHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &msg.GetHistoryResp{}, nil
}
