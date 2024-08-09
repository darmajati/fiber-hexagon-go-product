package config

import(
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    Client *mongo.Client
)

func InitMongoDB(uri string) *mongo.Client {
    clientOptions := options.Client().ApplyURI(uri)

    // Create a new client and connect to the server
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    log.Println("Connected to MongoDB!")
    Client = client
    return client
}