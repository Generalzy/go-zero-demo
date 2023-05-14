package logic

import (
	"context"

	"github.com/generalzy/zeropan/score/internal/svc"
	"github.com/generalzy/zeropan/score/types/score"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreCallbackLogic {
	return &SaveUserScoreCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreCallbackLogic) SaveUserScoreCallback(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	logx.Info("服务回滚,ScoreCallback")
	return &score.UserScoreResponse{Id: in.Id, Score: in.Score}, nil
}
