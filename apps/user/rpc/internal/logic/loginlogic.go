package logic

import (
	"context"
	"errors"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 	1、查找：根据手机号找用户。
	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Phone)
	//用户不存在
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 2、比对：用 bcrypt 比对输入密码和数据库密文。
	if !encrypt.ValidatePassword(in.Password, userInfo.Password) {
		return nil, errors.New("密码错误")
	}

	// 3、返回：返回用户的 ID 和基本信息。
	return &user.LoginResp{
		Id:       userInfo.Id,
		Nickname: userInfo.Nickname,
		Avatar:   userInfo.Avatar,
	}, nil
}
