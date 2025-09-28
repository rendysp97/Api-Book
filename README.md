# Api-Book
Create Api Book Based Requirement

 POST   /api/categories           --> api-book/controller/category.AddCategory (4 handlers)
 POST   /api/categories/:id       --> api-book/controller/category.GetDetailCategory (4 handlers)
 DELETE /api/categories/:id       --> api-book/controller/category.DeleteCategory (4 handlers)
 GET    /api/categories/:id/books --> api-book/controller/category.GetBookByCategory (4 handlers)
 PUT    /api/categories/:id       --> api-book/controller/category.UpdateDataCategory (4 handlers)

 
 GET    /api/books                --> api-book/controller/book.GetAllBook (4 handlers)
 POST   /api/books                --> api-book/controller/book.AddBook (4 handlers)
 GET    /api/books/:id            --> api-book/controller/book.GetBookDetail (4 handlers)
 DELETE /api/books/:id            --> api-book/controller/book.DeleteBook (4 handlers)
 PUT    /api/books/:id            --> api-book/controller/book.UpdateBook (4 handlers)
