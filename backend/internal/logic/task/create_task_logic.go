package task

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"backend/internal/model"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/util"
)

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.CreateTaskRequest) (resp *types.CommonResponse, err error) {
	tasks, err := l.svcCtx.UserTaskModel.FindByAccountAndChainAndAddressAndType(l.ctx,
		req.Account, req.Chain, req.Address, req.Type)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}

	// no task, insert task
	if len(tasks) == 0 {
		newTask := &model.UserTask{
			Account: req.Account,
			Chain:   int64(req.Chain),
			Address: req.Address,
			Type:    req.Type,
			Status:  util.UserTaskActiveStatus,
		}
		_, err := l.svcCtx.UserTaskModel.Insert(l.ctx, newTask)
		if err != nil {
			return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
		}
		return util.NewCommonResponseByCode(util.Success), nil
	}

	// task exists, update task
	if len(tasks) == 1 {
		tasks[0].Status = util.UserTaskActiveStatus
		err = l.svcCtx.UserTaskModel.Update(l.ctx, &tasks[0])
		if err != nil {
			return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
		}
		return util.NewCommonResponseByCode(util.Success), nil
	}

	return nil, util.NewErrorResponseByCode(util.DuplicateUserTasks)
}
