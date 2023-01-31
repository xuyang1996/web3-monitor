package user

import (
	"backend/internal/model"
	"backend/internal/util"
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	user := &model.User{
		Account:   req.Account,
		Signature: req.Signature,
		Status:    util.UserActiveStatus,
	}
	ret, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	logx.Infof("ret: %+v\n", ret)
	return util.NewCommonResponseByCode(util.Success), nil
}
