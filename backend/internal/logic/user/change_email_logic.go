package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/util"
)

type ChangeEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEmailLogic {
	return &ChangeEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEmailLogic) ChangeEmail(req *types.ChangeEmailRequest) (resp *types.CommonResponse, err error) {
	user, err := l.svcCtx.UserModel.FindOneByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}

	if user == nil {
		return util.NewCommonResponseByCode(util.FailToGetAccount), nil
	}

	user.Email = req.Email
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}
	return util.NewCommonResponseByCode(util.Success), nil
}
