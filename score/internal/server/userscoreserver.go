// Code generated by goctl. DO NOT EDIT.
// Source: score.proto

package server

import (
	"context"

	"github.com/generalzy/zeropan/score/internal/logic"
	"github.com/generalzy/zeropan/score/internal/svc"
	"github.com/generalzy/zeropan/score/types/score"
)

type UserScoreServer struct {
	svcCtx *svc.ServiceContext
	score.UnimplementedUserScoreServer
}

func NewUserScoreServer(svcCtx *svc.ServiceContext) *UserScoreServer {
	return &UserScoreServer{
		svcCtx: svcCtx,
	}
}

func (s *UserScoreServer) SaveUserScore(ctx context.Context, in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	l := logic.NewSaveUserScoreLogic(ctx, s.svcCtx)
	return l.SaveUserScore(in)
}

func (s *UserScoreServer) SaveUserScoreCallback(ctx context.Context, in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	l := logic.NewSaveUserScoreCallbackLogic(ctx, s.svcCtx)
	return l.SaveUserScoreCallback(in)
}