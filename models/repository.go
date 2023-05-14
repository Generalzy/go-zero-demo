package models

import "time"

// Repository 个人存储池
type Repository struct {
	// 用户ID
	UserID string `bson:"user_id"`
	// 上级ID
	ParentID string `bson:"parent_id"`
	// 扩展名
	Ext string `bson:"ext"`
	// 文件
	FileID string `bson:"file_id"`
	Name   string `bson:"name"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at"`
}
