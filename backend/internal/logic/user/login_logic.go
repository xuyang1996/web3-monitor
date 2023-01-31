package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"backend/internal/model"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/util"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.CommonResponse, err error) {
	user, err := l.svcCtx.UserModel.FindOneByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}

	// no user, insert user
	if user == nil {
		newUser := &model.User{
			Account:   req.Account,
			Signature: req.Signature,
			Email:     req.Email,
			Status:    util.UserActiveStatus,
		}
		_, err := l.svcCtx.UserModel.Insert(l.ctx, newUser)
		if err != nil {
			return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
		}
		return util.NewCommonResponseByCode(util.Success), nil
	}

	// user exists, update user
	user.Signature = req.Signature
	user.Email = req.Email
	user.Status = util.UserActiveStatus
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, util.NewErrorResponseByCode(util.FailToAccessDB)
	}
	return util.NewCommonResponseByCode(util.Success), nil
}
