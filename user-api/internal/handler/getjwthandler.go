package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-play-1/user-api/internal/logic"
	"go-zero-play-1/user-api/internal/svc"
)

func getJWTHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetJWTLogic(r.Context(), svcCtx)
		resp, err := l.GetJWT()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
