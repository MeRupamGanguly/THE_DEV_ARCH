package domain

import (
	"context"
	"errors"
)

var (
	ErrNoDocumentFound = errors.New("no documents found")
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Age      int    `json:"age" bson:"age"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserRepositoryContracts interface {
	AddUser(ctx context.Context, user User) (id string, err error)
	GetUser(ctx context.Context, id string) (user User, err error)
}
