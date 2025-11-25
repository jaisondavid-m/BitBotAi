package storage

import (
	"context"
	"library/config"
	"library/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func FindUserByEmail(email string) (*models.User, error) {
	collection := config.Mongo.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func CreateUser(user models.User) error {
	collection := config.Mongo.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	return err
}
