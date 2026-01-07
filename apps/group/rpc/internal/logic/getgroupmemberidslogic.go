package logic

import (
	"context"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"
	"easy-chat/pkg/xerr"

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
	// 1. 参数验证
	if in.GroupId == 0 {
		return nil, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}

	// 2. 查询群成员列表
	members, err := l.svcCtx.GroupMembersModel.FindUserIdsByGroupId(l.ctx, in.GroupId, in.Status)
	if err != nil {
		l.Logger.Errorf("find group members error: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	return &group.GetGroupMemberIdsResp{
		UserIds: members,
	}, nil
}
