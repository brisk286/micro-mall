package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"micro-mall-server/app/user/api/internal/logic/user"
	"micro-mall-server/app/user/api/internal/svc"
	"micro-mall-server/app/user/api/internal/types"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
