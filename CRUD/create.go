package crud

import (
	"database/sql"
	"practice/models"
)

func CreateBook(mydb *sql.DB, bookAuthor *models.BookAuthor) (*models.Book, error) {
	tx, err := mydb.Begin()
	if err != nil {
		return nil, err
	}
	var book models.Book
	query := `INSERT INTO books(name, genre, year) VALUES($1, $2, $3) returning id, name, genre, year`
	rowBook := tx.QueryRow(query, bookAuthor.Book.Name, bookAuthor.Book.Genre, bookAuthor.Book.Year)
	if err := rowBook.Scan(&book.Id, &book.Name, &book.Genre, &book.Year); err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func CreateAuthor(mydb *sql.DB, bookAuthor *models.BookAuthor) (*models.Author, error) {
	tx, err := mydb.Begin()
	if err != nil {
		return nil, err
	}
	var author models.Author
	query := `INSERT INTO authors(full_name, birth_date) VALUES($1, $2) returning id, full_name, birth_date`
	rowAuthor := tx.QueryRow(query, bookAuthor.Author.FullName, bookAuthor.Author.BirthDate)
	if err := rowAuthor.Scan(&author.Id, &author.FullName, &author.BirthDate); err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func CreateBookAuthor(mydb *sql.DB, booksId int, authorId int, bookAuthor *models.BookAuthor) (*models.BookAuthor, error) {
	tx, err := mydb.Begin()
	if err != nil {
		return nil, err
	}
	query := `INSERT INTO books_authors(book_id, author_id) VALUES($1, $2) returning id, book_id, author_id`
	rowBookAuthor := tx.QueryRow(query, booksId, authorId)
	if err := rowBookAuthor.Scan(&bookAuthor.Id, &bookAuthor.Book.Id, &bookAuthor.Author.Id); err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return bookAuthor, nil
}
