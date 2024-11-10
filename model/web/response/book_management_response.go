package response

import "time"

type BookManagementResponse struct {
	BookId       int       `json:"bookId"`
	CategoryId   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	AuthorId     int       `json:"authorId"`
	AuthorName   string    `json:"authorName"`
	Judul        string    `json:"judul"`
	CodeBook     string    `json:"codeBook"`
	DateofPublic time.Time `json:"dateofPublic"`
}
