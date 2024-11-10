package domain

import "time"

type BookManagement struct {
	BookId       int
	CategoryId   int
	CategoryName string
	AuthorId     int
	AuthorName   string
	Judul        string
	CodeBook     string
	DateofPublic time.Time
}
