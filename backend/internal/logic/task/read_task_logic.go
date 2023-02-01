package task

import (
	"backend/internal/model"
	"backend/internal/util"
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadTaskLogic {
	return &ReadTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadTaskLogic) ReadTask(req *types.ReadTaskRequest) (*types.ReadTaskResponse, error) {
	resp := new(types.ReadTaskResponse)
	tasks, err := l.svcCtx.UserTaskModel.FindByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}
	resp.Code = util.Success
	resp.Message = util.MessageMap[util.Success]
	resp.Data = GetTaskInfo(tasks)
	return resp, nil
}

func GetTaskInfo(tasks []model.UserTask) []types.TaskInfo {
	newTasks := make([]types.TaskInfo, 0)
	for _, task := range tasks {
		if task.Status == util.UserTaskInactiveStatus {
			continue
		}
		taskInfo := types.TaskInfo{
			Id:      task.Id,
			Account: task.Account,
			Chain:   int32(task.Chain),
			Address: task.Address,
			Type:    task.Type,
			Status:  task.Status,
		}
		newTasks = append(newTasks, taskInfo)
	}
	return newTasks
}
