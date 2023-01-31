package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"backend/internal/logic/user"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/util"
)

func ChangeEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeEmailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, util.FieldNotSetErrorResponse(err))
			return
		}

		l := user.NewChangeEmailLogic(r.Context(), svcCtx)
		resp, err := l.ChangeEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
