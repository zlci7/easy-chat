// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"easy-chat/apps/msg/api/internal/svc"
	"easy-chat/apps/msg/api/internal/types"
	"easy-chat/apps/msg/rpc/msg"
	"easy-chat/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryLogic {
	return &GetHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryLogic) GetHistory(req *types.GetHistoryReq) (resp *types.GetHistoryResp, err error) {
	// 1. 从 JWT Context 中获取用户 ID
	uid := ctxdata.GetUidFromCtx(l.ctx)

	// 2. 参数校验和默认值处理
	limit := req.Limit
	if limit <= 0 {
		limit = 20 // 默认 20 条
	}
	if limit > 100 {
		limit = 100 // 最大限制 100 条
	}

	// 3. 调用 Msg RPC 服务获取消息（单次请求，不循环）
	rpcResp, err := l.svcCtx.MsgRpc.GetHistory(l.ctx, &msg.GetHistoryReq{
		UserId:    uid,
		PeerId:    req.PeerId,
		Type:      req.Type,
		AnchorSeq: req.AnchorSeq,
		Direction: req.Direction,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	// 4. 转换格式（RPC -> API）
	var list []types.MsgData
	for _, item := range rpcResp.List {
		list = append(list, types.MsgData{
			MsgId:      item.MsgId,
			FromUserId: item.FromUserId,
			ToUserId:   item.ToUserId,
			GroupId:    item.GroupId, // 添加 GroupId 字段
			Type:       item.Type,
			Content:    item.Content,
			Timestamp:  item.Timestamp,
			Seq:        item.Seq,
		})
	}

	// 5. 返回完整响应（包括 HasMore 和 NextSeq，让客户端决定是否继续请求）
	return &types.GetHistoryResp{
		List:    list,
		HasMore: rpcResp.HasMore, // 是否还有更多消息
		NextSeq: rpcResp.NextSeq, // 下次请求的 anchor_seq
	}, nil
}
