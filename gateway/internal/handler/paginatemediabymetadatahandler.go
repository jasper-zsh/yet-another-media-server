package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yet-another-media-server/gateway/internal/logic"
	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"
)

func paginateMediaByMetadataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PaginateMediaByMetadataRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPaginateMediaByMetadataLogic(r.Context(), svcCtx)
		resp, err := l.PaginateMediaByMetadata(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
