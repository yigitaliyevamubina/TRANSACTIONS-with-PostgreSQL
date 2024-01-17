package crud

import (
	"database/sql"
	"practice/models"
)

func GetBookById(mydb *sql.DB, bookId int) (*models.Book, error) {
	tx, err := mydb.Begin()
	if err != nil {
		return nil, err
	}
	query := `SELECT id, name, genre, year FROM books WHERE id = $1`
	var book models.Book
	respBook := tx.QueryRow(query, bookId)
	if err := respBook.Scan(&book.Id, &book.Name, &book.Genre, &book.Year); err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func GetAuthorById(mydb *sql.DB, authorId int) (*models.Author, error) {
	tx, err := mydb.Begin()
	if err != nil {
		return nil, err
	}
	query := `SELECT id, full_name, birth_date FROM authors WHERE id = $1`
	var author models.Author
	respAuthor := tx.QueryRow(query, authorId)
	if err := respAuthor.Scan(&author.Id, &author.FullName, &author.BirthDate); err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func GetAllBooksAuthors(mydb *sql.DB) ([]*models.BookAuthor, error) {
	tx, err := mydb.Begin()
	if err != nil {
		return nil, err
	}
	query := `SELECT b.id, 
					b.name, 
					b.genre, 
					b.year, 
					a.id, 
					a.full_name, 
					a.birth_date 
					FROM books b JOIN 
					books_authors ba 
					ON b.id = ba.book_id 
					JOIN authors a ON a.id = ba.author_id`
	rows, err := tx.Query(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var booksAuthors []*models.BookAuthor
	for rows.Next() {
		var bookAuthor models.BookAuthor
		if err := rows.Scan(&bookAuthor.Book.Id,
			&bookAuthor.Book.Name,
			&bookAuthor.Book.Genre,
			&bookAuthor.Book.Year,
			&bookAuthor.Author.Id,
			&bookAuthor.Author.FullName,
			&bookAuthor.Author.BirthDate); err != nil {
			tx.Rollback()
			return nil, err
		}
		booksAuthors = append(booksAuthors, &bookAuthor)
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return booksAuthors, nil
}
