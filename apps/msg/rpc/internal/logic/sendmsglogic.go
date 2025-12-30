package logic

import (
	"context"
	"encoding/json"
	"time"

	"easy-chat/apps/msg/models"
	"easy-chat/apps/msg/rpc/internal/svc"
	"easy-chat/apps/msg/rpc/msg"
	"easy-chat/pkg/model/mq"
	"easy-chat/pkg/xerr"

	"github.com/google/uuid"
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

	//1. 生成基础信息
	msgId := uuid.New().String()
	now := time.Now().UnixMilli()

	//2.存入mysql
	newMsg := &models.Msg{
		MsgId:      msgId,
		FromUid:    in.FromUserId,
		ToUid:      in.ToUserId,
		Type:       int64(in.Type), //类型转换为int64
		Content:    in.Content,
		CreateTime: now,
	}
	//2.1 插入mysql
	_, err := l.svcCtx.MsgModel.Insert(l.ctx, newMsg)
	if err != nil {
		l.Logger.Errorf("insert msg to mysql error: %v", err)
		return nil, xerr.NewErrCode(xerr.MSG_SAVE_ERROR)
	}

	//3.构造广播消息
	broadcastMsg := &mq.BroadcastMsg{
		MsgId:      msgId,
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Type:       int(in.Type),
		Content:    in.Content,
		Timestamp:  now,
	}

	//4. 序列化广播消息
	broadcastMsgBytes, err := json.Marshal(broadcastMsg)
	if err != nil {
		l.Logger.Errorf("marshal broadcast msg to json error: %v", err)
		//如果序列化失败，则返回消息保存失败
		return nil, xerr.NewErrCode(xerr.MSG_SAVE_ERROR)
	}

	//5. 将广播消息推送到Redis
	_, err = l.svcCtx.RedisClient.Publish(mq.PushMsgKey, string(broadcastMsgBytes))
	if err != nil {
		l.Logger.Errorf("publish broadcast msg to redis error: %v", err)
		//如果存库成功但推送失败，属于系统异常，此处可以重试或者记录错误
		//但是不会影响RPC返回成功，因为客户端可以通过拉去历史消息补救
	}

	//5.返回给rpc服务
	return &msg.SendMsgResp{
		MsgId:     msgId,
		Timestamp: now,
	}, nil
}
