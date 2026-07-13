package app

import (
	"golang-banking-api/controller"
	"golang-banking-api/exception"
	"golang-banking-api/middleware"
	"golang-banking-api/model/domain"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(authHandler *controller.AuthHandler, userController controller.UserController, accountController controller.AccountController, transactionController controller.TransactionController) *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc("POST", "/register", authHandler.Register)
	router.HandlerFunc("POST", "/login", authHandler.Login)
	router.HandlerFunc("POST", "/refresh", authHandler.Refresh)
	router.HandlerFunc("POST", "/logout", authHandler.Logout)

	adminRouter := httprouter.New()
	adminRouter.GET("/users", userController.FindAll)
	adminRouter.GET("/users/:userId", userController.FindById)
	adminRouter.POST("/users", userController.Create)
	adminRouter.PUT("/users/:userId", userController.Update)
	adminRouter.DELETE("/users/:userId", userController.Delete)
	adminRouter.GET("/accounts", accountController.FindAll)
	adminRouter.GET("/accounts/:accountId", accountController.FindById)
	adminRouter.POST("/accounts", accountController.Create)
	adminRouter.PUT("/accounts/:accountId", accountController.Update)
	adminRouter.DELETE("/accounts/:accountId", accountController.Delete)
	adminRouter.GET("/transactions", transactionController.FindAll)
	adminRouter.GET("/transactions/:transactionId", transactionController.FindById)
	adminRouter.POST("/transactions", transactionController.Create)
	adminRouter.PUT("/transactions/:transactionId", transactionController.Update)
	adminRouter.DELETE("/transactions/:transactionId", transactionController.Delete)

	adminHandler := middleware.AuthRoleMiddleware(domain.RoleAdmin)(middleware.StripPrefix("/api/admin", adminRouter))
	for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		router.Handler(method, "/api/admin/*path", adminHandler)
	}

	userRouter := httprouter.New()
	userRouter.GET("/profile", userController.GetMe)
	userRouter.PATCH("/profile", userController.UpdateMe)
	userRouter.PUT("/profile", userController.UpdateMe)
	userRouter.GET("/accounts", accountController.GetMyAccounts)
	userRouter.GET("/accounts/:accountId", accountController.GetMyAccountById)
	userRouter.POST("/accounts", accountController.CreateMyAccount)
	userRouter.PATCH("/accounts/:accountId", accountController.UpdateMyAccount)
	userRouter.DELETE("/accounts/:accountId", accountController.DeleteMyAccount)
	userRouter.POST("/transactions/deposit", transactionController.Deposit)
	userRouter.POST("/transactions/withdraw", transactionController.Withdraw)
	userRouter.POST("/transactions/transfer", transactionController.Transfer)
	userRouter.GET("/transactions", transactionController.GetMyTransactions)
	userRouter.GET("/transactions/:transactionId", transactionController.GetMyTransactionById)

	userHandler := middleware.AuthRoleMiddleware(domain.RoleAdmin, domain.RoleUser)(middleware.StripPrefix("/api/user", userRouter))
	for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		router.Handler(method, "/api/user/*path", userHandler)
	}

	router.PanicHandler = exception.ErrorHandler

	swaggerHandler := NewSwaggerHandler()
	router.Handler("GET", "/swagger.json", swaggerHandler)
	router.Handler("GET", "/swagger/", swaggerHandler)

	return router
}
