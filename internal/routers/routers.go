package routers

import (
	"api_admin_server/internal/controllers"
	"api_admin_server/internal/middleware"
	"api_admin_server/store"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Init() *gin.Engine {
	r := gin.Default()
	config := new(store.Config)
	gin.SetMode(config.Get().Runmode)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "*",
		ExposedHeaders:  "*",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.GET("health", func(context *gin.Context) {
		context.JSON(200, map[string]string{"msg": "ok"})
		return
	})

	userController := new(controllers.UserController)

	userOpenApi := r.Group("open/user")
	userOpenApi.POST("code/email", userController.SendEmailCode)
	userOpenApi.POST("register", userController.Register)
	userOpenApi.POST("login", userController.Login)
	userOpenApi.POST("password/back", userController.PasswordBack)

	//
	userApi := r.Group("user").Use(middleware.VerifyToken())
	userApi.POST("data/full", userController.DataFull)

	publicController := new(controllers.PublicController)
	publicApi := r.Group("public").Use(middleware.VerifyToken())
	publicApi.POST("file/upload", publicController.UploadFile)

	return r
}
