package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupsInfoLogic {
	return &GetGroupsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupsInfoLogic) GetGroupsInfo(in *group.GetGroupsInfoReq) (*group.GetGroupsInfoResp, error) {
	// todo: add your logic here and delete this line

	return &group.GetGroupsInfoResp{}, nil
}
