package users

import (
	domain "THE_DEV_ARCH/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	repo domain.UserRepositoryContracts
}

func NewUserService(repo domain.UserRepositoryContracts) *userService {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) AddUser(ctx context.Context, user domain.User) (id string, err error) {
	user.ID = primitive.NewObjectID().Hex()
	id, err = svc.repo.AddUser(ctx, user)
	if err != nil {
		return
	}
	fmt.Println(id)
	return
}

func (svc *userService) GetUser(ctx context.Context, id string) (user domain.User, err error) {
	user, err = svc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}
