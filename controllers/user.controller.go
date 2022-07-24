package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlxPatidar/go-restful-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Database
}

func NewUserController(client *mongo.Database) *UserController {
	return &UserController{client}
}

// get all users
func (userController UserController) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all user")
	res.Header().Add("Content-Type", "application/json")
	users := models.Users(userController.client)
	json.NewEncoder(res).Encode(users)
}

// get user information based on user id
func (userController UserController) GetUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get user")
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("user")
}

// create new user
func (userController UserController) CreateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Create user")
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("user")
}

// update user detail
func (userController UserController) UpdateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Update user")
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("user")
}

// update user detail
func (userController UserController) DeleteUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Delete user")
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode("user")
}
