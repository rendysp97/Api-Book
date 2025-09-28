package repository

import (
	"api-book/model"
	"database/sql"
)

func GetAllBookRepo(db *sql.DB, data *[]model.Book) error {

	rows, err := db.Query("SELECT * FROM book")

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		var newData model.Book

		err := rows.Scan(&newData.Id,
			&newData.Title,
			&newData.Description,
			&newData.Image_url,
			&newData.Release_year,
			&newData.Price,
			&newData.Total_page,
			&newData.Thickness,
			&newData.Category_id,
			&newData.Created_at,
			&newData.Modified_at,
			&newData.Created_by,
			&newData.Modified_by)

		if err != nil {
			return err
		}

		*data = append(*data, newData)

	}

	return nil

}

func AddBookRepo(db *sql.DB, newData *model.Book) error {
	sqlStatement := `
		INSERT INTO book (
			title, description, image_url, release_year, price, total_page, thickness, category_id, created_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id
	`

	err := db.QueryRow(sqlStatement,
		newData.Title,
		newData.Description,
		newData.Image_url,
		newData.Release_year,
		newData.Price,
		newData.Total_page,
		newData.Thickness,
		newData.Category_id,
		newData.Created_by,
	).Scan(&newData.Id)

	return err
}

func GetBookDetailRepo(db *sql.DB, newData *model.Book, id int) error {
	err := db.QueryRow(`
		SELECT * FROM book WHERE id = $1
	`, id).Scan(
		&newData.Id,
		&newData.Title,
		&newData.Description,
		&newData.Image_url,
		&newData.Release_year,
		&newData.Price,
		&newData.Total_page,
		&newData.Thickness,
		&newData.Category_id,
		&newData.Created_at,
		&newData.Modified_at,
		&newData.Created_by,
		&newData.Modified_by,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBookRepo(db *sql.DB, book *model.Book, id int) (sql.Result, error) {
	res, err := db.Exec(`
		UPDATE book 
		SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=NOW()
		WHERE id=$9
	`, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Category_id, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteBookRepo(db *sql.DB, id int) (sql.Result, error) {
	return db.Exec("delete from book where id = $1", id)
}
