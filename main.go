package main

import (
	"database/sql"
	"encoding/json"
	crud "practice/CRUD"
	"practice/models"

	"github.com/k0kubun/pp/v3"
	_ "github.com/lib/pq"
)

func main() {
	connection := "user=postgres password=mubina2007 dbname=postgres sslmode=disable"
	mydb, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	reqString := []byte(`
	{
		"book": {
			"name": "book1",
			"genre": "genre1",
			"year": 2000
		},
		"author": {
			"fullName": "Author1",
			"birthDate": "1892-01-03"
		}
	}
	`)
	var booksAuthors models.BookAuthor
	if err := json.Unmarshal(reqString, &booksAuthors); err != nil {
		panic(err)
	}
	book, err := crud.CreateBook(mydb, &booksAuthors)
	if err != nil {
		panic(err)
	}
	pp.Println(book)

	pp.Println("----------------------------------------------------------")
	author, err := crud.CreateAuthor(mydb, &booksAuthors)
	if err != nil {
		panic(err)
	}
	pp.Println(author)

	bookAuthor2, err := crud.CreateBookAuthor(mydb, book.Id, author.Id, &booksAuthors)
	pp.Println("----------------------------------------------------------")
	pp.Println(bookAuthor2)
	pp.Println(crud.GetAuthorById(mydb, 3))
	pp.Println(crud.GetAllBooksAuthors(mydb))
}
