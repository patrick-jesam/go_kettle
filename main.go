package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"kettle/config/db"
	"kettle/pkg"
	"log"
	"net/http"
	"os"
)

func RegisterServices(db *db.MongoDB, router *mux.Router) {
	pkg.InitializeAccountService(db, router)
}

func main() {
	log.Println("welcome to kettle go boilerplate for grpc and rest services")

	_ = godotenv.Load()
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	dbConn := db.NewMongoConnection()

	RegisterServices(dbConn, router)

	log.Printf("server running on port %v", port)

	err := http.ListenAndServe(":"+port, nil)

	log.Println(err)
}
