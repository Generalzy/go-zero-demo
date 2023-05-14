package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/generalzy/zeropan/score/internal/svc"
	"github.com/generalzy/zeropan/score/types/score"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreLogic {
	return &SaveUserScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreLogic) SaveUserScore(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		// 重试
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		// 此处需要用tx或其他链接写实现saveScore的事务逻辑
		logx.Info("saveUserScore called>>>>>>>>>>>>>")
		return errors.New("this is a error")
	})
	if err != nil {
		// 回滚状态
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &score.UserScoreResponse{
		Id:    in.Id + "100",
		Score: in.Score + 100,
	}, nil
}
