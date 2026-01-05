// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package group

import (
	"net/http"

	"easy-chat/apps/group/api/internal/logic/group"
	"easy-chat/apps/group/api/internal/svc"
	"easy-chat/apps/group/api/internal/types"
	"easy-chat/pkg/resultx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JoinGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JoinGroupReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewJoinGroupLogic(r.Context(), svcCtx)
		resp, err := l.JoinGroup(&req)
		resultx.HttpResult(r, w, resp, err)
	}
}
