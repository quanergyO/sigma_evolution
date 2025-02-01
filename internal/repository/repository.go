package repository

import "database/sql"

type IBooks interface {
}

type IPlayLists interface {
}

type ICourses interface {
}

type ISchool21Project interface {
}

type ISkils interface {
}

type Repository struct {
	IBooks
	IPlayLists
	ICourses
	ISchool21Project
	ISkils
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
