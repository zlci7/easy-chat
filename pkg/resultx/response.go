package resultx

import (
	"net/http"
	"strconv"
	"strings"

	"easy-chat/pkg/xerr" // 假设你有自定义错误码包

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type ResponseBean struct {
	Code uint32      `json:"code"` // 使用 int32 更明确
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 失败返回
		var errCode uint32 = xerr.SERVER_COMMON_ERROR // 明确声明为 uint32
		errMsg := "服务器开小差了，稍后再试"

		// 优先级1：检查是否为本地 CodeError
		if e, ok := err.(*xerr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else if s, ok := status.FromError(err); ok {
			// 优先级2：gRPC 错误，尝试从消息中解析自定义错误码
			errMsg = s.Message()
			// 消息格式: "ErrCode:200003，ErrMsg:密码错误"
			if strings.Contains(errMsg, "ErrCode:") && strings.Contains(errMsg, "ErrMsg:") {
				parts := strings.Split(errMsg, "，")
				if len(parts) >= 2 {
					codeStr := strings.TrimPrefix(parts[0], "ErrCode:")
					if code, parseErr := strconv.ParseUint(codeStr, 10, 32); parseErr == nil {
						errCode = uint32(code)
						errMsg = strings.TrimPrefix(parts[1], "ErrMsg:")
					}
				}
			}
			// 如果解析失败，errCode 和 errMsg 保持默认值
		} else if causeErr := errors.Cause(err); causeErr != err {
			// 优先级3：尝试用 pkg/errors 解包
			if e, ok := causeErr.(*xerr.CodeError); ok {
				errCode = e.GetErrCode()
				errMsg = e.GetErrMsg()
			}
		}

		httpx.WriteJson(w, http.StatusOK, ResponseBean{
			Code: errCode,
			Msg:  errMsg,
			Data: nil,
		})
	}
}

func Success(data interface{}) *ResponseBean {
	return &ResponseBean{
		Code: uint32(xerr.OK),
		Msg:  "OK",
		Data: data,
	}
}
