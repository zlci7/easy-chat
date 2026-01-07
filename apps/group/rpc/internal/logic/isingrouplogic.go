package logic

import (
	"context"

	"easy-chat/apps/group/models"
	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"
	"easy-chat/pkg/xerr"

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
	user, err := l.svcCtx.GroupMembersModel.FindOneByGroupIdUserId(l.ctx, uint64(in.GroupId), in.UserId)
	if err != nil {
		if err == models.ErrNotFound {
			return &group.IsInGroupResp{
				IsMember: false,
			}, nil
		}
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	return &group.IsInGroupResp{
		IsMember: true,
		Role:     int32(user.Role),
	}, nil
}
