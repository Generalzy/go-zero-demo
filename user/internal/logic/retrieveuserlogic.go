package logic

import (
	"context"
	"github.com/generalzy/zeropan/models"
	"github.com/generalzy/zeropan/user/internal/svc"
	"github.com/generalzy/zeropan/user/types/user"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RetrieveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveUserLogic {
	return &RetrieveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveUserLogic) RetrieveUser(in *user.RetrieveUserRequest) (*user.RetrieveUserResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, err
	}

	result := l.svcCtx.MongoClient.Database("zeropan").Collection("user").FindOne(
		context.TODO(), bson.M{"_id": objectID})
	if result.Err() != nil {
		return nil, result.Err()
	}

	u := new(models.User)
	if err := result.Decode(u); err != nil {
		return nil, err
	}
	return &user.RetrieveUserResponse{
		Id:       in.Id,
		Name:     u.Name,
		Email:    u.Email,
		CreateAt: u.CreatedAt.Format("2006-01-02 15-04-05"),
		UpdateAt: u.UpdatedAt.Format("2006-01-02 15-04-05"),
	}, nil
}
