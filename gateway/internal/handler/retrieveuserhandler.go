package handler

import (
	"net/http"

	"github.com/generalzy/zeropan/gateway/internal/logic"
	"github.com/generalzy/zeropan/gateway/internal/svc"
	"github.com/generalzy/zeropan/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RetrieveUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RetrieveUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRetrieveUserLogic(r.Context(), svcCtx)
		resp, err := l.RetrieveUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
