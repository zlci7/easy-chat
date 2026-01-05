package logic

import (
	"context"
	"time"

	"easy-chat/apps/group/models"
	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/svc"
	"easy-chat/pkg/idgen"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 基础功能
func (l *CreateGroupLogic) CreateGroup(in *group.CreateGroupReq) (*group.CreateGroupResp, error) {
	// todo: add your logic here and delete this line
	//1、生成群id(雪花算法)
	groupId := idgen.GenInt64ID()
	if groupId == 0 {
		return nil, xerr.NewErrCode(xerr.GEN_GROUP_ID_ERROR)
	}

	now := time.Now().Unix()

	//2、创建群实体
	newGroup := &models.Groups{
		GroupId:       uint64(groupId),
		Name:          in.Name,
		OwnerUid:      in.OwnerUid,
		Type:          1, // 默认普通群
		Avatar:        in.Avatar,
		Status:        1, // 正常状态
		Description:   in.Description,
		MaxMembers:    500,
		MemberCount:   1, // 群主算一个成员
		JoinType:      1, // 默认自由加入
		InviteConfirm: 0, //默认不邀请确认
		MuteAll:       0, //默认不禁言
		CreateTime:    now,
		UpdateTime:    now,
	}

	//3、创建群成员实体
	newGroupMember := &models.GroupMembers{
		GroupId: uint64(groupId),
		UserId:  in.OwnerUid,
		Role:    3,
		Status:  1, //默认正常
		// Nickname: in.Name,   //这里需要调用user rpc服务去查询nickname
		// InviterUid: 0,       //暂时不实现邀请
		// LastAckMsgId: 0, //暂时不实现已读消息
		// MuteEndTime:  0, //暂时不实现禁言
		JoinTime: time.Now().Unix(),
	}

	//4、调用model方法添加群信息及群主信息
	err := l.svcCtx.GroupModel.InsertWithGroupAndMember(l.ctx, newGroup, newGroupMember)
	if err != nil {
		l.Logger.Error("创建群失败", err)
		return nil, xerr.NewErrCode(xerr.CREATE_GROUP_ERROR)
	}

	// 5、封装返回结果
	return &group.CreateGroupResp{
		GroupId: groupId,
		GroupInfo: &group.GroupInfo{
			GroupId:     int64(groupId),
			Name:        in.Name,
			OwnerUid:    in.OwnerUid,
			Type:        1, // 默认普通群
			Avatar:      in.Avatar,
			Status:      1, // 正常状态
			Description: in.Description,
			MaxMembers:  500,
			MemberCount: 1, // 群主算一个成员
			JoinType:    1, // 默认自由加入
			MuteAll:     0, //默认不禁言
			CreateTime:  now,
			UpdateTime:  now,
		},
	}, nil
}
