package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberIdsLogic {
	return &GetGroupMemberIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberIdsLogic) GetGroupMemberIds(in *group.GetGroupMemberIdsReq) (*group.GetGroupMemberIdsResp, error) {
	// todo: add your logic here and delete this line

	return &group.GetGroupMemberIdsResp{}, nil
}
