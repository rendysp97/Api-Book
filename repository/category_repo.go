package repository

import (
	"api-book/model"
	"database/sql"
)

func GetAllCategoriesRepo(db *sql.DB, dataCategory *[]model.Categories) {
	rows, err := db.Query("SELECT * FROM categories")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data model.Categories

		err := rows.Scan(&data.Id, &data.Name, &data.Created_at, &data.Created_by, &data.Modified_at, &data.Modified_by)

		if err != nil {
			panic(err)
		}

		*dataCategory = append(*dataCategory, data)

	}

}

func AddCategoryRepo(db *sql.DB, newData *model.Categories) error {

	sqlStatement := `
	insert into categories (name,created_by) Values ($1,$2) returning id 
	`
	err := db.QueryRow(sqlStatement, newData.Name, newData.Created_by).Scan(&newData.Id)

	return err
}

func GetDetailCategoryRepo(db *sql.DB, newData *model.Categories, id int) error {

	err := db.QueryRow("Select * from categories WHERE id = $1", id).Scan(&newData.Id, &newData.Name, &newData.Created_at, &newData.Created_by, &newData.Modified_at, &newData.Modified_by)

	if err != nil {
		return err
	}

	return err
}

func DeleteCategoryRepo(db *sql.DB, id int) (sql.Result, error) {
	return db.Exec("delete from categories where id = $1", id)

}

func GetBookByCategoryRepo(db *sql.DB, data *[]model.Book, id int) error {
	rows, err := db.Query("SELECT * FROM book WHERE category_id = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var b model.Book
		if err := rows.Scan(
			&b.Id,
			&b.Title,
			&b.Description,
			&b.Image_url,
			&b.Release_year,
			&b.Price,
			&b.Total_page,
			&b.Thickness,
			&b.Category_id,
			&b.Created_at,
			&b.Modified_at,
			&b.Created_by,
			&b.Modified_by,
		); err != nil {
			return err
		}
		*data = append(*data, b)
	}
	return nil
}
func UpdateDataCategoryRepo(db *sql.DB, data *model.Categories, id int) (sql.Result, error) {

	res, err := db.Exec(
		"UPDATE categories SET name = $1, created_by = $2 WHERE id = $3",
		data.Name, data.Created_by, id,
	)

	if err != nil {
		panic(err)
	}

	return res, err

}
