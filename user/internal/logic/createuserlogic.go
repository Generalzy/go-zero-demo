package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/generalzy/zeropan/models"
	"github.com/generalzy/zeropan/user/internal/svc"
	"github.com/generalzy/zeropan/user/types/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		// 重试
		return nil, status.Error(codes.Internal, err.Error())
	}

	now := time.Now()
	data := models.User{
		Name:      in.Name,
		Password:  "",
		Email:     in.Email,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: now,
	}
	err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		// 此处需要用tx或其他链接写实现saveScore的事务逻辑
		if _, err := l.svcCtx.MongoClient.Database("zeropan").Collection("user").InsertOne(context.TODO(), data); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		// 回滚状态
		return nil, status.Error(codes.Aborted, err.Error())
	}

	//if result, err := l.svcCtx.MongoClient.Database("zeropan").Collection("user").InsertOne(context.TODO(), data); err != nil {
	//	return nil, err
	//} else {
	//	return &user.CreateUserResponse{
	//		Id:       result.InsertedID.(primitive.ObjectID).String(),
	//		Name:     in.Name,
	//		Email:    in.Email,
	//		CreateAt: now.Format("2006-01-02 15-04-05"),
	//		UpdateAt: now.Format("2006-01-02 15-04-05"),
	//	}, nil
	//}
	return &user.CreateUserResponse{
		Id:       "xxxx",
		Name:     in.Name,
		Email:    in.Email,
		CreateAt: now.Format("2006-01-02 15-04-05"),
		UpdateAt: now.Format("2006-01-02 15-04-05"),
	}, nil
}
