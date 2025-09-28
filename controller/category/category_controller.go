package category

import (
	"api-book/database/connection"
	"api-book/model"
	"api-book/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET ALL CATEGORY

func Getallcategories(ctx *gin.Context) {

	var allData []model.Categories

	repository.GetAllCategoriesRepo(connection.Db, &allData)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    allData,
		"message": "success",
	})

}

// ADD CATEGORY

func AddCategory(ctx *gin.Context) {

	var newData model.Categories

	if err := ctx.ShouldBindJSON(&newData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newData.Created_by == "" {
		username, _ := ctx.Get("username")
		newData.Created_by = username.(string)
	}

	err := repository.AddCategoryRepo(connection.Db, &newData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    newData,
		"message": "success",
	})

}

// GET DETAIL CATEGORY

func GetDetailCategory(ctx *gin.Context) {
	fromParamsId := ctx.Param("id")

	id, err := strconv.Atoi(fromParamsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Need Valid ID",
		})
		return
	}

	var dataCat model.Categories

	err = repository.GetDetailCategoryRepo(connection.Db, &dataCat, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Cant Find Category Id",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    dataCat,
		"message": "Success",
	})
}

// DELETE CATEGORY

func DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	res, err := repository.DeleteCategoryRepo(connection.Db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Id Exist Based Book Id"})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success Delete Category"})
}

// GET BOOK BY CATEGORY
func GetBookByCategory(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Need Valid Id"})
		return
	}

	var data []model.Book
	err = repository.GetBookByCategoryRepo(connection.Db, &data, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(data) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No books found for this category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Success",
	})
}

func UpdateDataCategory(ctx *gin.Context) {

	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Need Valid Id",
		})
		return
	}

	var data model.Categories

	data.Id = id

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	res, _ := repository.UpdateDataCategoryRepo(connection.Db, &data, id)

	x, _ := res.RowsAffected()

	if x == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Cant Update Not Found Id ",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Success Update",
	})
}
