package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *group.JoinGroupReq) (*group.JoinGroupResp, error) {
	// todo: add your logic here and delete this line

	return &group.JoinGroupResp{}, nil
}
