package task

import (
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

func (l *ReadTaskLogic) ReadTask(req *types.ReadTaskRequest) (resp *types.ReadTaskResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
