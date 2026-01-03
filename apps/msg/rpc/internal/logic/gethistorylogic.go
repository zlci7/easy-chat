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
	// todo: add your logic here and delete this line
	//1. 调取model查询数据库
	// list, err := l.svcCtx.MsgModel.FindListBySession(l.ctx, in.UserId, in.PeerId, 1, in.Page, in.PageSize)
	list, err := l.svcCtx.MsgModel.FindListBySession(l.ctx, in.UserId, in.PeerId, in.Type, in.Page, in.PageSize)
	if err != nil && err != models.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	//2. 转换数据格式（model -> proto）

	//3. 返回结果

	return &msg.GetHistoryResp{}, nil
}
