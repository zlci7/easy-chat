package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMembersLogic {
	return &GetGroupMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMembersLogic) GetGroupMembers(in *group.GetGroupMembersReq) (*group.GetGroupMembersResp, error) {
	// todo: add your logic here and delete this line

	return &group.GetGroupMembersResp{}, nil
}
