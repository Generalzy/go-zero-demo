package logic

import (
	"context"

	"github.com/generalzy/zeropan/gateway/internal/svc"
	"github.com/generalzy/zeropan/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveUserScoreCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreCallbackLogic {
	return &SaveUserScoreCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveUserScoreCallbackLogic) SaveUserScoreCallback(req *types.UserScoreRequest) (resp *types.UserScoreResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
