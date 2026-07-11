package main

import (
	"golang-banking-api/app"
	"golang-banking-api/controller"
	"golang-banking-api/helper"
	"golang-banking-api/middleware"
	"golang-banking-api/repository"
	"golang-banking-api/service"
	"golang-banking-api/model/domain"
	"net/http"

	"github.com/go-playground/validator/v10"
)


func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepo := repository.NewMySQLUserRepository(db)
	tokenRepo := repository.NewMySQLTokenRepository(db)
	authUsecase := service.NewAuthUsecase(userRepo, tokenRepo)
	authHandler := controller.NewAuthHandler(authUsecase)
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(authHandler, userController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.AuthRoleMiddleware(domain.RoleAdmin, domain.RoleUser)(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}