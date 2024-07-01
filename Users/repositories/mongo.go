package repositories

import (
	domain "THE_DEV_ARCH/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepoMongo struct {
	dbName      string
	mongoClient *mongo.Client
}

func NewuserRepoMongo(dbName string, mongoClient *mongo.Client) *userRepoMongo {
	return &userRepoMongo{
		dbName:      dbName,
		mongoClient: mongoClient,
	}
}

func (repo *userRepoMongo) AddUser(ctx context.Context, user domain.User) (id string, err error) {
	res, err := repo.mongoClient.Database(repo.dbName).Collection(userCol).InsertOne(ctx, user)
	if err != nil {
		return
	}
	return res.InsertedID.(string), nil
}
func (repo *userRepoMongo) GetUser(ctx context.Context, id string) (user domain.User, err error) {
	filter := bson.M{"_id": id}
	err = repo.mongoClient.Database(repo.dbName).Collection(userCol).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = domain.ErrNoDocumentFound
			return
		}
		return
	}
	return
}
