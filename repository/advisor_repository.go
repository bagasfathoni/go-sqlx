package repository

import (
	"github.com/jmoiron/sqlx"
)

type advisorRepository struct {
	repo *sqlx.DB
}

type AdvisorRepository interface {
}

// func (a *advisorRepository) FindAllBy(by map[string]interface{}) []model.Advisor {
// 	var results []model.Advisor

// }

func InitAdvisorRepository(db *sqlx.DB) AdvisorRepository {
	advRepo := new(advisorRepository)
	advRepo.repo = db
	return advRepo
}
