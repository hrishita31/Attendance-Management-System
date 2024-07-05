package infra

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Coll *mongo.Database
var SubjAtt *mongo.Database

func NewMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection successful")

	Client = client
	createIndexForMongo()
	return Client
}

func DisconnectMongo() {

	if err := Client.Disconnect(context.Background()); err != nil {
		fmt.Println(err)
	}

}

func GetClient() *mongo.Client {
	return Client
}

func GetDatabase() *mongo.Database {
	Coll := Client.Database("users_db")
	return Coll
}

func createIndexForMongo() {

	att := GetDatabase().Collection("attendance")
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "date", Value: 1},
			{Key: "username", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	name, err := att.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Error in creating index: %v", err)
	}
	fmt.Println("indexes created for: ", name)
}
