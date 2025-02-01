package service

import "github.com/quanergyO/sigma_evolution/internal/repository"

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

type Service struct {
	IBooks
	IPlayLists
	ICourses
	ISchool21Project
	ISkils
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
