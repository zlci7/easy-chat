package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchIsInGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchIsInGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchIsInGroupLogic {
	return &BatchIsInGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchIsInGroupLogic) BatchIsInGroup(in *group.BatchIsInGroupReq) (*group.BatchIsInGroupResp, error) {
	// todo: add your logic here and delete this line

	return &group.BatchIsInGroupResp{}, nil
}
