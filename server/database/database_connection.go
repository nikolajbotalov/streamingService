package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

func Connect() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")

	if MongoDb == "" {
		log.Fatalln("MONGODB_URI not set")
	}

	fmt.Printf("MongoDB URI:%s", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Println("Error connecting to MongoDB")
	}

	return client
}

func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	dataBaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE_NAME: %s", dataBaseName)

	collection := client.Database(dataBaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}

	return collection
}
