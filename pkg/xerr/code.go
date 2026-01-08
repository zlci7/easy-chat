package xerr

import "fmt"

// 常用通用错误码
const (
	OK = 200
	//通用错误
	SERVER_COMMON_ERROR = 100001 // 服务器开小差了
	REUQEST_PARAM_ERROR = 100002 // 参数错误
	TOKEN_EXPIRE_ERROR  = 100003 // Token过期
	DB_ERROR            = 100004 // 数据库错误
	TOO_MANY_REQUESTS   = 100005

	// 用户模块错误码 (200xxx)
	USER_ALREADY_EXISTS = 200001 // 用户已注册
	USER_NOT_FOUND      = 200002 // 用户不存在
	USER_PASSWORD_ERROR = 200003 // 密码错误
	USER_ENCRYPT_ERROR  = 200004 // 密码加密失败
	USER_SAVE_ERROR     = 200005 // 用户保存失败
	USER_ID_GET_ERROR   = 200006 // 用户ID获取失败
	USER_NOT_LOGIN      = 200007 // 用户未登录

	//消息传输
	MSG_SAVE_ERROR = 300001 // 消息保存失败
	SEQ_GET_ERROR  = 300002 // 序列号获取失败

	// 群组模块错误码 (400xxx)
	CREATE_GROUP_ERROR              = 400001 // 创建群失败
	GEN_GROUP_ID_ERROR              = 400002 // 生成群id失败
	GROUP_NOT_FOUND                 = 400003 // 群不存在
	GROUP_ALREADY_JOINED            = 400004 // 已加入该群
	GROUP_DISBANDED                 = 400005 // 群已解散
	GROUP_BANNED                    = 400006 // 群已被封禁
	GROUP_MEMBER_LIMIT              = 400007 // 群成员已满
	GROUP_UPDATE_MEMBER_COUNT_ERROR = 400008 // 更新群成员数失败)
)

// CodeError 自定义错误结构体
type CodeError struct {
	errCode uint32
	errMsg  string
}

// 1. 获取错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// 2. 获取错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// 3. 实现 error 接口
func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

// NewErrCode 工厂方法：通过错误码创建错误
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

// NewErrMsg 工厂方法：创建自定义消息的错误
func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
