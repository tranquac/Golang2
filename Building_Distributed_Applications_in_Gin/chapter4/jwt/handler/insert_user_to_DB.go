package handler

import (
	"context"
	"crypto/sha256"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	users := map[string]string{
		"admin": "fCRmh4Q2J7Rseqkz",
		"quac":  "RE4zfHB35VPtTkbT",
		"hong":  "L3nSFRcZzNQ67bcc",
	}
	ctx := context.Background()
	MONGO_DATABASE := "demo"
	MONGODB_URI := "mongodb://admin:password@103.81.86.132:27017/demo?authSource=admin"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGODB_URI))

	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(MONGO_DATABASE).Collection("users")
	h := sha256.New()
	for username, password := range users {
		collection.InsertOne(ctx, bson.M{
			"username": username,
			"password": string(h.Sum([]byte(password))),
		})
	}
}
