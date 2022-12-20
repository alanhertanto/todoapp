// main.go

package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"todoapp/database"
	"todoapp/controllers"
	"todoapp/platform/authenticator"
	"todoapp/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

	// Load Configurations from config.json using Viper
	LoadAppConfig()
	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	RegisterPostRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterPostRoutes(router *mux.Router) {
	router.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controllers.GetPosById).Methods("GET")
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("DELETE")
}
