package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsInGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsInGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsInGroupLogic {
	return &IsInGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 辅助功能（供其他RPC服务调用）
func (l *IsInGroupLogic) IsInGroup(in *group.IsInGroupReq) (*group.IsInGroupResp, error) {
	// todo: add your logic here and delete this line

	return &group.IsInGroupResp{}, nil
}
