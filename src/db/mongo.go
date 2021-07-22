package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoDBConnect(collectionName string) *mongo.Collection {
	client, err := MongoDBConnect(collectionName)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetMongoDBDisConnect(client *mongo.Collection) {
	client.Database().Client().Disconnect(context.TODO())
}
