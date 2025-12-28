package logic

import (
	"context"
	"errors"

	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	//1. 查重：手机号是不是注册过了。
	// FindOneByMobile 是 goctl 生成的 Model 方法
	u, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Phone)
	//用户已注册
	if err == nil && u != nil {
		return nil, status.Error(codes.AlreadyExists, "用户已注册")
	}
	//如果err不是用户不存在，则返回错误
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	// 2. 加密：把明文密码变成哈希值。
	hashedPassword, err := encrypt.EncryptPassword(in.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "密码加密失败")
	}
	// 3. 入库：保存到 MySQL。
	newUser := &models.User{
		Mobile:   in.Phone,
		Password: string(hashedPassword),
		Nickname: in.Nickname,
		Avatar:   "", // 默认头像或随机
		Gender:   0,  // 默认性别
	}

	//将User对象写入mysql
	res, err := l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, errors.New("用户保存失败")
	}
	//获取用户新id
	userId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.New("用户id获取失败")
	}
	//返回id，不签发token，api去做
	return &user.RegisterResp{
		UserId: userId,
	}, nil

}
