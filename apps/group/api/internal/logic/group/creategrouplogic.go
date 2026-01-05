// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package group

import (
	"context"

	"easy-chat/apps/group/api/internal/svc"
	"easy-chat/apps/group/api/internal/types"
	"easy-chat/apps/group/rpc/group"
	"easy-chat/pkg/ctxdata"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	// 1. 从 JWT Context 中获取当前登录用户ID（作为群主），不需要额外校验
	uid := ctxdata.GetUidFromCtx(l.ctx)

	// 2. 参数校验
	if req.Name == "" {
		return nil, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}

	// 3. 调用 Group RPC 服务创建群
	rpcResp, err := l.svcCtx.GroupRpc.CreateGroup(l.ctx, &group.CreateGroupReq{
		OwnerUid:    uid, // 使用JWT中的user_id作为群主
		Name:        req.Name,
		Avatar:      req.Avatar,
		Description: req.Description,
		MemberIds:   req.MemberIds, // 初始成员列表（可选）
	})
	if err != nil {
		l.Logger.Errorf("CreateGroup RPC error: %v", err)
		return nil, err
	}

	// 4. 转换格式（RPC -> API）
	return &types.CreateGroupResp{
		GroupId: rpcResp.GroupId,
		GroupInfo: types.GroupInfo{
			GroupId:     rpcResp.GroupInfo.GroupId,
			Name:        rpcResp.GroupInfo.Name,
			OwnerUid:    rpcResp.GroupInfo.OwnerUid,
			Type:        rpcResp.GroupInfo.Type,
			Avatar:      rpcResp.GroupInfo.Avatar,
			Status:      rpcResp.GroupInfo.Status,
			Description: rpcResp.GroupInfo.Description,
			Notice:      rpcResp.GroupInfo.Notice,
			MaxMembers:  rpcResp.GroupInfo.MaxMembers,
			MemberCount: rpcResp.GroupInfo.MemberCount,
			JoinType:    rpcResp.GroupInfo.JoinType,
			MuteAll:     rpcResp.GroupInfo.MuteAll,
			CreateTime:  rpcResp.GroupInfo.CreateTime,
			UpdateTime:  rpcResp.GroupInfo.UpdateTime,
		},
	}, nil
}
