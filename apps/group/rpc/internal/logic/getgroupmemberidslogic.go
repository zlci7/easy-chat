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
	// 1. 调用我们在 Model 层刚写好的查询方法
	uids, err := l.svcCtx.GroupMembersModel.FindUserIdsByGroupId(l.ctx, in.GroupId, in.Status)
	if err != nil {
		l.Logger.Errorf("GetGroupMemberIds db error: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 3. 确保返回空切片而不是 nil（已在 Model 层处理）
	if uids == nil {
		uids = make([]int64, 0)
	}

	return &group.GetGroupMemberIdsResp{
		UserIds: uids,
	}, nil
}
