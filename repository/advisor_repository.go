package repository

import (
	"github.com/bagasfathoni/go-sqlx/config"
	"github.com/bagasfathoni/go-sqlx/model"
	"github.com/jmoiron/sqlx"
)

type advisorRepository struct {
	repo *sqlx.DB
}

type AdvisorRepository interface {
	GetAll() ([]model.Advisor, error)
}

func (a *advisorRepository) GetAll() ([]model.Advisor, error) {
	var results []model.Advisor
	err := a.repo.Get(&results, config.SELECT_ALL_ADVISORS)
	return results, err
}

func InitAdvisorRepository(db *sqlx.DB) AdvisorRepository {
	advRepo := new(advisorRepository)
	advRepo.repo = db
	return advRepo
}
