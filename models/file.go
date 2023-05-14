package models

import "time"

// File 公共存储池
type File struct {
	Hash string `bson:"hash"`
	Name string `bson:"name"`
	// Ext 文件扩展名
	Ext  string  `bson:"ext"`
	Size float64 `bson:"size"`
	// Path 文件路径
	Path string `bson:"path"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at"`
}
