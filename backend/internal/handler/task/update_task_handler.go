package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"backend/internal/logic/task"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/util"
)

func UpdateTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTaskRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, util.FieldNotSetErrorResponse(err))
			return
		}

		l := task.NewUpdateTaskLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTask(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
