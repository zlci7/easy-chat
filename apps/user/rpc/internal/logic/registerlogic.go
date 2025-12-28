package logic

import (
	"context"

	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/encrypt"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, xerr.NewErrCode(xerr.USER_ALREADY_EXISTS)
	}
	//如果err不是用户不存在，则返回错误
	if err != nil && err != models.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 2. 加密：把明文密码变成哈希值。
	hashedPassword, err := encrypt.EncryptPassword(in.Password)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.USER_ENCRYPT_ERROR)
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
		return nil, xerr.NewErrCode(xerr.USER_SAVE_ERROR)
	}
	//获取用户新id
	userId, err := res.LastInsertId()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.USER_ID_GET_ERROR)
	}
	//返回id，不签发token，api去做
	return &user.RegisterResp{
		UserId: userId,
	}, nil

}
