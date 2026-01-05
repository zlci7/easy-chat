// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package group

import (
	"context"

	"easy-chat/apps/group/api/internal/svc"
	"easy-chat/apps/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MuteMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMuteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MuteMemberLogic {
	return &MuteMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MuteMemberLogic) MuteMember(req *types.MuteMemberReq) (resp *types.MuteMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
