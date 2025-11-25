package storage

import (
	"context"
	"library/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Material struct {
	Content string `bson:"content"`
}

func SaveMaterial(text string) error {
	collection := config.Mongo.Collection("study_material")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{
		"content": text,
	})
	return err
}

func GetMaterial() ([]string, error) {
	collection := config.Mongo.Collection("study_material")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []string

	for cursor.Next(ctx) {
		var doc Material
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		result = append(result, doc.Content)
	}

	return result, nil
}
