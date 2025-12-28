package resultx

import (
	"net/http"

	"easy-chat/pkg/xerr" // 假设你有自定义错误码包

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
		errCode := uint32(xerr.SERVER_COMMON_ERROR) // 默认为服务器通用错误
		errMsg := "服务器开小差了，稍后再试"

		// 判断是否为自定义错误 (如果是 grpc 返回的错误，这里也能解析)
		// 这里的逻辑取决于你怎么定义 xerr，通常需要解析 err 类型
		// 简单示例：
		if s, ok := status.FromError(err); ok {
			// 拿到 gRPC 的错误码和错误信息
			errCode = uint32(s.Code())
			errMsg = s.Message()
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
		Code: 0,
		Msg:  "OK",
		Data: data,
	}
}
