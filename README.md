# Book API with Go, Gin, and PostgreSQL

## Overview
This project is a simple RESTful API built with **Go** and **Gin**, using **PostgreSQL** for the database.  
It supports **JWT authentication** for user login and CRUD operations for **Books** and **Categories**.

---

## Features

### User
- `POST /api/users/login` → User login with JWT authentication

### Books
- `GET /books` → Get all books
- `POST /books` → Add a new book (JWT required)
- `GET /books/:id` → Get book details by ID
- `PUT /books/:id` → Update book by ID (JWT required)
- `DELETE /books/:id` → Delete book by ID (JWT required)

### Categories
- `GET /categories` → Get all categories
- `POST /categories` → Add a new category (JWT required)
- `POST /categories/:id` → Get detail category by ID
- `GET /categories/:id/books` → Get all books under a category
- `PUT /categories/:id` → Update category by ID (JWT required)
- `DELETE /categories/:id` → Delete category by ID (JWT required)

---
