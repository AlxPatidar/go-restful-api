package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	error := godotenv.Load("config/.env")
	if error != nil {
		log.Fatal("Error loading .env file.")
		fmt.Printf("Error loading .env file %s", error.Error())
		return
	}
}

// welcome page to check server is runing
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

}
func Start() {
	// get port from env file
	port := os.Getenv("PORT")
	fmt.Println("Go server is started on", port)
	// create router using mux
	router := mux.NewRouter()
	// add router path
	router.HandleFunc("/", helloWorldHandler)
	// run server on env port
	http.ListenAndServe(":"+port, router)
}
