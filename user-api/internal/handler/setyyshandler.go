package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-play-1/user-api/internal/logic"
	"go-zero-play-1/user-api/internal/svc"
	"go-zero-play-1/user-api/internal/types"
)

func setYysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.YysReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSetYysLogic(r.Context(), svcCtx)
		resp, err := l.SetYys(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
