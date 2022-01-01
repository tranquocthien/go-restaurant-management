package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB(user, password string) *MongoDB {
	strMongo := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.wn1j0.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", "tranquocthien", "Th184357617")
	client, err := mongo.NewClient(options.Client().ApplyURI(strMongo))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to database successfully")
	Mongo.Client = client
	return Mongo
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)

	return collection
}
