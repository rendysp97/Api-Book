package book

import (
	"api-book/database/connection"
	"api-book/model"
	"api-book/repository"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET ALL BOOK

func GetAllBook(ctx *gin.Context) {

	var alldata []model.Book

	repository.GetAllBookRepo(connection.Db, &alldata)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    alldata,
		"message": "success",
	})

}

// ADD BOOK

func AddBook(ctx *gin.Context) {

	var newData model.Book

	if err := ctx.ShouldBindJSON(&newData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newData.Total_page > 100 {
		newData.Thickness = "Tebal"
	} else {
		newData.Thickness = "Tipis"
	}

	if newData.Release_year < 1980 || newData.Release_year > 2024 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "release_year harus antara 1980 sampai 2024",
		})
		return
	}

	err := repository.AddBookRepo(connection.Db, &newData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newData,
		"message": "success",
	})

}

// GET BOOK DETAIL

func GetBookDetail(ctx *gin.Context) {
	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cant Find Book With This Id",
		})
		return
	}

	var newData model.Book

	err = repository.GetBookDetailRepo(connection.Db, &newData, id)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Book Data Not Found"})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newData,
		"message": "Success ",
	})
}

// DELETE BOOK

func DeleteBook(ctx *gin.Context) {
	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cant Delete Book With This Id",
		})
		return
	}

	res, err := repository.DeleteBookRepo(connection.Db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	x, _ := res.RowsAffected()

	if x == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Book Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Book",
	})
}

// UPDATE BOOK

func UpdateBook(ctx *gin.Context) {

	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Need Valid Id",
		})
		return
	}

	var data model.Book

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	data.Id = id

	if data.Total_page > 100 {
		data.Thickness = "Tebal"
	} else {
		data.Thickness = "Tipis"
	}

	res, _ := repository.UpdateBookRepo(connection.Db, &data, id)

	x, _ := res.RowsAffected()

	if x == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Cant Update Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Success Update",
	})

}
