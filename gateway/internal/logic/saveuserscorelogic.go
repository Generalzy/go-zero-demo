package logic

import (
	"context"
	"github.com/generalzy/zeropan/score/types/score"

	"github.com/generalzy/zeropan/gateway/internal/svc"
	"github.com/generalzy/zeropan/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreLogic {
	return &SaveUserScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveUserScoreLogic) SaveUserScore(req *types.UserScoreRequest) (resp *types.UserScoreResponse, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.ScoreRpc.SaveUserScore(context.TODO(), &score.UserScoreRequest{
		Id:    "xxx",
		Score: 100,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserScoreResponse{
		ID:    response.Id,
		Score: response.Score,
	}, nil
}
