package repository

import "github.com/nattatorn-dev/go-clean-arch/author"

type AuthorRepository interface {
	GetByID(id int64) (*author.Author, error)
}
