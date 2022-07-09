package repository

import (
	"github.com/bagasfathoni/go-sqlx/model"
	"github.com/jmoiron/sqlx"
)

type departmentRepository struct {
	repo *sqlx.DB
}

type DepartmentRepository interface {
}

func (d *departmentRepository) GetAll() ([]model.Department, error) {
	var results []model.Department
	err := d.repo.Select(&results, config.SELECT_ALL_DEPARTMENTS)
	return results, err
}

func InitDepartmentRepository(db *sqlx.DB) DepartmentRepository {
	return &departmentRepository{repo: db}
}
