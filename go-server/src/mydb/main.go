package mydb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type T struct {
	Client         *mongo.Client
	Database       *mongo.Database
	UserCollection *mongo.Collection
	mongodbURI     string
}

func Init(mongodbURI string) *T {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("golangdb")

	return &T{
		Client:         client,
		Database:       database,
		UserCollection: database.Collection("test"),
	}
}

type MongoFields struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`

	EMAIL string `bson:"email,omitempty"`
	NOTE  string `bson:"note,omitempty"`
}

func (data T) GetOrCreateUser(email string) primitive.ObjectID {
	var result MongoFields
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := data.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne() ObjectIDFromHex ERROR:", err)
		result2, _ := data.UserCollection.InsertOne(ctx, bson.D{{Key: "email", Value: email}})
		return result2.InsertedID.(primitive.ObjectID)
	}
	return result.ID
}

func (data T) GetUser(id primitive.ObjectID) MongoFields {
	var result MongoFields
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := data.UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne() ObjectIDFromHex ERROR:", err)
	}
	return result
}

func (data T) UpdateUserNote(id primitive.ObjectID, note string) {
	var result MongoFields
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := data.UserCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"note": note}},
	).Decode(&result)
	if err != nil {
		fmt.Println("FindOne() ObjectIDFromHex ERROR:", err)
	}
}
