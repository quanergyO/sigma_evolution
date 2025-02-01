package service

import (
	"github.com/quanergyO/sigma_evolution/internal/repository"
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

type Service struct {
	IBooks
	IPlayLists
	ICourses
	ISchool21Project
	ISkils
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ISkils: NewSkillService(repo),
	}
}
