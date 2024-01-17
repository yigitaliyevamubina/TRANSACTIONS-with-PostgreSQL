package models

type Book struct {
	Id    int
	Name  string
	Genre string
	Year  int
}

type Author struct {
	Id        int
	FullName  string
	BirthDate string
}

type BookAuthor struct {
	Id     int
	Book   Book
	Author Author
}
