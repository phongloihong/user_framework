package db

import (
	"fmt"
	"log"
	"time"

	"github.com/phongloihong/user_framework/config"
	"golang.org/x/net/context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Database *mongo.Database

func init() {
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		config.Config.MongoDB.User,
		config.Config.MongoDB.Password,
		config.Config.MongoDB.Host,
		config.Config.MongoDB.Port,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	Database = client.Database(config.Config.MongoDB.Name)
}
