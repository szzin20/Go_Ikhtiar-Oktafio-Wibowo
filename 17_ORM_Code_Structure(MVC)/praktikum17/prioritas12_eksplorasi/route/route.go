package route

import (
	"rest/controllers"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBookByID)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)

	e.GET("/blogs", controllers.GetAllBlogs)
	e.GET("/blogs/:id", controllers.GetBlogByID)
	e.POST("/blogs", controllers.CreateBlog)
	e.PUT("/blogs/:id", controllers.UpdateBlog)
	e.DELETE("/blogs/:id", controllers.DeleteBlog)
}
