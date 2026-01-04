package logic

import (
	"context"

	"easy-chat/apps/msg/models"
	"easy-chat/apps/msg/rpc/internal/svc"
	"easy-chat/apps/msg/rpc/msg"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryLogic {
	return &GetHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取消息历史
func (l *GetHistoryLogic) GetHistory(in *msg.GetHistoryReq) (*msg.GetHistoryResp, error) {
	// 参数校验和默认值设置
	limit := in.Limit
	if limit <= 0 {
		limit = 20 // 默认返回20条
	}
	if limit > 100 {
		limit = 100 // 最大不超过100条
	}

	// 1. 调用 model 基于 seq 查询数据库
	list, err := l.svcCtx.MsgModel.FindListBySeq(
		l.ctx,
		in.UserId,
		in.PeerId,
		in.Type,
		in.AnchorSeq,
		in.Direction,
		limit,
	)
	if err != nil && err != models.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 2. 转换数据格式（model -> proto）
	var respList []*msg.MsgData
	for _, data := range list {
		respList = append(respList, &msg.MsgData{
			MsgId:      data.MsgId,
			FromUserId: data.FromUid,
			ToUserId:   data.ToUid,
			GroupId:    data.GroupId,
			Type:       data.Type,
			Content:    data.Content,
			Timestamp:  data.CreateTime,
			Seq:        data.Seq,
		})
	}

	// 3. 计算是否还有更多消息和下次请求的 anchor_seq
	hasMore := len(respList) >= int(limit)
	var nextSeq int64
	if len(respList) > 0 {
		if in.Direction == 0 {
			// 向前查询（查历史），下次从最早的消息继续向前
			nextSeq = respList[len(respList)-1].Seq
		} else {
			// 向后查询（查新消息），下次从最新的消息继续向后
			nextSeq = respList[len(respList)-1].Seq
		}
	}

	// 4. 返回结果
	return &msg.GetHistoryResp{
		List:    respList,
		HasMore: hasMore,
		NextSeq: nextSeq,
	}, nil
}
