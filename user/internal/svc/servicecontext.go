package svc

import (
	"context"
	"database/sql"
	"github.com/generalzy/zeropan/user/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	clientOptions := options.Client().ApplyURI(c.MongoUri.Uri)
	clientOptions.SetMaxPoolSize(100)
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	DB, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dtm_barrier")
	return &ServiceContext{
		Config:      c,
		MongoClient: client,
		DB:          DB,
	}
}
