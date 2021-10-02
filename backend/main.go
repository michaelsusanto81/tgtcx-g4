package main

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/tgtcx-g4/coupon/backend/database"
	"github.com/tgtcx-g4/coupon/backend/server"
)

func main() {

	// Init database connection
	database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
