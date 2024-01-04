package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gox/zero/apiDemo/internal/logic"
	"gox/zero/apiDemo/internal/svc"
	"gox/zero/apiDemo/internal/types"
)

func ApiDemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiDemoLogic(r.Context(), svcCtx)
		resp, err := l.ApiDemo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
