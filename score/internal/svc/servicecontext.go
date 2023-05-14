package svc

import (
	"database/sql"
	"github.com/generalzy/zeropan/score/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config config.Config
	DB     *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dtm_barrier")
	return &ServiceContext{
		Config: c,
		DB:     DB,
	}
}
