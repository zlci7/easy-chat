package logic

import (
	"context"
	"time"

	"easy-chat/apps/group/models"
	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *group.JoinGroupReq) (*group.JoinGroupResp, error) {
	// 1. 检查群是否存在
	groupInfo, err := l.svcCtx.GroupModel.FindOneByGroupId(l.ctx, uint64(in.GroupId))
	if err != nil {
		if err == models.ErrNotFound {
			return nil, xerr.NewErrCode(xerr.GROUP_NOT_FOUND)
		}
		l.Logger.Errorf("find group error: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 2. 检查群状态
	if groupInfo.Status == 0 {
		return nil, xerr.NewErrCode(xerr.GROUP_DISBANDED)
	}
	if groupInfo.Status == 2 {
		return nil, xerr.NewErrCode(xerr.GROUP_BANNED)
	}

	// 3. 检查群是否已满
	if groupInfo.MemberCount >= groupInfo.MaxMembers {
		return nil, xerr.NewErrCode(xerr.GROUP_MEMBER_LIMIT)
	}

	// 4. 检查是否已加入过
	existMember, err := l.svcCtx.GroupMembersModel.FindOneByGroupIdUserId(
		l.ctx,
		uint64(in.GroupId),
		in.UserId,
	)
	if err != nil && err != models.ErrNotFound {
		l.Logger.Errorf("check member exist error: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 如果已存在且状态正常，返回已加入
	if existMember != nil && existMember.Status == 1 {
		return nil, xerr.NewErrCode(xerr.GROUP_ALREADY_JOINED)
	}

	now := time.Now().Unix()
	isRejoin := existMember != nil // 判断是否是重新加入

	// 5. 准备成员数据
	member := &models.GroupMembers{
		GroupId:    uint64(in.GroupId),
		UserId:     in.UserId,
		Role:       1, // 1-普通成员
		Status:     1, // 1-正常
		InviterUid: in.InviterUid,
		JoinTime:   now,
		UpdateTime: now,
	}

	// 🔥 6. 事务：添加成员 + 更新成员数（原子操作）
	err = l.svcCtx.GroupModel.AddMemberWithIncrCount(l.ctx, member, isRejoin)
	if err != nil {
		l.Logger.Errorf("add member with incr count error: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	return &group.JoinGroupResp{
		NeedApproval: false,
	}, nil
}
