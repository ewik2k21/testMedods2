package provider

import (
	"testMedods2/cmd/server"
	"testMedods2/internals/handler"
	"testMedods2/internals/repository"
	"testMedods2/internals/routes"
	"testMedods2/internals/services"

	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	tokenService := services.NewTokenService(userRepo)
	userHandler := handler.NewUserHandler(userService, tokenService)
	routes.RegistterUserRoutes(server, userHandler)
}
