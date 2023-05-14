package models

import "time"

// User 用户信息
type User struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
	Email    string `bson:"email"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at"`
}
