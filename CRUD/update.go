package crud

import "database/sql"

func UpdateBookNameById(mydb *sql.DB, bookId int, newName string) error {
	tx, err := mydb.Begin()
	if err != nil {
		return err
	}
	query := `UPDATE books SET name = $1 WHERE id = $2`
	_, err = tx.Exec(query, newName, bookId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateBookGenreById(mydb *sql.DB, bookId int, newGenre string) error {
	tx, err := mydb.Begin()
	if err != nil {
		return err
	}
	query := `UPDATE books SET genre = $1 WHERE id = $2`
	_, err = tx.Exec(query, newGenre, bookId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateAuthorFullNameById(mydb *sql.DB, authorId int, newFullName string) error {
	tx, err := mydb.Begin()
	if err != nil {
		return err
	}
	query := `UPDATE authors SET full_name = $1 WHERE id = $2`
	_, err = tx.Exec(query, newFullName, authorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateAuthorBirthDateByid(mydb *sql.DB, authorId int, newBirthDate string) error {
	tx, err := mydb.Begin()
	if err != nil {
		return err
	}
	query := `UPDATE authors SET birth_date = $1 WHERE id = $2`
	_, err = tx.Exec(query, newBirthDate, authorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
