package main

import (
	"golang-banking-api/app"
	"golang-banking-api/controller"
	"golang-banking-api/helper"
	"golang-banking-api/repository"
	"golang-banking-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := app.NewDB()
	validate := validator.New()
	userRepo := repository.NewMySQLUserRepository(db)
	tokenRepo := repository.NewMySQLTokenRepository(db)
	authUsecase := service.NewAuthUsecase(userRepo, tokenRepo)
	authHandler := controller.NewAuthHandler(authUsecase)
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	accountRepository := repository.NewAccountRepository()
	accountService := service.NewAccountService(accountRepository, userRepository, db, validate)
	accountController := controller.NewAccountController(accountService)
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, accountRepository, userRepository, db, validate)
	transactionController := controller.NewTransactionController(transactionService)
	router := app.NewRouter(authHandler, userController, accountController, transactionController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
