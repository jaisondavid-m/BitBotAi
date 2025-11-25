package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func ConnectMongo() {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	if uri == "" {
		log.Fatal("❌ MONGO_URI missing in .env")
	}
	if dbName == "" {
		log.Fatal("❌ MONGO_DB missing in .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Could not ping MongoDB:", err)
	}

	Mongo = client.Database(dbName)
	fmt.Println("✅ MongoDB connected successfully")
}
