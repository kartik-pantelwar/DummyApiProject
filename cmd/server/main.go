package main

import (
	"dummyProject/external"
	"dummyProject/internal/adaptors/persistance"
	"dummyProject/internal/config"
	userhandler "dummyProject/internal/interfaces/handler"
	"dummyProject/internal/interfaces/middleware"
	"dummyProject/internal/interfaces/routes"
	user "dummyProject/internal/usecase"
	"dummyProject/pkg/migrate"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	database, err := persistance.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to Database: %v", err)
	}
	fmt.Println("Connected to database")

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory %v", err)
	}

	migrate := migrate.NewMigrate(
		database.GetDB(),
		cwd+"/migrations")

	err = migrate.RunMigrations()
	if err != nil {
		log.Fatalf("failed to run migrations %v", err)
	}

	userRepo := persistance.NewUserRepo(database)
	sessionRepo := persistance.NewSessionRepo(database)
	userService := user.NewUserService(userRepo, sessionRepo)
	userHandler := userhandler.NewUserHandler(userService)
	pHandler := external.NewHandler()

	router := routes.InitRoutes(&userHandler, pHandler)
	wrapper:= middleware.TimeoutMiddleware(200*time.Millisecond,router)

	configP, err := config.LoadConfig()
	if err != nil {
		panic("Unable to use port")
	}
	err = http.ListenAndServe(fmt.Sprintf(":%s", configP.APP_PORT), wrapper)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
