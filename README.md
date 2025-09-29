# Api-Book
Create Api Book Based Requirement


Api Route User : 

Using Middleware Jwt 

router.POST("/api/users/login", auth.LoginUser)


API Route Book : 

cat.GET("/books", book.GetAllBook)
		cat.POST("/books", book.AddBook)
		cat.GET("/books/:id", book.GetBookDetail)
		cat.DELETE("/books/:id", book.DeleteBook)
		cat.PUT("/books/:id", book.UpdateBook)


API Route Categories : 

cat.GET("/categories", category.Getallcategories)
		cat.POST("/categories", category.AddCategory)
		cat.POST("/categories/:id", category.GetDetailCategory)
		cat.DELETE("/categories/:id", category.DeleteCategory)
		cat.GET("/categories/:id/books", category.GetBookByCategory)
		cat.PUT("/categories/:id", category.UpdateDataCategory)
