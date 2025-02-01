package service

import (
	"github.com/quanergyO/sigma_evolution/internal/repository"
	"github.com/quanergyO/sigma_evolution/types"
)

type SkillService struct {
	repo *repository.Repository
}

func NewSkillService(repo *repository.Repository) *SkillService {
	return &SkillService{
		repo: repo,
	}
}

func (s *SkillService) GetAll() ([]types.Skill, error) {
	return s.repo.ISkils.GetAll()
}

func (s *SkillService) GetById(id int) (types.Skill, error) {
	return s.repo.ISkils.GetById(id)
}

func (s *SkillService) Update(id int) error {
	return s.repo.ISkils.Update(id)
}

func (s *SkillService) Create(skill types.Skill) (int, error) {
	return s.repo.ISkils.Create(skill)
}

func (s *SkillService) Delete(id int) error {
	return s.repo.ISkils.Delete(id)
}
