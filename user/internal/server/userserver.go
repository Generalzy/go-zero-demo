// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/generalzy/zeropan/user/internal/logic"
	"github.com/generalzy/zeropan/user/internal/svc"
	"github.com/generalzy/zeropan/user/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) RetrieveUser(ctx context.Context, in *user.RetrieveUserRequest) (*user.RetrieveUserResponse, error) {
	l := logic.NewRetrieveUserLogic(ctx, s.svcCtx)
	return l.RetrieveUser(in)
}

func (s *UserServer) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServer) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UserServer) CreateUserCallback(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	l := logic.NewCreateUserCallbackLogic(ctx, s.svcCtx)
	return l.CreateUserCallback(in)
}