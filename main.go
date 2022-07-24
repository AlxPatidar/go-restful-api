package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
func main() {
	fmt.Println("Go Mongo Restful API")
	port := os.Getenv("PORT")
	fmt.Println("Go server is started on", port)
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":"+port, router)
}
