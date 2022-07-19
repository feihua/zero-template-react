package system

import (
	"net/http"

	"zero-template-react/api/internal/logic/system"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserStatusUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserStatusUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := system.NewUserStatusUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserStatusUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
