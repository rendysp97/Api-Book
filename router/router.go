package router

import (
	"api-book/controller/auth"
	"api-book/controller/book"
	"api-book/controller/category"
	"api-book/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})

	router.POST("/api/users/login", auth.LoginUser)

	cat := router.Group("/api")

	cat.Use(middleware.AuthMiddleware())
	{
		cat.GET("/categories", category.Getallcategories)
		cat.POST("/categories", category.AddCategory)
		cat.POST("/categories/:id", category.GetDetailCategory)
		cat.DELETE("/categories/:id", category.DeleteCategory)
		cat.GET("/categories/:id/books", category.GetBookByCategory)
		cat.PUT("/categories/:id", category.UpdateDataCategory)

		//book

		cat.GET("/books", book.GetAllBook)
		cat.POST("/books", book.AddBook)
		cat.GET("/books/:id", book.GetBookDetail)
		cat.DELETE("/books/:id", book.DeleteBook)
		cat.PUT("/books/:id", book.UpdateBook)
	}

	return router
}
