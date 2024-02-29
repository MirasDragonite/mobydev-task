package main

import (
	"miras/internal/handlers"
	"miras/internal/repository"
	"miras/internal/services"
)

func main() {
	db, err := repository.OpenDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	handler.Router()

	err = handler.Gin.Run("localhost:8000")
	if err != nil {
		panic(err)
	}
}
