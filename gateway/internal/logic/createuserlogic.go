package logic

import (
	"context"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/generalzy/zeropan/score/types/score"
	"github.com/generalzy/zeropan/user/types/user"

	"github.com/generalzy/zeropan/gateway/internal/svc"
	"github.com/generalzy/zeropan/gateway/internal/types"

	_ "github.com/dtm-labs/dtmdriver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
)

var dtmServer = "etcd://localhost:2379/dtmservice"

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (resp *types.CreateUserResponse, err error) {

	gid := dtmgrpc.MustGenGid(dtmServer)
	sagaGrpc := dtmgrpc.NewSagaGrpc(dtmServer, gid)
	userServer, err := l.svcCtx.Config.UserRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	scoreServer, err := l.svcCtx.Config.ScoreRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	userReq := &user.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		CreateAt: "",
		UpdateAt: "",
	}
	// dtm添加user
	sagaGrpc.Add(userServer+"/user.User/createUser", userServer+"/user.User/createUserCallback", userReq)
	// dtm添加score
	scoreReq := &score.UserScoreRequest{
		// Id可以另写一个获取ID的服务
		Id:    "0x12231",
		Score: 0,
	}
	sagaGrpc.Add(scoreServer+"/score.UserScore/saveUserScore", scoreServer+"/score.UserScore/saveUserScoreCallback", scoreReq)
	sagaGrpc.WaitResult = true
	err = sagaGrpc.Submit()
	if err != nil {
		return nil, err
	}
	//response, err := l.svcCtx.UserRpc.CreateUser(context.TODO(), &user.CreateUserRequest{
	//	Name:     req.Name,
	//	Email:    req.Email,
	//	CreateAt: "",
	//	UpdateAt: "",
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//scoreResponse, err := l.svcCtx.ScoreRpc.SaveUserScore(context.TODO(), &score.UserScoreRequest{
	//	Id:    response.Id,
	//	Score: 0,
	//})
	//if err != nil {
	//	return nil, err
	//}
	return &types.CreateUserResponse{
		ID:        scoreReq.Id,
		Name:      userReq.Name,
		Email:     userReq.Email,
		CreatedAt: userReq.CreateAt,
		UpdatedAt: userReq.UpdateAt,
	}, nil
}
