package repositories

import (
	domain "THE_DEV_ARCH/Domain"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type userRepoRedis struct {
	redisClient *redis.Client
}

func NewuserRepoRedis(redisClient *redis.Client) *userRepoRedis {
	return &userRepoRedis{
		redisClient: redisClient,
	}
}

func (repo *userRepoRedis) AddUser(ctx context.Context, user domain.User) (id string, err error) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return
	}
	err = repo.redisClient.Set(ctx, user.ID, string(userJson), 0).Err()
	if err != nil {
		return
	}
	return
}
func (repo *userRepoRedis) GetUser(ctx context.Context, id string) (user domain.User, err error) {
	res, err := repo.redisClient.Get(ctx, id).Result()
	if err != nil {
		if err == redis.Nil {
			err = domain.ErrNoDocumentFound
		}
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	return
}
