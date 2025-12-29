// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"time"

	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"easy-chat/apps/user/rpc/userclient"
	"easy-chat/pkg/jwtx"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	//1、调用rpc用户登录服务
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	//2、用户验证正常，生成token
	token, err := jwtx.GetToken(l.svcCtx.Config.Jwt.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Jwt.AccessExpire, rpcResp.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("生成token失败")
	}
	//3、返回信息
	return &types.LoginResp{
		Token:  token,
		Expire: time.Now().Unix() + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
