package logic

import (
	"context"

	"github.com/generalzy/zeropan/user/internal/svc"
	"github.com/generalzy/zeropan/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserCallbackLogic {
	return &CreateUserCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserCallbackLogic) CreateUserCallback(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	logx.Info("createUserCallback...............................")
	return &user.CreateUserResponse{}, nil
}
