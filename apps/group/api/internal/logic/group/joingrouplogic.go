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

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup(req *types.JoinGroupReq) (resp *types.JoinGroupResp, err error) {
	// 1. 从 JWT Context 中获取当前登录用户ID，不需要校验
	uid := ctxdata.GetUidFromCtx(l.ctx)

	// 2. 参数校验
	if req.GroupId == 0 {
		return nil, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}

	// 3. 调用 Group RPC 服务加入群
	rpcResp, err := l.svcCtx.GroupRpc.JoinGroup(l.ctx, &group.JoinGroupReq{
		GroupId:    req.GroupId,
		UserId:     uid,            // 使用JWT中的user_id
		InviterUid: req.InviterUid, // 邀请人ID（可选）
		ApplyMsg:   req.ApplyMsg,   // 申请消息（可选）
	})
	if err != nil {
		l.Logger.Errorf("JoinGroup RPC error: %v", err)
		return nil, err
	}

	// 4. 返回结果
	return &types.JoinGroupResp{
		NeedApproval: rpcResp.NeedApproval,
	}, nil
}
