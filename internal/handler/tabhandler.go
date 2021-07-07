package handler

import (
	"net/http"

	"e_tab_be/internal/logic"
	"e_tab_be/internal/svc"
	"e_tab_be/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func TabHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTabLogic(r.Context(), ctx)
		resp, err := l.SaveTabs(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
