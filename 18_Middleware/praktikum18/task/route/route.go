package route

import (
	"middleware/controllers"
	"middleware/middlewares"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {

	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)

	Protectusers := e.Group("/users", middlewares.JWTMiddleware())

	Protectusers.GET("", controllers.GetUsersController)
	Protectusers.GET("/:id", controllers.GetUserController)
	Protectusers.DELETE("/:id", controllers.DeleteUserController)
	Protectusers.PUT("/:id", controllers.UpdateUserController)

	books := e.Group("/books", middlewares.JWTMiddleware())

	books.GET("", controllers.GetBooksController)
	books.GET("/:id", controllers.GetBookController)
	books.POST("", controllers.CreateBookController)
	books.PUT("/:id", controllers.UpdateBookController)
	books.DELETE("/:id", controllers.DeleteBookController)

}
