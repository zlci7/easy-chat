// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"easy-chat/apps/msg/api/internal/logic"
	"easy-chat/apps/msg/api/internal/svc"
	"easy-chat/apps/msg/api/internal/types"
	"easy-chat/pkg/resultx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetHistory(&req)
		resultx.HttpResult(r, w, resp, err)
	}
}
