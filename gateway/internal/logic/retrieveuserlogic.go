package logic

import (
	"context"
	"github.com/generalzy/zeropan/gateway/internal/errorx"
	"github.com/generalzy/zeropan/user/types/user"

	"github.com/generalzy/zeropan/gateway/internal/svc"
	"github.com/generalzy/zeropan/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveUserLogic {
	return &RetrieveUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveUserLogic) RetrieveUser(req *types.RetrieveUserRequest) (resp *types.RetrieveUserResponse, err error) {
	userID := req.ID
	if userID == "" {
		return nil, errorx.ParamsError
	}
	// 获取用户信息
	result, e := l.svcCtx.UserRpc.RetrieveUser(context.TODO(), &user.RetrieveUserRequest{
		Id: userID,
	})
	if e != nil {
		return nil, e
	}

	return &types.RetrieveUserResponse{
		Name:      result.GetName(),
		Password:  "",
		Email:     result.GetEmail(),
		CreatedAt: result.GetCreateAt(),
		UpdatedAt: result.GetUpdateAt(),
		DeletedAt: "",
	}, nil
}
