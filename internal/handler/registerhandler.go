// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"user_api/common/response"
	"user_api/internal/logic"
	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Output(w, response.RESP_BAD_REQUEST, nil, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			response.Output(w, response.RESP_BAD_GATEWAY, err.Error(), nil)
		} else {
			response.Output(w, response.RESP_SUCC, nil, resp)
		}
	}
}
