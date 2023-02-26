package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Connect a database
func Start() {
	// Load dot env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Create client option
	option := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Create a client
	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		log.Fatal(err)
	}

	// Create context to check client is running
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check client created successfully
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a database with this client
	DB = client.Database("devnotes")
	fmt.Println("\"" + DB.Name() + "\" database connected successful")
}

// Create a collection in database and return that
func Collection(s string) *mongo.Collection {
	return DB.Collection(s)
}
