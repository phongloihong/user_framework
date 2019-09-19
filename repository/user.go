package repository

import (
	"context"
	"sync"

	"github.com/phongloihong/user_framework/utils"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/phongloihong/user_framework/db"
	"github.com/phongloihong/user_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	collection *mongo.Collection
}

var (
	Once     sync.Once
	UserRepo *userRepo
)

func init() {
	Once.Do(func() {
		UserRepo = &userRepo{db.Database.Collection("users")}
		db.Database.Collection("users").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		})
	})
}

func (u *userRepo) Create(req types.UserAddReq) (*mongo.InsertOneResult, error) {
	hashedPassword, err := utils.Hash([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	return u.collection.InsertOne(context.TODO(), bson.M{
		"email":    req.Email,
		"password": hashedPassword,
	})
}

func (u *userRepo) GetByEmail(email string) *mongo.SingleResult {
	return u.collection.FindOne(context.TODO(), bson.M{"email": email})
}
