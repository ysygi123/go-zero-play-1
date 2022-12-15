package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-play-1/user-api/internal/logic"
	"go-zero-play-1/user-api/internal/svc"
	"go-zero-play-1/user-api/internal/types"
)

func calculateInventoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CalculateInventoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCalculateInventoryLogic(r.Context(), svcCtx)
		resp, err := l.CalculateInventory(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
