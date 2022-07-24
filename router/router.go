package router

import (
	"encoding/json"
	"net/http"

	"github.com/AlxPatidar/go-restful-api/config"
	"github.com/AlxPatidar/go-restful-api/controllers"
	"github.com/gorilla/mux"
)

// welcome page to check server is runing
func helloWorldHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	json.NewEncoder(res).Encode("Server is runing")
}

func Router() *mux.Router {
	// create router using mux
	router := mux.NewRouter()
	// add router path
	router.HandleFunc("/", helloWorldHandler)
	client := config.GetDatabase("users")
	userController := controllers.NewUserController(client)
	router.HandleFunc("/api/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{userId}", userController.GetUser).Methods("GET")
	router.HandleFunc("/api/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{userId}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{userId}", userController.DeleteUser).Methods("DELETE")
	return router
}
