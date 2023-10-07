package route

import (
	"clean/constant"
	"clean/controller"
	"clean/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(app *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepo(db)
	userController := controller.NewUserController(userRepository)

	AppJWT := app.Group("/users")
	AppJWT.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	AppJWT.GET("", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
}
