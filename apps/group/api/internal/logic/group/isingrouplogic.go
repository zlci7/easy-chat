// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package group

import (
	"context"

	"easy-chat/apps/group/api/internal/svc"
	"easy-chat/apps/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsInGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsInGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsInGroupLogic {
	return &IsInGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsInGroupLogic) IsInGroup(req *types.IsInGroupReq) (resp *types.IsInGroupResp, err error) {
	// todo: add your logic here and delete this line

	return
}
