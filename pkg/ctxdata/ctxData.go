package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId 必须和 jwtx.GetToken 里的 key 保持一致
var CtxKeyJwtUserId = "uid"

// GetUidFromCtx 从 context 中获取用户ID
func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64

	// go-zero 解析 JWT 后，会把 claims 里的 json number 放入 context
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx parse error: %+v", err)
		}
	} else {
		// 如果取不到，说明可能是未登录或 token 无效
		logx.WithContext(ctx).Error("GetUidFromCtx: uid not found in context")
	}
	return uid
}
