package repositories

import (
	"context"

	domain "THE_DEV_ARCH/Domain"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCol = "User"

type userRepo struct {
	dbName      string
	mongoClient *mongo.Client
	redisClient *redis.Client
	mongorepo   *userRepoMongo
	redisRepo   *userRepoRedis
}

func NewUserRepo(dbName string, redisClient *redis.Client, mongoClient *mongo.Client) *userRepo {
	return &userRepo{
		dbName:      dbName,
		redisClient: redisClient,
		mongoClient: mongoClient,
		mongorepo:   NewuserRepoMongo(dbName, mongoClient),
		redisRepo:   NewuserRepoRedis(redisClient),
	}
}
func (repo *userRepo) AddUser(ctx context.Context, user domain.User) (id string, err error) {
	// add user to Mongodb
	id, err = repo.mongorepo.AddUser(ctx, user)
	if err != nil {
		return
	}
	// add user to redis
	user.ID = id
	_, err = repo.redisRepo.AddUser(ctx, user)
	if err != nil {
		return
	}
	return
}
func (repo *userRepo) GetUser(ctx context.Context, id string) (user domain.User, err error) {
	//get user from Redis
	user, err = repo.redisRepo.GetUser(ctx, id)
	if err != nil {
		if err == domain.ErrNoDocumentFound {
			// get from mongodb
			user, err = repo.mongorepo.GetUser(ctx, id)
			if err != nil {
				return
			}
			// add to redis
			_, err = repo.redisRepo.AddUser(ctx, user)
			if err != nil {
				return
			}
		}
		return
	}
	return
}
