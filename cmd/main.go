package main

import (
	"log"
	"miras/internal/handlers"
	"miras/internal/repository"
	"miras/internal/services"
)

func main() {
	db, err := repository.OpenDB()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	handler.Router()

	handler.Gin.Run("localhost:8000")
}
