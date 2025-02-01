package repository

import (
	"database/sql"

	"github.com/quanergyO/sigma_evolution/internal/repository/postgres"
	"github.com/quanergyO/sigma_evolution/types"
)

type IBooks interface {
}

type IPlayLists interface {
}

type ICourses interface {
}

type ISchool21Project interface {
}

type ISkils interface {
	GetAll() ([]types.Skill, error)
	GetById(id int) (types.Skill, error)
	Update(id int) error
	Create(skill types.Skill) (int, error)
	Delete(id int) error
}

type Repository struct {
	IBooks
	IPlayLists
	ICourses
	ISchool21Project
	ISkils
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ISkils: postgres.NewSkill(db),
	}
}
