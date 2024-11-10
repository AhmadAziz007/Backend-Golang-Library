package update

import "time"

type BookManagementUpdateRequest struct {
	BookId       int       `validate:"required"`
	CategoryId   int       `validate:"required"`
	AuthorId     int       `validate:"required"`
	Judul        string    `validate:"required,min=1,max=100" json:"judul"`
	CodeBook     string    `validate:"required,min=1,max=100" json:"codeBook"`
	DateofPublic time.Time `validate:"required" json:"dateofPublic"`
}
