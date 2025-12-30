package logic

import (
	"context"

	"easy-chat/apps/msg/rpc/internal/svc"
	"easy-chat/apps/msg/rpc/msg"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送消息：处理存储、生成ID、广播
func (l *SendMsgLogic) SendMsg(in *msg.SendMsgReq) (*msg.SendMsgResp, error) {
	// todo: add your logic here and delete this line

	return &msg.SendMsgResp{}, nil
}
