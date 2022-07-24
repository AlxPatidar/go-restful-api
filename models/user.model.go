package models

import (
	"context"
	"log"
	"os/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"emai,omitempty"`
	IsActive bool               `json:"isActive,omitempty" bson:"isActive,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Website  string             `json:"website,omitempty" bson:"website,omitempty"`
}

func CreateUser(user user.User) {
}

func Users(client *mongo.Database) []primitive.M {
	cursors, error := client.Collection("users").Find(context.Background(), bson.D{{}})
	if error != nil {
		log.Fatal(error)
	}
	var users []primitive.M
	defer cursors.Close(context.Background())
	for cursors.Next(context.Background()) {
		var user bson.M
		err := cursors.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
