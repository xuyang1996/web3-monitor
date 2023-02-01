package task

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"

	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/util"
)

type DeleteTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTaskLogic) DeleteTask(req *types.DeleteTaskRequest) (resp *types.CommonResponse, err error) {
	tasks, err := l.svcCtx.UserTaskModel.FindByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}

	taskIdMap := make(map[int64]bool)
	for _, task := range tasks {
		taskIdMap[task.Id] = true
	}
	for _, id := range req.IdArray {
		if _, ok := taskIdMap[id]; !ok {
			return nil, util.NewErrorResponseByCode(util.InvalidIdArrayForThisAccount)
		}
	}

	err = l.svcCtx.UserTaskModel.BatchUpdateStatusById(l.ctx, util.UserTaskInactiveStatus, req.IdArray)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}
	return util.NewCommonResponseByCode(util.Success), nil
}
