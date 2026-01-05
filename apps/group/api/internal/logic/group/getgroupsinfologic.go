// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package group

import (
	"context"

	"easy-chat/apps/group/api/internal/svc"
	"easy-chat/apps/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupsInfoLogic {
	return &GetGroupsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupsInfoLogic) GetGroupsInfo(req *types.GetGroupsInfoReq) (resp *types.GetGroupsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
