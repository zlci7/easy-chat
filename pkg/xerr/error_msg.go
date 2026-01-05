package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差了,请稍后再试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"

	// --- 用户模块错误 200xxx ---
	message[USER_ALREADY_EXISTS] = "用户已注册"
	message[USER_NOT_FOUND] = "用户不存在"
	message[USER_PASSWORD_ERROR] = "密码错误"
	message[USER_ENCRYPT_ERROR] = "密码加密失败"
	message[USER_SAVE_ERROR] = "用户保存失败"
	message[USER_ID_GET_ERROR] = "用户ID获取失败"

	//消息传输
	message[MSG_SAVE_ERROR] = "消息保存失败"
	message[SEQ_GET_ERROR] = "序列号获取失败"

	// 群组模块错误码 (400xxx)
	message[CREATE_GROUP_ERROR] = "创建群失败"
	message[GEN_GROUP_ID_ERROR] = "生成群id失败"
	message[GROUP_NOT_FOUND] = "群不存在"
	message[GROUP_ALREADY_JOINED] = "您已加入该群"
	message[GROUP_DISBANDED] = "该群已解散"
	message[GROUP_BANNED] = "该群已被封禁"
	message[GROUP_MEMBER_LIMIT] = "群成员已满"
	message[GROUP_UPDATE_MEMBER_COUNT_ERROR] = "更新群成员数失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差了,请稍后再试"
	}
}

// 判断是否为自定义错误
func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
