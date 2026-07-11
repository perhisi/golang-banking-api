package app

import (
	"golang-banking-api/controller"
	"golang-banking-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(authHandler *controller.AuthHandler, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc("POST", "/register", authHandler.Register)
	router.HandlerFunc("POST", "/login", authHandler.Login)
	router.HandlerFunc("POST", "/refresh", authHandler.Refresh)
	router.HandlerFunc("POST", "/logout", authHandler.Logout)

	router.GET("/api/admin/users", userController.FindAll)
	router.GET("/api/admin/users/:userId", userController.FindById)
	router.POST("/api/admin/users", userController.Create)
	router.PUT("/api/admin/users/:userId", userController.Update)
	router.DELETE("/api/admin/users/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}