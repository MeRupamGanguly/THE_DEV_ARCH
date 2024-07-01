package main

import (
	users "THE_DEV_ARCH/Users"
	"THE_DEV_ARCH/Users/repositories"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBNAME = "RGGMA"
)

func main() {
	// Config
	// Log
	// Consul
	// Redis connection
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	// MongoDB connection
	mongoOption := options.Client().ApplyURI("mongodb://admin:adminpass@localhost:27017/")
	mongoContext, contextCancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer contextCancle()
	mongoClient, err := mongo.Connect(mongoContext, mongoOption)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repositories.NewUserRepo(DBNAME, redisClient, mongoClient)
	usersvc := users.NewUserService(userRepo)
	// id, err := usersvc.AddUser(context.Background(), domain.User{Name: "Rupam", Age: 28, Email: "rg@ganguly.com", Password: "SonuSud"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	user, err := usersvc.GetUser(context.Background(), "66744eb72325d5d89f2fb750")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
