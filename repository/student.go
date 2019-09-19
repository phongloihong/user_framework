package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/phongloihong/user_framework/sequence"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/phongloihong/user_framework/db"
	"github.com/phongloihong/user_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var collectionSequence *mongo.Collection

func init() {
	collection = db.Database.Collection("student")
	collectionSequence = db.Database.Collection("student_sequence")
}

func Fetch() ([]*types.Student, error) {
	var students []*types.Student
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	// cur.All(context.TODO(), &students)
	for cur.Next(context.TODO()) {
		var student types.Student
		err := cur.Decode(&student)
		if err != nil {
			return nil, err
		}

		students = append(students, &student)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())

	return students, nil
}

func Insert(studentRequest types.StudentAddReq) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nextId, err := sequence.GetNextID(collectionSequence, "student")
	if err != nil {
		return nil, err
	}

	return collection.InsertOne(ctx, bson.M{
		"id":         nextId,
		"first_name": studentRequest.FirstName,
		"last_name":  studentRequest.LastName,
		"class_name": studentRequest.ClassName,
	})
}

func GetById(studentId primitive.ObjectID) (types.Student, error) {
	var student types.Student
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": studentId}).Decode(&student)

	return student, err
}

func UpdateById(studentId primitive.ObjectID, studentReq types.StudentUpdateReq) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.UpdateOne(
		ctx,
		bson.M{"_id": studentId},
		bson.M{
			"$set": bson.M{
				"first_name": studentReq.FirstName,
				"last_name":  studentReq.LastName,
				"class_name": studentReq.ClassName,
			},
		})
}

func DeleteById(studentId primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return collection.DeleteOne(ctx, bson.M{"_id": studentId})
}

func Find(studentRequest types.StudentSearchRequest) ([]*types.Student, error) {
	findOptions := options.Find()
	findOptions.SetLimit(30)

	var filter map[string]interface{}
	bs, _ := json.Marshal(studentRequest)
	json.Unmarshal(bs, &filter)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var students []*types.Student
	if err := cur.All(context.TODO(), &students); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	return students, nil
}

func GroupStudent(name string) (*[]map[string]interface{}, error) {
	var students []map[string]interface{}

	pipeline := bson.A{
		bson.M{"$match": bson.M{"first_name": name}},
		bson.M{"$group": bson.M{
			"_id":        "$last_name",
			"first_name": bson.M{"$push": "$first_name"},
			"ids":        bson.M{"$push": "$id"},
		}},
	}
	cur, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}

	if err = cur.All(context.TODO(), &students); err != nil {
		return nil, err
	}

	return &students, nil
}

// func getFilters(studentRequest types.StudentSearchRequest) (bson.D, error) {
// 	filter := bson.M{}

// 	if studentRequest.Id != "" {
// 		objectID, err := primitive.ObjectIDFromHex(studentRequest.Id)
// 		if err != nil {
// 			return nil, err
// 		}

// 		filter["_id"] = objectID
// 	}

// 	if studentRequest.ClassName != "" {
// 		filter["class_name"] = primitive.Regex{Pattern: studentRequest.ClassName, Options: "i"}
// 	}

// 	if req.Name != "" {
// 		filter["$or"] = bson.A{
// 			bson.M{"first_name": primitive.Regex{Pattern: studentRequest.FirstName, Options: "i"}},
// 			bson.M{"last_name": primitive.Regex{Pattern: studentRequest.LastName, Options: "i"}},
// 		}
// 	}

// 	return filter, nil
// }
