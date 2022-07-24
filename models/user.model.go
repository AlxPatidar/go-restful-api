package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"emai,omitempty"`
	IsActive  bool               `json:"-" bson:"isActive"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Website   string             `json:"website,omitempty" bson:"website,omitempty"`
	IsDeleted bool               `json:"-" bson:"isDeleted"`
}

// get all user list item from user collection
func Users(client *mongo.Collection) []primitive.M {
	cursors, error := client.Find(context.Background(), bson.M{"isDelete": false})
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

// get user information based on user id
func UserOne(client *mongo.Collection, userId primitive.ObjectID) primitive.M {
	filter := bson.M{"_id": userId}
	var result primitive.M
	error := client.FindOne(context.Background(), filter).Decode(&result)
	if error != nil {
		return nil
	}
	return result
}

func CreateUser(client *mongo.Collection, user primitive.M) bool {
	_, error := client.InsertOne(context.Background(), user)
	if error != nil {
		return false
	}
	return true
}

// update user detail based item update values
func UpdateUser(client *mongo.Collection, userId primitive.ObjectID, user primitive.M) bool {
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": user}
	_, error := client.UpdateByID(context.Background(), filter, update)
	if error != nil {
		return false
	}
	return true
}

// delete user from users collection using update is deleted key
func DeleteUser(client *mongo.Collection, userId primitive.ObjectID) bool {
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"isDeleted": true}}
	error := client.FindOneAndUpdate(context.Background(), filter, update)
	if error != nil {
		return false
	}
	return true
}
