package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/quanergyO/sigma_evolution/types"
)

type Skill struct {
	db *sql.DB
}

func NewSkill(db *sql.DB) *Skill {
	return &Skill{
		db: db,
	}
}

func (s *Skill) GetAll() ([]types.Skill, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *Skill) GetById(id int) (types.Skill, error) {
	return types.Skill{}, fmt.Errorf("not implented")
}

func (s *Skill) Update(id int) error {
	return fmt.Errorf("not implented")
}

func (s *Skill) Create(skill types.Skill) (int, error) {
	var skillId int
	query := fmt.Sprintf("INSERT INTO %s(name) VALUES($1) RETURNING ID", skillTable)
	slog.Info(query)
	row := s.db.QueryRow(query, skill.Name)
	if err := row.Scan(&skillId); err != nil {
		return 0, err
	}

	return skillId, nil
}

func (s *Skill) Delete(id int) error {
	return fmt.Errorf("not implented")
}
