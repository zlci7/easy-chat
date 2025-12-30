package server

import (
	"context"
	"easy-chat/pkg/model/mq"
	"encoding/json"

	"github.com/redis/go-redis/v9" // 依然需要这个库做底层支持
	"github.com/zeromicro/go-zero/core/logx"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading" // <--- 关键：引入 go-zero 的线程管理
)

type Subscriber struct {
	rds     *redis.Client
	connMgr *ConnectionManager
}

func NewSubscriber(c zredis.RedisConf, connMgr *ConnectionManager) *Subscriber {
	// 1. 复用 go-zero 的配置 (RedisConf)
	// 这样你就不用在 yaml 里写两遍 redis 配置了
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Pass,
		DB:       1,
		// 生产环境可能还需要配置 PoolSize 等，但在 Sub 模式下连接数只需 1 个
	})

	return &Subscriber{
		rds:     rdb,
		connMgr: connMgr,
	}
}

func (s *Subscriber) Start() {
	// 使用 go-zero 的安全协程启动器
	// 这样写就不需要在外层调用 go s.Start()，直接 s.Start() 即可，它内部会异步执行
	threading.GoSafe(func() {
		s.subscribeLoop()
	})
}

// 把具体的循环逻辑抽离出来，保持代码整洁
func (s *Subscriber) subscribeLoop() {
	logx.Info("Attributes: [Redis Sub] starting...")

	//创建一个上下文
	ctx := context.Background()
	//订阅消息
	pubsub := s.rds.Subscribe(ctx, mq.PushMsgKey)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		// 这里的处理逻辑保持不变
		var broadcastMsg mq.BroadcastMsg
		if err := json.Unmarshal([]byte(msg.Payload), &broadcastMsg); err != nil {
			// 使用 logx 记录结构化日志
			logx.Errorf("[Redis Sub] Unmarshal error: %v, payload: %s", err, msg.Payload)
			continue
		}

		// // 发送逻辑
		// if err := s.connMgr.SendMsg(broadcastMsg.ToUserId, []byte(broadcastMsg.Content)); err != nil {
		// 	// 这种通常是 debug 级别的日志，避免生产环境日志爆炸
		// 	// logx.Debugf("User %d not on this node", broadcastMsg.UserId)
		// }

		// 发送逻辑：把完整的消息结构体发送给客户端
		jsonData, err := json.Marshal(broadcastMsg)
		if err != nil {
			logx.Errorf("[Redis Sub] Marshal error: %v", err)
			continue
		}

		if err := s.connMgr.SendMsg(broadcastMsg.ToUserId, jsonData); err != nil {
			// 用户可能不在当前节点，这是正常的
			// logx.Debugf("User %d not on this node", broadcastMsg.ToUserId)
		}
	}
}
