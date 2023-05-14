package models

import "time"

// Share 分享
type Share struct {
	UserID       string    `bson:"user_id"`
	RepositoryID string    `bson:"repository_id"`
	ExpiredTime  time.Time `bson:"expired_time"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at"`
}
