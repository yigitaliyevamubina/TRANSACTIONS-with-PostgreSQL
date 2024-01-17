package crud

import "database/sql"

func DeleteBookById(mydb *sql.DB, bookId int) error {
	tx, err := mydb.Begin()
	if err != nil {
		return err
	}

	query1 := `DELETE FROM books_authors WHERE book_id = $1`
	_, err = tx.Exec(query1, bookId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query2 := `DELETE FROM books WHERE id = $1`
	_, err = tx.Exec(query2, bookId)
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

func DeleteAuthorById(mydb *sql.DB, authorId int) error {
	tx, err := mydb.Begin()
	if err != nil {
		return err
	}
	query1 := `DELETE FROM books_authors WHERE author_id = $1`
	_, err = tx.Exec(query1, authorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query2 := `DELETE FROM authors WHERE id = $1`
	_, err = tx.Exec(query2, authorId)
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
