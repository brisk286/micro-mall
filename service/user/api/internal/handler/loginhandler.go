package handler

import (
	"micro-mall-server/service/user/api/internal/logic"
	"micro-mall-server/service/user/api/internal/svc"
	"micro-mall-server/service/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest

		// 解析Request到req
		if err := httpx.Parse(r, &req); err != nil {
			// 将err写入ResponseWriter返回
			httpx.Error(w, err)
			return
		}

		// 登录并返回response
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
